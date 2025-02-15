function App() {
    // event handlers inside the component
    const handleChange = (event) => {
        if (event.target.value.length > 10) {
            alert("Character limit exceeded");
        }
    };
    return (
        <section>
            <h1>Character Limit</h1>
            <input
                onChange={handleChange}
                placeholder="Enter some text"
            ></input>
        </section>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
