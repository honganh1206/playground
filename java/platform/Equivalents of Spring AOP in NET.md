---
id: Equivalents of Spring AOP in NET
aliases:
  - Equivalents of Spring AOP in NET
tags: []
---

# Equivalents of Spring AOP in NET

ASP.NET Core doesn't have a direct one-to-one equivalent to Spring AOP, but it offers several mechanisms that achieve similar results:

1. Middleware

Middleware in ASP.NET Core handles cross-cutting concerns in HTTP request processing.

```cs
public class RequestLoggingMiddleware
{
    private readonly RequestDelegate _next;
    private readonly ILogger _logger;

    public RequestLoggingMiddleware(RequestDelegate next, ILoggerFactory loggerFactory)
    {
        _next = next;
        _logger = loggerFactory.CreateLogger<RequestLoggingMiddleware>();
    }

    public async Task InvokeAsync(HttpContext context)
    {
        _logger.LogInformation($"Request starting: {context.Request.Path}");

        var stopwatch = Stopwatch.StartNew();
        await _next(context);
        stopwatch.Stop();

        _logger.LogInformation($"Request completed in {stopwatch.ElapsedMilliseconds}ms");
    }
}

// In Program.cs or Startup.cs
app.UseMiddleware<RequestLoggingMiddleware>();
```

2. Filters

Filters in ASP.NET Core are similar to aspects but specific to MVC actions.

```cs
public class LogActionFilter : IActionFilter
{
    private readonly ILogger _logger;
    private Stopwatch _stopwatch;

    public LogActionFilter(ILoggerFactory loggerFactory)
    {
        _logger = loggerFactory.CreateLogger<LogActionFilter>();
    }

    public void OnActionExecuting(ActionExecutingContext context)
    {
        _logger.LogInformation($"Executing action {context.ActionDescriptor.DisplayName}");
        _stopwatch = Stopwatch.StartNew();
    }

    public void OnActionExecuted(ActionExecutedContext context)
    {
        _stopwatch.Stop();
        _logger.LogInformation($"Action {context.ActionDescriptor.DisplayName} executed in {_stopwatch.ElapsedMilliseconds}ms");
    }
}

// In Controller or globally
[TypeFilter(typeof(LogActionFilter))]
public class HomeController : Controller
{
    // ...
}
```
