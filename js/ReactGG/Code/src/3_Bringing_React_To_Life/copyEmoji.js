const emojis = [
    {
        id: "smile",
        emoji: ":)",
    },
    {
        id: "cry",
        emoji: "T.T",
    },
    {
        id: "frown",
        emoji: ":(",
    },
];

const selectRandomElement = (arr) => {
    return arr[Math.floor(Math.random() * arr.length)];
};

function App() {
    const selected = selectRandomElement(emojis);

    // event handlers inside the component
    const handleCopy = (e) => {
        if (e.target.innerText !== selected.emoji) {
            alert("Wrong emoji");
        } else if (e.timeStamp > 5000) {
            alert("Too slow.Try again");
        } else {
            alert("Winner!");
        }
    };

    const handleReset = () => {
        window.location.reload();
    };

    return (
        <div>
            <h1>{selected.emoji}</h1>
            <p>Copy this emoji</p>
            <ul>
                {emojis.map(({ id, emoji }) => {
                    return (
                        <li onCopy={handleCopy} key={id}>
                            {emoji}
                        </li>
                    );
                })}
            </ul>
            <button onClick={handleReset}>Reset</button>
        </div>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
