function Greeting({ name }) {
    const [index, setIndex] = React.useState(0);

    const greetings = ["Hello", "Hola", "Bonjour"];

    const handleClick = () => {
        // If we reach the end of the array, take the 1st element
        // Otherwise, continue incrementing the index by 1
        const nextIndex = index === greetings.length - 1 ? 0 : index + 1;
        setIndex(nextIndex);
    };

    return (
        <main>
            <h1>
                {greetings[index]}, {name}
            </h1>
            <button onClick={handleClick}>Next Greeting</button>
            <span role="img" aria-label="hand waving">
                {console.log("rendering wave, if you see this twice, that means strict mode is on")}
            </span>
        </main>
    );
}

const container = document.getElementById("root");
const root = ReactDOM.createRoot(container);
root.render(
  // When using React.createElement, the second argument should be an object containing the props.
    React.createElement(
        React.StrictMode,
        null,
        React.createElement(Greeting, { name: "Hong Anh Pham" })
    )
);
