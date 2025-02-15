
## Props

- Props are to components **what arguments are to functions**
- Props is always an **object** and if you pass a prop without a value, that will be `true` in the component (implicit Boolean value)

```jsx
<Pokemon id="bulbasaur" starter />

{ id: "bulbasaur", starter: true}
```

### Access the data

- Use functions to pass props into components

```jsx
function Layout (props) {
	return (
		<div classname="layout">
			<SideBar />
				{props.children}
			<Footer />
		</div>
	)
}

...
<Layout>
	<Pizza/>
</Layout>
```


## React Elements

- An object representation of a DOM node

```json
// Example of an element as an object representation of a DOM node
{
  type: "h1",
  props: {
    children: "Profile",
    className: "header"
  }
}
```

```html
<!-- How the object looks in DOM-->
<h1 class="header">
  Profile
</h1>
```

> [!info] How React transform the object to the DOM node
> 
> The way React works is **it'll keep references to all of the elements** in your application. If one of those references changes, React will then know where exactly in the DOM it needs to update.

- You can also pass in other React components into JSX

```jsx
import { jsx } from "react/jsx-runtime"

function Badge({ name, handle, img }) {
  return jsx("div", {
    className: "badge",
    children: [
      jsx("img", {
        alt: name,
        src: img
      }),
      jsx("div", {
        children: [
          jsx("h4", {
            children: name
          }),
          jsx("p", {
            children: ["@", handle]
          })
        ]
      })
    ]
  });
}

export default function App() {
  return jsx(Badge, {
    name: "Tyler McGinnis",
    handle: "tylermcginnis",
    img: "https://avatars0.githubusercontent.com/u/2933430"
  });
}
```
---

## Quiz


### Props

#### Q4

```jsx
<DatePicker
  eventName={`${event} - ${date}`}
  renderFooter={() => (
    <button onClick={handleClose}>
      Close
    </button>
  )}
  // An object should be passed here => Use double curly brackets {{}}
  settings={
    format: "yy-mm-dd",
    animate: false,
    days: "abbreviate"
  }
/>
```


### Elements & Components

#### Q2: Why does React use an object representation of the DOM?

- Objects are lightweight + easy to create/destroy
- Easily compare the current element with the previously rendered element
- Update the DOM nodes that have changed
- Avoid wiping out + recreating the entire DOM tree

---
