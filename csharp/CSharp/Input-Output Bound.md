---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#computer"
cssclasses:
  - center-images
---
The `await` keyword pauses the execution while *waiting for the method marked with `async` to return a task*


```cs
// Notice that the method is declared as 'async' and returns a Task<string>
public async Task<string> DownloadDataAsync() {
    using (var httpClient = new HttpClient()) {

        // Use the await keyword to execute a non-blocking GET request
        string stringResult = await httpClient.GetStringAsync("https://ndepend.com/data");

        // Return the result obtained when the request is completed
        return stringResult;
    }
}

```

The .NET library offers many async methods for different I/O tasks


```cs
var tasks = new Task<string>[] {
   File.ReadAllTextAsync(@"C:\Program Files\dotnet\dotnet.exe"),
   File.ReadAllTextAsync(@"C:\Windows\explorer.exe"),
   File.ReadAllTextAsync(@"C:\Windows\py.exe"),
};

await Task.WhenAll(tasks);

```
