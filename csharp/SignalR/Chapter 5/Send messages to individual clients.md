```cs
// LearningHub.cs
public async Task SendToOthers(string message)
{
	await Clients.Others.ReceiveMessage(message);
}
```