function Counter () {
  const [count, setCount] = React.useState(0)

  const handleClick = () => {
    setCount(1)
    setCount(2)
    setCount(3)
  }

  return (
    <main>
      <h1>{count}</h1>
      <button onClick={handleClick}>
        +
      </button>
    </main>
  )
}


const container = document.getElementById("mydiv");
const root = ReactDOM.createRoot(container);
root.render(<Counter />);
