
## Hooks

- Special functions that allow you to get help from React
- Must start with `use...`
- Cannot be called insides conditions/loops/nested functions


> [!tip] Lifting state up
> 
> Whenever you have a state that multiple components depend on, you might want to **lift that state up** to the nearest parent component, and pass it down as a prop


> [!tip] Avoid mutating state directly
> 
> There are functions that help us avoid mutating the state directly like `filter()`, `map()` and the spread `...` operator


```js
const handleUpdateTodo = (updatedTodo) => {
	const newTodos = todos.map((todo) =>
		todo.id === updatedTodo.id ? updatedTodo : todo
	);
	setTodos(newTodos);
};

const handleDeleteTodo = (id) => {
	// filter out todos that are not going to be deleted and update the state => Avoid mutating state directly
	const newTodos = todos.filter((todo) => todo.id !== id);
	setTodos(newTodos);
};

const handleAddTodo = (newTodo) => {
	const newTodos = [...todos, newTodo];
	setTodos(newTodos);
};
```
## Quiz

```jsx
// spread state into a new state
const [state, setState] = React.useState({
	loading: true,
	authed: false,
	repos: []
});

const handleAuthComplete = () => {
	setState({...state, authed: true})
}
```