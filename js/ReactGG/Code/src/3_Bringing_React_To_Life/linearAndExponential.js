function App() {
    const [linear, setLinear] = React.useState(0);
    const [exponential, setExponential] = React.useState(1);

    const handleClick = () => {
        // React will render the state to the UI first
        // UI will show Linear: 1 and Exponential: 2
        setLinear(linear + 1);
        setExponential(exponential * 2);

        // The log will show the current state after we click the button
        console.log({ linear, exponential });
    };

    return (
        <main>
            <p>Linear: {linear}</p>
            <p>Exponential: {exponential}</p>
            <button onClick={handleClick}>Do Math</button>
        </main>
    );
}

const container = document.getElementById("root");
const root = ReactDOM.createRoot(container);
root.render(
    React.createElement(React.StrictMode, null, React.createElement(App))
);
