function App() {
    const [count, setCount] = React.useState(0);
    const handleIncrement = () => {
        setCount(count + 1);
    };
    const handleDecrement = () => {
        setCount(count - 1);
    };
    return (
        <main>
            <span>{count}</span>
            <div>
                <button onClick={handleDecrement}>-</button>
                <button onClick={handleIncrement}>+</button>
            </div>
        </main>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
