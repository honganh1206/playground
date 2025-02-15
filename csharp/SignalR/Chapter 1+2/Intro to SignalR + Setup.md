- The ability for applications to **communicate with the server asynchronously**
- The need to **receive a real-time** update from the server that was *actually triggered by an event on the server* and  not by a client request
- Use cases:
	- Clients need to **receive real-time updates from the server**
	- **High frequency of data exchange** between the client and the server

# Setting up the solution

```bash
dotnet new sln
# Create a proj based on ASP.NET MVC
dotnet new mvc -o SignalRServer
```

# Setting up SignalR hub

- **Hub** is where the **server-side SignalR components center around**
- Hub is equivalent to MVC/WebAPI controller

```ad-warning
title: Hub lifetime
SignalR hub lifetime is **restricted to a single call** => Do not store any durable data in it
```

# Make SignalR hub strongly-typed

---
# Exercises

1. C
2. D
3. C