
$$v = f(s)$$

With $v$ as View, $f$ as Function and $s$ as state


## Rendering

### What is rendering?

- Rendering in React = React calls your component to update the view
- React creates a **snapshot** of your component at that particular moment of time => **Props/State/Event handlers/UI description** are all captured in the snapshot

```jsx
import { createRoot } from "react-dom/client"
import App from "./App"

const rootElement = document.getElementById("root") // HTML element to mount your React app to
const root = createRoot(rootElement)

root.render(<App/>) // React element as the starting point
```

### When does React re-render?

- React will only re-render when the **state of the component changes**
#### How? 

1. When an **event handler** is invoked, **that event handler gains to the props/state** at the moment the snapshot was created
2. If the handler contains the invocation of `useState()` and React sees that the state has changed => React will triggers a re-render of the component, creating a new snapshot

![[Pasted image 20240420160934.png]]

> [!info] Batching
> 
> Even if the updater function `setCount()` inside the event handler function is invoked multiple times, **React will only re-render one** as it only keeps the result of the **last invocation**
> 

```jsx
const handleClick = () => {
  setCount(0 + 1) // Only this works per click
  setCount(0 + 1)
  setCount(0 + 1) // Keep the result of this
}

```



> [!tip]
> 
> React is **good at rendering** and components are not always pure.


> [!todo] Re-rendering child components
> 
> Whenever state changes, **React will re-render the component that owns that state and all of its child components** – regardless of whether or not those child components accept any props.

![[Pasted image 20240420170044.png]]




### `StrictMode`

- Whenever you have `StrictMode` enabled, React will re-render your components **twice** per click

## Quiz

Q5

```jsx
import * as React from "react"

function LightSwitch() {
  const [state, setState] = React.useState("on")

  const handleClick = () => {
    setState(state === "on" ? "off" : "on")
    alert(`The switch is ${state}`) // The state is On when the button is clicked the 1st time
  }

  return (
    <button onClick={handleClick}>
      {state}
    </button>
  )

```

Q8

```jsx
const handleClick = () => {
  setCount(10)
  setCount((c) => c + 30)
  setCount(70)
  setCount((c) => c + 100) // Get the latest state (70) then increment by 100 => 170
}
```