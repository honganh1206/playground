# Binding Configuration

- The `Configuration` class in ASP.NET core is used to retrieve values from `appsettings.json` files or environment variables
- The `Configuration.Bind` method is used to bind configuration values from a specific section of the configuration (`appsettings.json`, environment variables) to an instance of a POCO (Plain Old CLR Object) class
- **POCO** is like an **empty box** that you can use to **store and organize specific types of information**.

```json
{
  "AppSettings": {
    "ApiKey": "your-api-key",
    "BaseUrl": "https://api.example.com",
    "MaxItemsPerPage": 10
  }
}
```

```cs
// POCO
public class AppSettings
{
    public string ApiKey { get; set; }
    public string BaseUrl { get; set; }
    public int MaxItemsPerPage { get; set; }
}

// Bind config values with POCO
public void ConfigureServices(IServiceCollection services)
{
    // Other service configurations...

    // Bind configuration to AppSettings class
    var appSettings = new AppSettings();
    Configuration.GetSection("AppSettings").Bind(appSettings);
    services.AddSingleton(appSettings);
}


```

---
# Options Pattern

- Similar to Binding Configuration but with validation + live reloading + easier testing
- Once we configure the class containing our configuration we can inject it via dependency injection with `IOptions<T>` and thus injecting only part of our configuration or rather only the part that we need.


### Modify values in configuration files with `IOptionSnapshot` and `IOptionsMonitor`

|     | `Ioptions<T>`                                                                                                                                      | `IOptionsSnapshot<T>`                                                                                                                                                   | `IoptionsMonitor<T>`                                                                                                                                                                                   |
| --- | -------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
|     | - Original Options interface </br> - NO support config reloading </br> - Only bind config during registration </br> - NO support for named options | - Scoped service </br> - Support config reloading </br> - **Cannot** be injected into singleton service </br> - Values reload per request </br> - Support named options | - Registered as **singleton** service </br> - Support config reloading </br> - **Can** be injected into any service lifetime </br> - Cache and reload values immediately </br> - Support named options |



