# Add SignalR client dependencies

- Two options:
	- Via Content Delivery Network (CDN)
	- Via NPM

# Overview of JS SignalR client implementation

```js
const connection = new signalR.HubConnectionBuilder()
    .withUrl("/learningHub")
    .configureLogging(signalR.LogLevel.Information)
    .build();

// Event listener
connection.on("ReceiveMessage", (message) => {
    $("#signalr-message-panel").prepend($("<div />").text(message));
});

// Associate a click event with an HTML button
$("#btn-broadcast").click(function () {
    var message = $("#broadcast").val();
    connection
        .invoke("BroadcastMessage", message)
        .catch((err) => console.log(err.toString()));
});

// Trigger the SignalR connection to start
async function start() {
    try {
        await connection.start();
        console.log('connected');
    } catch (err) {
        console.log(err);
        // Execute function after a delay
        setTimeout(() => {
            start();
        }, 5000);
    }
};

//If the connection breaks at any point, it will automatically restart
connection.onclose(async () => {
    await start();
})

start();
```

1. Event listener => When the server-side code triggers the event, the message is *prepended inside a panel*
2. Associate a click event with the HTML button => Trigger the `BroadcastMessage()` method inside the server-side hub while passing the message
3. Associate the connection with the `onclose()` event to keep the connection alive as long as possible

## Blazor client

```html
<div>
	<label for="broadcastMsg">Message</label>
	<input @bind="message" type="text" id="broadcastMsg" name="broadcastMsg"/>
</div>
```

- the `@bind`directive **binds the content of the `input` element to a C# variable called `message`** => Whenever you change the text, the content will change
- the `@onclick` directive **triggers the `BroadcastMessage` C# method**

```cs
protected override async Task OnInitializedAsync()
{
	hubConnection = new HubConnectionBuilder()
		.WithUrl(NavigationManager.ToAbsoluteUri("/learningHub"))
		.Build();

	hubConnection.On<string>("ReceiveMessage", (message) =>
	{
		messages.Add(message);
		StateHasChanged(); // Blazor-specific method
	});

	await hubConnection.StartAsync();
}
```

- The `StateHasChanged()` method specific to Blazor monitor changes to elements => Panel always populated with messages received from the server