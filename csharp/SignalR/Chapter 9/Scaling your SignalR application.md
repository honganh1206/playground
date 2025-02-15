# Setting up Redis backplane

- A **backplane** (in the context of Redis backplane) is a communication channel that **allows communication and message distribution across multiple (SignalR) servers**
- Redis backplane allows the hubs to **exchange information between the instances of SignalR hubs** connected to it. 
- Redis backplane can **transfer the message to the right instance of the application/the right client**

![[Pasted image 20240103135546.png]]

```bash
# Access the cli
redis cli
# Insertion command
set client:1 "Test"
# Verify Redis cache is working
get client:1

```

# Running multiple hub instances via Redis backplane

- Entry point into the application includes
	- Reverse proxy
	- Load balancer
	- Combination of both
- Current goal: **Manually create two copies of the web application** and **connect them to the same instance of Redis backplane**

# Hosting the same SignalR hub in in different applications

- Same hub code in multiple different applications => SignalR can be scaled **independently** from the web application that hosts it with Redis
- Scenario: 
	- Component #1: A **service** for devices to *make occasional client-initiated call* and **SignalR connection** for the server application to *issue real-time instructions to these devices*
	- Component #2: A **web page** for users to *access and view real-time device status information* managed by SignalR
	- Issue: The number of IoT devices that will be connected to the application will be **different** from the number of users looking at the devices
	- Solution: 
		- #1: Scale the two components independently to **two applications**
		- #2: Use the same hub hosted by **both the user-facing web app and the web API service for the IoT device**

# Using `HubContext` to send messages from outside SignalR hub

- `HubContext` allows you to **access all fields of a specific SignalR hub, including its groups, clients and so on**.

> [!tip] Using `HubContext` for scaling
> 
> `HubContext` is not designed specifically for a scaled SignalR hub, so you can use it with monolithic SignalR hubs too


