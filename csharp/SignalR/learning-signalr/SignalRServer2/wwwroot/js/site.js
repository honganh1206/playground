// Please see documentation at https://docs.microsoft.com/aspnet/core/client-side/bundling-and-minification
// for details on configuring this project to bundle and minify static web assets.

// Write your JavaScript code.
const connection = new signalR.HubConnectionBuilder()
    .withUrl("/learningHub", {
        transport:
            signalR.HttpTransportType.WebSockets |
            signalR.HttpTransportType.LongPolling,
        headers: { Key: "Value" },
        // Apply authentication token
        accessTokenFactory: null,
        logMessageContent: true,
        // If WebSocket is used => skip negotiation to establish connection faster
        skipNegotiation: false,
        // Apply credentials to the request => CORS requests
        withCredentials: true,
        // Maximum allowed time for HTTP requests
        timeout: 100000,
    })
    // .withHubProtocol(new signalR.protocols.msgpack.MessagePackHubProtocol())
    .configureLogging(signalR.LogLevel.Information)
    .build();

// Determine how long the connection should wait for a message from the server
connection.serverTimeoutInMilliseconds = 30000; // 30 seconds
// Client sends ping messages to the server to keep the connection alive
connection.keepAliveIntervalInMilliseconds = 15000; // 15 seconds


// Event listener
connection.on("ReceiveMessage", (message) => {
    $("#signalr-message-panel").prepend($("<div />").text(message));
});

// Server streaming
$("#btn-trigger-stream").click(function () {
    var numOfJobs = parseInt($("#number-of-jobs").val(), 10);

    connection.stream("TriggerStream", numOfJobs).subscribe({
        next: (message) =>
            $("#signalr-message-panel").prepend($("<div />").text(message)),
    });
});

// Client streaming
$("#btn-broadcast").click(function () {
    var message = $("#broadcast").val();
    if (message.includes(";")) {
        var messages = message.split(";");

        var subject = new signalR.Subject(); // Represent an abstraction of the open stream
        connection
            .send("BroadcastStream", subject)
            .catch((err) => console.error(err.toString()));

        for (let i = 0; i < messages.length; i++) {
            subject.next(messages[i]);
        }

        subject.complete();
    } else {
        connection
            .invoke("BroadcastMessage", message)
            .catch((err) => console.log(err.toString()));
    }
});

$("#btn-others-message").click(function () {
    var message = $("#others-message").val();
    connection
        .invoke("SendToOthers", message)
        .catch((err) => console.log(err.toString()));
});

$("#btn-self-message").click(function () {
    var message = $("#self-message").val();
    connection
        .invoke("SendToCaller", message)
        .catch((err) => console.log(err.toString()));
});

$("#btn-individual-message").click(function () {
    var message = $("#individual-message").val();
    var connectionId = $("#connection-for-message").val();
    connection
        .invoke("SendToIndividual", connectionId, message)
        .catch((err) => console.log(err.toString()));
});

// Btns for groups

$("#btn-group-message").click(function () {
    var message = $("#group-message").val();
    var group = $("#group-for-message").val();
    connection
        .invoke("SendToGroup", group, message)
        .catch((err) => console.error(err.toString()));
});

$("#btn-group-add").click(function () {
    var group = $("#group-to-add").val();
    connection
        .invoke("AddUserToGroup", group)
        .catch((err) => console.error(err.toString()));
});

$("#btn-group-remove").click(function () {
    var group = $("#group-to-remove").val();
    connection
        .invoke("RemoveUserFromGroup", group)
        .catch((err) => console.error(err.toString()));
});

// Trigger the SignalR connection to start
async function start() {
    try {
        await connection.start();
        console.log("connected");
    } catch (err) {
        console.log(err);
        // Execute function after a delay
        setTimeout(() => {
            start();
        }, 5000);
    }
}

// If the connection breaks at any point, it will automatically restart
connection.onclose(async () => {
    await start();
});

start();
