function Avatar({ img, alt }) {
    return <img src={img} alt={alt} />;
}

function Name({ name }) {
    return <h4>{name}</h4>;
}

function Handle({ handle }) {
    return <p>@{handle}</p>;
}

function Badge({ user, style, addFriend }) {
    return (
        <div style={style}>
            <Avatar img={user.name} alt={user.alt} />
            <div>
                <Name name={user.name} />
                <Handle handle={user.handle} />
                <button onClick={addFriend}>Add Friend</button>
            </div>
        </div>
    );
}

function App() {
    return (
        <Badge
            user={{
                name: "Lynn Fisher",
                img: "https://avatars.githubusercontent.com/u/871315",
                handle: "lynnandtonic",
            }}
            style={{
                width: 300,
                margin: "0 auto",
            }}
        />
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
