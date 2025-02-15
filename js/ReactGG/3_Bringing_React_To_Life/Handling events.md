> [!info] Good practices
> 
> - Encapsulate event handlers to their own functions with naming `handle` + name of the event
> - Pass the handler as a **reference** (`handleClick`), not an invocation (`handleClick()`)
> - Create the event handler **inside** the component for **access to any state/props via closures**

```jsx
export default function AlertButton({msg}) {
	const handleClick = () => alert(msg)
	return (
		<button onClick={handleClick}>
			Alert
		</button>
	)
}
```