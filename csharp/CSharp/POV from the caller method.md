---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---
If the caller method is `async`:


```cs
async Task<string> CallerAsync() {
	// Wait for the time-consuming task to finish
	string stringResult = await DownloadDataAsync();
	// Or gather the task -> do some work -> await for the task later
   var taskDownloadData = DownloadDataAsync();
   string stringResult = await taskDownloadData;
   return stringResult;
}

```

The caller does not have to be `async`:


```cs
Task<string> Caller() {
	// We return the task to the caller of this method
	// And the caller of this method will await it
   var taskDownloadData = DownloadDataAsync();
   return taskDownloadData;
}
```
