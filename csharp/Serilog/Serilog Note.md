---
tags:
  - "#study"
  - "#review"
cssclasses:
  - center-images
---

- [Guide to configure Serilog](https://blog.postsharp.net/serilog-aspnetcore)
# [Logging levels](https://code-maze.com/csharp-different-log-levels-in-serilog/)

![[Pasted image 20231224000027.png]]

- Configuring `MinimumLevel`
	- Generate application logs for levels **equal to or higher than `Information` level**
	- Override the default logging of Microsoft and System => Record logs of level `Warning` or higher

```json
"Serilog": {
  "MinimumLevel": {
    "Default": "Information",
    "Override": {
      "Microsoft": "Warning",
      "System": "Warning"
    }
  }
}
```

## Config example

```json
// logsettings.json
{
  "Serilog": {
    "Using": [
      "Serilog.Sinks.File",
      "Serilog.Sinks.Console"
    ],
    "MinimumLevel": {
      "Default": "Debug",
      "Override": {
        "Microsoft.AspNetCore": "Warning"
      }
    },
    "Enrich": [ "FromLogContext", "WithMachineName", "WithProcessId", "WithThreadId" ],
    "WriteTo": [
      {
        "Name": "File",
        "Args": {
          "path": "logs/log.json",
          "rollingInterval": "Day",
          "formatter": "Serilog.Formatting.Compact.CompactJsonFormatter, Serilog.Formatting.Compact"
        }
      },
      {
        "Name": "Console",
        "Args": {
          "theme": "Serilog.Sinks.SystemConsole.Themes.AnsiConsoleTheme::Code, Serilog.Sinks.Console",
          "outputTemplate": "{Timestamp:yyyy-MM-dd HH:mm:ss.fff zzz} [{Level:u3}] {Message:lj}{NewLine}{Exception}"
        }
      }
    ]
  },
  "Serilog.Sinks.Console": {
    "theme": {
      "ansiTheme": {
        "fatal": "red,bold",
        "error": "red",
        "warning": "yellow",
        "information": "green",
        "debug": "cyan",
        "verbose": "gray"
      }
    }
  }
}
```


# Sinks

- Destinations for our log messages