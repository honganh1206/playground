function App() {
    const [mode, setMode] = React.useState("dark");

    const handleDarkMode = () => {
        setMode("dark");
    };
    const handleLightMode = () => {
        setMode("light");
    };

    return (
        <main className={mode}>
            {mode === "light" ? (
                <button onClick={handleDarkMode}>Activate Dark Mode</button>
            ) : (
                <button onClick={handleLightMode}>Activate Light Mode</button>
            )}
        </main>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
