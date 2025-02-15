const initialFormData = {
    name: "",
    email: "",
    address: "",
    city: "",
    zipcode: "",
};
function App() {
    // Persist formFields with React state
    const [formFields, setFormFields] = React.useState([]);

    const handleAddFormField = (e) => {
        e.preventDefault(); // Ensure the form is not submitted the server by default
        const formData = new FormData(e.target);

        const newField = {
            id: new Date().getTime(),
            type: formData.get("type"),
            label: formData.get("label"),
            placeholder: formData.get("placeholder"),
            required: formData.get("required"),
            value: "",
        };

        setFormFields([...formFields, newField]);
        e.target.reset();
    };

    const handleUpdateFormField = (id, updatedField) => {
        const updatedFormField = formFields.map((field) =>
            field.id === id ? { ...field, ...updatedField } : field
        );
        setFormFields(updatedFormField);
    };

    const handleDeleteFormField = (id) => {
        const remainingFormFields = formFields.filter(
            (field) => field.id !== id
        );
        setFormFields(remainingFormFields);
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        alert(JSON.stringify(formFields, null, 2));
    };

    return (
        <div>
            <h1>Form Builder</h1>
            <form id="form-builder" onSubmit={handleAddFormField}>
                <fieldset>
                    <legend>Add a field</legend>
                    <label htmlFor="type">Field Type</label>
                    <select name="type" id="type">
                        <option value="text">Text</option>
                        <option value="number">Number</option>
                        <option value="email">Email</option>
                        <option value="password">Password</option>
                    </select>
                    <div>
                        <label htmlFor="required">Required</label>
                        <input type="checkbox" name="required" id="required" />
                    </div>
                    <label htmlFor="label">Label</label>
                    <input
                        required
                        type="text"
                        name="label"
                        id="label"
                        placeholder="Enter a label"
                    />
                    <label htmlFor="placeholder">Placeholder</label>
                    <input
                        required
                        type="text"
                        name="placeholder"
                        id="placeholder"
                        placeholder="Enter a placeholder"
                    />
                    <button type="submit" className="secondary">
                        Add Form Field
                    </button>
                </fieldset>
            </form>
            <form id="form-fields" submit={handleSubmit}>
                <fieldset>
                    <ul>
                        {formFields.map((field) => (
                            <li key={field.id}>
                                <label htmlFor={`input-${field.id}`}>
                                    {field.label}
                                </label>
                                <input
                                    id={`input-${field.id}`}
                                    required={field.required}
                                    placeholder={field.placeholder}
                                    type={field.type}
                                    value={field.value}
                                    onChange={(event) =>
                                        handleUpdateFormField(field.id, {
                                            value: event.target.value,
                                        })
                                    }
                                />
                                <button
                                    type="button"
                                    className="secondary"
                                    onClick={() =>
                                        handleDeleteFormField(field.id)
                                    }
                                >
                                    Delete
                                </button>
                            </li>
                        ))}
                    </ul>
                    <button type="submit" className="primary">
                        Submit
                    </button>
                </fieldset>
            </form>
        </div>
    );
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
