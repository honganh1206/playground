
## Components

- React components == JS functions


> [!important] Single Responsibility Principle
> Components/Functions should do **just one thing**


## JSX

- Write HTML-ish looking syntax inside your components

> [!note]
> 
> You cannot render adjacent parent elements in React. You can only return 1 element at a time.

```jsx
import * as React from 'react'
export default function Authors() {
	return (
		<!-- Fragment or <> only to generate adjacent elements -->
		<React.Fragment> 
			<h2>Authors</h2>
			<ul>
				<li>Bob</li>
				<li>Tim</li>
			</ul>
		</React.Fragment>
	)
}
```

- Elements must be self-closing 

```html
<input name="term" type="text"/>
```

- Components must be capitalized `<User/>

- Making a list in JSX

```jsx
export default function Authors() {
	return (
		<ul id="tweets">
			{tweets.map(tweet) => (
				// Unique key is a must
				<li key={tweet.id}>
					{tweet.text}
				</li>
			)}
		</ul>
	)
}
```

---

## Quiz

```jsx
import Authors from "./Authors"

export default function About () {

// React components is NOT a pure function here
  localStorage.setItem('viewed_about', true)

  return <main>

    <h1>About Us</h1>

    <p>We write JavaScript and words about JavaScript.</p>

    <Authors />

  </main>

}
```

> [!important]
> 
> ["_React’s rendering process must always be pure._"](https://react.dev/learn/keeping-components-pure#side-effects-unintended-consequences)


- To render no UI => The function should return `null` or `false`


---

## Exercises

### Badge variable


```jsx
import React from 'react';
function Badge() {
  const name = "Tyler McGinnis";
  const handle = "tylermcginnis";
  const img = "https://avatars0.githubusercontent.com/u/2933430";

  return (
    <div className="badge">
      <img alt={handle} src={img} />
      <div>
        <h4>{name}</h4>
        <p>@{handle}</p>
      </div>
    </div>
  );
}

export function App(props) {
  return (
    <Badge/>
  );
}
```

### Adjacent elements

```jsx
import React from 'react';
function Layout() {
  return (
    <>
      <header>Header</header>
      <main>Main</main>
      <aside>Aside</aside>
      <footer>Footer</footer>
    </>
  );
}

export function App(props) {
  return (
    <Layout/>
  );
}
```


### Conditional rendering

```jsx
import React from 'react';

function LactoseIntolerant() {
  return (
    <h1>
      <span role="img" aria-label="broccoli and strawberry">
        🥦🍓
      </span>
    </h1>
  );
}

function LactoseTolerant() {
  return (
    <h1>
      <span role="img" aria-label="milk and cheese">
        🥛🧀
      </span>
    </h1>
  );
}


export function App(props) {
  const isLactoseTolerant = true;
  return (
    <div>
      {isLactoseTolerant ? <LactoseTolerant/> : <LactoseIntolerant/>}
    </div>);
}
```


### Rendering lists

```jsx
import React from 'react';

function List() {
  const friends = [
    { id: 893, name: 'Lynn' },
    { id: 871, name: 'Alex' },
    { id: 982, name: 'Ben' },
    { id: 61, name: 'Mikenzi' },
  ];

  return (
    <ul id='friends'>
      {friends.map(fr => (
        <li key={fr.id}>{fr.name}</li>
      ))}
    </ul>
  );
}

export function App(props) {
  return <List />;
}

```

### Rendering lists no keys

```jsx
import React from 'react';

function List() {
  const friends = ['Ben', 'Lynn', 'Alex'];

  return (
    <ul>
      {friends.map((fr, index) => (
        <li key={index}>{fr}</li>
      ))}
    </ul>
  );
}

export function App(props) {
  return <List />;
}

```