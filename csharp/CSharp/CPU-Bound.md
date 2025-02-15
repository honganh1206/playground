---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---
Use `Task.Run()` to return a task object that *can be awaited* with `await`


```cs
// This method returns IMMEDIATELY once await is met => Non-blocking
async Task<string> CalculateResultAsync(string input)
{
	// A time-consuming job here executed by current thread or runtime pool thread
    string result = await Task.Run(() => CalculateComplexOutput(input));
    return result; // Resume when the time-consuming task ends
}

string CalculateComplexOutput(string input)
{
    StringBuilder sb = new StringBuilder();
    for (int i = input.Length - 1; i >= 0; i--)
    {
        sb.Append(input[i]);
    }

    return sb.ToString();
}

string reversedInput = await CalculateResultAsync("Hello world!");

Console.WriteLine(reversedInput);

```