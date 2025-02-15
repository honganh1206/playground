Tags: #review #golang #programming 

Interface values = A tuple of a value and a  type, both are dynamic.


```go
var a Animal = Dog{"Buddy"} // The type here is `Dog` and the value is `Dog{"Buddy"}`

```


If the concrete value inside the interface is nil then *the method will be called with a nil receiver*. Note that an interface value that holds a nil concrete value is non-nil it-self, as an interface is only nil when *both its dynamic type and dynamic value are nil*


```go
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T  // Nil pointer
	i = t // Assign the nil pointer to the interface + dynamic type *T is set
	describe(i) // (<nil>, *main.T)
	i.M() // <nil>

	i = &T{"hello"} 
	describe(i) // (&{hello}, *main.T)
	i.M() // hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

An example of the nil interface values:


```go
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // panic here
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```