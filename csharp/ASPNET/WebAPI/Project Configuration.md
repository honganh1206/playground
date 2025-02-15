# Middleware 

- ASP.NET Core middleware is **a piece of code integrated inside the application’s pipeline** that we can use to **handle requests and responses**.
- Configure request delegates by using **Run/Map/Use** extension methods

# Extension methods
## `Run`
- Terminate the middleware pipeline and generate a response => Simple tasks like returning a fixed response/performing an action without further processing
```cs
app.Run(async context =>
{
    Console.WriteLine($"Writing the response to the client in the Run method");
    await context.Response.WriteAsync("Hello from the middleware component.");
});
```
## `Use`: 
- Add middleware components to pipeline such as authentication, logging, exception handling, etc.
```cs
app.Use(async (context, next) =>
{
    Console.WriteLine($"Logic BEFORE executing the next delegate in the Use method");
    await next.Invoke();
    Console.WriteLine($"Logic AFTER executing the next delegate in the Use method");
});
```
## `Map`
- Branch the pipeline based on URL prefix (can be combined with other extension methods)
```cs
app.Map("/usingmapbranch", builder =>
{
    builder.Use(async (context, next) =>
    {
        Console.WriteLine("Map branch logic in the Use method BEFORE the next delegate");
        await next.Invoke();
        Console.WriteLine("Map branch logic in the Use method AFTER the next delegate");
    });
    builder.Run(async context =>
    {
        Console.WriteLine($"Writing the response to the client in the Run method");
        await context.Response.WriteAsync("Hello from the map branch");
    });
});
```
- `MapWhen`: Similar to **Map** but instead of checking the URL prefix, it checks a predicate function to **determine whether to branch the pipeline or not**
```cs
app.MapWhen(context => context.Request.Query.ContainsKey("testquerystring"), builder =>
{
    builder.Run(async context =>
    {
        await context.Response.WriteAsync("Hello from the MapWhen Branch");
    });
});
```

# Middleware flow

![[Pasted image 20230714153113.png]]

### Order for middleware components

![[Pasted image 20230714153441.png]]

