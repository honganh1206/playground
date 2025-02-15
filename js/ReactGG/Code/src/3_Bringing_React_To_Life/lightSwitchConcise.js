function App() {
    const [mode, setMode] = React.useState("dark");

    const handleMode = () => {
        setMode(mode === "dark" ? "light" : "dark");
    };

    return (
        <main className={mode}>
            <button onClick={handleMode}>Activate {mode} mode</button>
        </main>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
