function App({ characterLimit = 20 }) {
    const [inputValue, setInputValue] = React.useState("");
    const remainingChars = characterLimit - inputValue.length;
    const limitExceeded = remainingChars < 0;

    const handleChange = (e) => {
        setInputValue(e.target.value);
    };

    const handleSubmit = (e) => {
        e.preventDefault();

        // if input length is too long
        // alert "The input exceeds the character limit. Please shorten your text."
        // else
        // alert "Thanks for your submission"
        if (limitExceeded) {
            alert(
                "The input exceeds the character limit. Please shorten your text"
            );
        } else {
            alert("Thanks for your submission");
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label htmlFor="limited-text-input">Limited Text Input:</label>
                <span className={limitExceeded ? "error" : "no-error"}>
                    Characters remaining: {remainingChars}
                </span>
            </div>
            <input
                type="text"
                placeholder="Enter some text"
                id="limited-text-input"
                value={inputValue}
                onChange={handleChange}
            />

            <button type="submit" className="primary">
                Submit
            </button>
        </form>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
