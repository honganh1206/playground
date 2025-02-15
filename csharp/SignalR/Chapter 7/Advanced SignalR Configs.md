
```ad-summary
1. SignalR server has **three levels of configuration**:
	a. Top-level SignalR configs
	b. Message protocol configs
	c. HTTP transport configs
2. Different SignalR client types have different config options, but **all clients support configs of transport mechanism**
3. MessagePack protocol can store data in binary format
	a. Stricter than JSON
	b. Better for performance

```


# MessagePack protocol 

- Similar to JSON in terms of structure but **it is binary rather than textual** => Cannot read the message while it is being transferred + 