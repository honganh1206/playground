using System.Runtime.CompilerServices;
using Microsoft.AspNetCore.SignalR;

namespace SignalRHubs
{
    public class LearningHub : Hub<ILearningHubClient>
    {
        /// <summary>
        /// Receive a message from a client then re-send the messages to all clients, including the one sending the message
        /// </summary>
        /// <param name="message"></param>
        /// <returns></returns>
        public async Task BroadcastMessage(string message)
        {
            // Clients include all details of all connected clients
            await Clients.All.ReceiveMessage(GetMessageToSend(message));
        }

        // Client streaming
        public async Task BroadcastStream(IAsyncEnumerable<string> stream)
        {
            await foreach (var item in stream)
            {
                await Clients.Caller.ReceiveMessage($"Server received {item}");
            }
        }

        // Server streaming
        public async IAsyncEnumerable<string> TriggerStream(
            int jobsCount,
            [EnumeratorCancellation] // Trigger cancellation if calling client issues a cancellation
            CancellationToken cancellationToken
        )
        {
            for (int i = 0; i < jobsCount; i++)
            {
                cancellationToken.ThrowIfCancellationRequested();
                yield return $"Job {i} executed successfully";
                await Task.Delay(1000, cancellationToken);
            }
        }

        public override async Task OnConnectedAsync()
        {
            await base.OnConnectedAsync();
        }

        public override async Task OnDisconnectedAsync(Exception? exception)
        {
            await base.OnDisconnectedAsync(exception);
        }

        // Broadcast message to all clients except sender
        public async Task SendToOthers(string message)
        {
            await Clients.Others.ReceiveMessage(GetMessageToSend(message));
        }

        // Broadcast messages to the sender only
        public async Task SendToCaller(string message)
        {
            // await Clients.Caller.ReceiveMessage(GetMessageToSend(message));
            throw new Exception("You cannot send messages to yourself");
        }

        // Broadcast messages to specific client(s)
        public async Task SendToIndividual(string connectionId, string message)
        {
            await Clients.Client(connectionId).ReceiveMessage(GetMessageToSend(message));
        }

        public async Task SendToMultipleIndividuals(string connectionIds, string message)
        {
            // TODO: Add a method to process connectionIds as string to List<string>
            await Clients.Clients(connectionIds).ReceiveMessage(GetMessageToSend(message));
        }
        private string GetMessageToSend(string originalMsg)
        {
            return $"User connection Id: {Context.ConnectionId}. Message: {originalMsg}";
        }

        private string GetMessageToSendToGroup(string originalMsg, string groupName)
        {
            return $"User connection Id: {Context.ConnectionId}. Group name: {groupName}. Message: {originalMsg}";
        }

        // Send messages to groups
        public async Task SendToGroup(string groupName, string message)
        {
            await Clients.Group(groupName).ReceiveMessage(GetMessageToSendToGroup(message, groupName));
        }

        public async Task AddUserToGroup(string groupName)
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, groupName);
            await Clients.Caller.ReceiveMessage($"Current user added to {groupName} group");
            await Clients.Others.ReceiveMessage($"User {Context.ConnectionId} added to {groupName} group");
        }

        public async Task RemoveUserFromGroup(string groupName)
        {
            await Groups.RemoveFromGroupAsync(Context.ConnectionId, groupName);
            await Clients.Caller.ReceiveMessage($"Current user removed from {groupName} group");
            await Clients.Others.ReceiveMessage($"User {Context.ConnectionId} removed from {groupName} group");
        }
    }
}
