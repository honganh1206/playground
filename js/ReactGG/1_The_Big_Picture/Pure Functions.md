## Lessons

### Functions

- Functions can be unpredictable due to the **side effects** and **inconsistent outputs**

```js
// A function relying on the state => Side effects
function addTodo(todo) {
	this.todos.push(todo);
}
// A function with inconsistent outputs
function getGitHubProfile(username) {
	return fetch(`https://api.github.com/users/${username}`)
		.then((res) => res.json());
}

const p1 = getGitHubProfile("username");
const p2 = getGitHubProfile("username");

p1 === p2 // false because two different objects
```

### Rules

1: No side effects

2: Keep your outputs consistent


### Signs of a good function

#### Cacheable

```js

let primeCache = {
	1: false
}

// The output will not change if the input stays the same
const isPrime = (n) => {
	// The code is more impure as it depends on the external state i.e., primeCache
	if (typeof primeCache[n] === 'boolean') {
		return primeCache[n];
	}
	// We iterate up to `Math.sqrt(n)` because if a number `n` has a factor (a number or algebraic expression that divides another number or expression evenly—i.e., with no remainder) `f` greater than its square root, then it also has a factor smaller than or equal to its square root.
	for (let i = 2; i <= Math.sqrt(n); i++) {
		if (n % i === 0) {
			primeCache[n] = false;
			return false;
		}
	}
	primeCache[n] = true;
	return true;
}

export default isPrime;
```

#### Testable

- Given the same input, the output stays the same

---

## Quiz


1. This function is NOT pure, as the local storage might change => A different result (game progress) each time we call it 

```js
function getGameProgress(gameId) {

  return localStorage.getItem(gameId);

}
```

