# Client/Server rendered UI or Hybrid
## Server
- Pros:
	- Minimal client requirement
	- Access flexibility
- Cons:
	- The server needs to handle most of the computing power + memory usage
- Examples: ASP.NET Core Razor Pages, Core MVC, Blazor Server
## Client
- Pros
	- Rich interactivity
	- Disconnected mode + Updates to client-side model after reconnection
- Cons:
	- Increase load time (Clients' device needs to handle the login)
	- Little support for low-end devices
- Examples: Blazor WebAssembly, ASP.NET Core SPA with JS frameworks
## Hybrid: ASP.NET Core MVC or Razor Pages plus Blazor