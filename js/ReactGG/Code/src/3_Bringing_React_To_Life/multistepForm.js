const initialFormData = {
    name: "",
    email: "",
    address: "",
    city: "",
    zipcode: "",
};
function App() {
    const [currentStep, setCurrentStep] = React.useState(1);
    const [formData, setFormData] = React.useState(initialFormData);

    const handleChange = (event) => {
        setFormData({
            ...formData,
            [event.target.name]: event.target.value,
        });
    };

    const handleNextStep = () => {
        setCurrentStep(currentStep + 1);
    };

    const handlePreviousStep = () => {
        setCurrentStep(currentStep - 1);
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        alert("Done");
        setCurrentStep(1);
        setFormData(initialFormData);
    };

    if (currentStep === 1) {
        return (
            <form onSubmit={handleSubmit}>
                <h2>Personal Information </h2>
                <div>
                    <label>Step {currentStep} of 3</label>
                    <progress value={currentStep} max={3}></progress>
                </div>
                <div>
                    <label htmlFor="name">Name</label>
                    <input
                        required
                        name="name"
                        id="name"
                        placeholder="Enter your name"
                        value={formData.name}
                        onChange={handleChange}
                    ></input>
                </div>
                <div>
                    <label htmlFor="email">Email</label>
                    <input
                        required
                        name="email"
                        id="email"
                        placeholder="Enter your email"
                        value={formData.email}
                        onChange={handleChange}
                    ></input>
                </div>
                <button
                    type="submit"
                    className="secondary"
                    onClick={handleNextStep}
                >
                    Next
                </button>
            </form>
        );
    } else if (currentStep === 2) {
        return (
            <form onSubmit={handleSubmit}>
                <h2>Personal Information </h2>
                <div>
                    <label>Step {currentStep} of 3</label>
                    <progress value={currentStep} max={3}></progress>
                </div>
                <div>
                    <label htmlFor="address">Address</label>
                    <input
                        required
                        name="address"
                        id="address"
                        placeholder="Enter your address"
                        value={formData.address}
                        onChange={handleChange}
                    ></input>
                </div>
                <div>
                    <label htmlFor="city">City</label>
                    <input
                        required
                        name="city"
                        id="city"
                        placeholder="Enter your city"
                        value={formData.city}
                        onChange={handleChange}
                    ></input>
                </div>
                <div>
                    <label htmlFor="zipcode">Zipcode</label>
                    <input
                        required
                        name="zipcode"
                        id="zipcode"
                        placeholder="Enter your zipcode"
                        value={formData.zipcode}
                        onChange={handleChange}
                    ></input>
                </div>
                <button
                    type="submit"
                    className="secondary"
                    onClick={handlePreviousStep}
                >
                    Previous
                </button>
                <button
                    type="submit"
                    className="secondary"
                    onClick={handleNextStep}
                >
                    Next
                </button>
            </form>
        );
    } else if (currentStep === 3) {
        return (
            <form onSubmit={handleSubmit}>
                <h2>Personal Information </h2>
                <div>
                    <label>Step {currentStep} of 3</label>
                    <progress value={currentStep} max={3}></progress>
                </div>
                <div>
                    <p>{formData.name}</p>
                    <p>{formData.email}</p>
                    <p>{formData.address}</p>
                    <p>{formData.city}</p>
                    <p>{formData.zipcode}</p>
                </div>
                <button
                    type="submit"
                    className="secondary"
                    onClick={handlePreviousStep}
                >
                    Previous
                </button>
                <button
                    type="submit"
                    className="primary"
                    onClick={handleSubmit}
                >
                    Submit
                </button>
            </form>
        );
    }
}

const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<App />);
