// Each Todo component handle its own update/delete event
function Todo({ todo, handleUpdateTodo, handleDeleteTodo }) {
    const [completed, setCompleted] = React.useState(false);
    const [editing, setEditing] = React.useState(false);

    const handleCheckboxClick = () =>
        handleUpdateTodo({
            ...todo,
            completed: !todo.completed,
        });
    const handleEditClick = () => setEditing(!editing);
    const handleUpdateLabel = (event) =>
        handleUpdateTodo({
            ...todo,
            label: event.target.value,
        });
    const handleDeleteClick = () => handleDeleteTodo(todo.id);
    return (
        <label htmlFor={todo.id}>
            <div>
                <input
                    type="checkbox"
                    id={todo.id}
                    checked={todo.completed}
                    onChange={handleCheckboxClick}
                />
                <span />
            </div>
            <span>
                {editing === true ? (
                    <input
                        type="text"
                        value={todo.label}
                        onChange={handleUpdateLabel}
                    />
                ) : (
                    <span>{todo.label}</span>
                )}
            </span>
            <button onClick={handleEditClick}>
                {editing ? "Save" : "Edit"}
            </button>
            <button onClick={handleDeleteClick}>Delete</button>
        </label>
    );
}
// Handle adding new todo
function TodoComposer({ handleAddTodo }) {
    const [label, setLabel] = React.useState("");
    const handleUpdateLabel = (event) => setLabel(event.target.value);

    function createTodo(label) {
        return {
            id: Math.floor(Math.random() * 1000),
            label,
            completed: false,
        };
    }

    const handleAddTodoClick = () => {
        const todo = createTodo(label);
        handleAddTodo(todo);
        setLabel(""); // Reset the content of the add todo input field
    };
    return (
        <li>
            <input
                placeholder="Add a new todo"
                type="text"
                value={label}
                onChange={handleUpdateLabel}
            />
            <button disabled={label.length === 0} onClick={handleAddTodoClick}>
                Add Todo
            </button>
        </li>
    );
}
// Maintain the overall state of the todo list
function TodoList() {
    // State that multiple components depende on
    const [todos, setTodos] = React.useState([
        {
            id: 1,
            label: "Learn React",
            completed: false,
        },
        {
            id: 2,
            label: "Sleep",
            completed: false,
        },
        {
            id: 3,
            label: "Workout",
            completed: false,
        },
    ]);

    const handleUpdateTodo = (updatedTodo) => {
        const newTodos = todos.map((todo) =>
            todo.id === updatedTodo.id ? updatedTodo : todo
        );
        setTodos(newTodos);
    };
    const handleDeleteTodo = (id) => {
        // filter out todos that are not going to be deleted and update the state => Avoid mutating state directly
        const newTodos = todos.filter((todo) => todo.id !== id);
        setTodos(newTodos);
    };
    const handleAddTodo = (newTodo) => {
        const newTodos = [...todos, newTodo];
        setTodos(newTodos);
    };

    return (
        <ul>
            <TodoComposer handleAddTodo={handleAddTodo} />
            {todos.map((todo) => (
                <Todo
                    key={todo.id}
                    todo={todo}
                    handleUpdateTodo={handleUpdateTodo}
                    handleDeleteTodo={handleDeleteTodo}
                />
            ))}
        </ul>
    );
}

function App() {
    return <TodoList />;
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
