function Badge(props) {
    return (
        <div className="badge">
            <img alt={props.name} src={props.img} />
            <div>
                <h4>{props.name}</h4>
                <p>@{props.handle}</p>
            </div>
        </div>
    );
}

function App() {
    return (
        <Badge
            name="Tyler McGinnis"
            handle="tylermcginnis"
            img="https://avatars0.githubusercontent.com/u/2933430"
        />
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
