function VibeCheck() {
    const [status, setStatus] = React.useState("clean");

    const handleClick = () => {
        setStatus(status == "dirty" ? "clean" : "dirty");
        alert(status);
    };

    return <button onClick={handleClick}>{status}</button>;
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<VibeCheck />);
