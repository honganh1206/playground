
## Lesson

### Overview

- Imperative concerns the HOW, while declarative concerns the WHAT

### Differences

- Describe HOW we want to achieve a certain task => Think like a computer >< Think like a human in Declarative
- Declarative code can be context-dependent => Concerning the WHAT (ultimate goal) and ignoring the HOW (steps taken to achieve the goal) allows the code to be **reusable** in different programs

```js
// Imperative
function double(arr) {
	let results = [];
	for (let i = 0; i < arr.length; i++) {
		results.push(arr[i] * 2);
	}
	return results;
}
// Declarative
function double(arr) {
	return arr.map((item) => item * 2);
}
```

> [!info]
> 
> Many (if not all) **declarative** APIs have some sort of **imperative** implementation as we are abstracting imperative code behind a declarative APIs

---

## Quiz


1. What makes React declarative
	- Describes the UI based on state
	- Abstract imperative code