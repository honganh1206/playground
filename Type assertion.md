Tags: #review #golang #programming 

We use type assertion to *access* an interface's value underlying concrete value.


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


