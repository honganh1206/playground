Tags: #review #golang #programming 

We use type assertion to *access* an interface's value underlying concrete value. Type assertion in Go is quite similar to the usual type casting, but with some differences:

1. Type assertions can **only** be used with interface types, while type casting can work with non-interface types.
2. Type assertions provide the return `ok` value to check if the conversion is possible.

This is useful when we need to access methods/fields that are specific to a concrete type e.g., `string` but you are working with a value through an interface `interface{}`

```go
func main() {
	var i interface{} = "hello"

	s := i.(string) // Assign the underlying value of string concrete type
	fmt.Println(s) // s holds the underlying value of "hello" here

	s, ok := i.(string) // Test if an interface value holds a concrete type
	fmt.Println(s, ok) // hello tru

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic
	fmt.Println(f)
}

```

But why? Interfaces in Go are implemented *implicitly*, so we need a safe way to convert back from an interface to a concrete type
