

- [.NET Background Service](https://blog.jetbrains.com/dotnet/2023/05/09/dotnet-background-services/)

- [BackgroundService on Microsoft Learn](https://learn.microsoft.com/en-us/aspnet/core/fundamentals/host/hosted-services?view=aspnetcore-8.0&tabs=visual-studio#backgroundservice-base-class)


- A base class for implementing [IHostedService](https://learn.microsoft.com/en-us/dotnet/api/microsoft.extensions.hosting.ihostedservice)
- Use cases: Update globally-used cache, process queued work, monitor service health, etc.

## What is `BackgroundService` class?

- Allow us to run tasks *independently of the main application thread*
- A class for implementing **long-running processes**
- `ExecuteAsync` helps reduce the implementation complexity

```cs
namespace App.Sample;

public class Worker : BackgroundService
{
    private readonly ILogger<Worker> _logger;

    public Worker(ILogger<Worker> logger)
    {
        _logger = logger;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
	    // This will keep running as long as the application executes
        while (!stoppingToken.IsCancellationRequested)
        {
	        // Handle processing logic here
            _logger.LogInformation("Worker running at: {time}", DateTimeOffset.Now);
            await Task.Delay(1000, stoppingToken);
        }
    }
}
```



> [!warning]
> The host application is **not** limited to 1 background service, but with many registered background services, *one can cause issues to others*.


With `IHostedService`, this will run the log information first before the application starts

```cs
// Implementing IHostedService
public async Task StartAsync(CancellationToken cancellationToken)
{
	_logger.LogInformation("StartAsync");
	await Task.Delay(10000);
}
```

But we can configure so that they don’t have to wait for each other. In this way, the background service will *start concurrently with the application*:

```cs
builder.Services.Configure<HostOptions>(x => {
	x.ServicesStartConcurrently = true;
	x.ServicesStopConcurrently = false;
})
```

More control with `IHosttedLifecycleService` 