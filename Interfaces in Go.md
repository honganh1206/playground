---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#golang"
cssclasses:
  - center-images
---
An **interface** type contains a set of *method signatures*. A value of interface type can hold *any value that implements those method*.

We implement an interface by *implementing its methods*. There is no explicit declaration of intent or “implements” keyword.

```go
type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())
}

// Implementation of Abser of value receiver type
// We implicitly implement the Abser interface here
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Implementation of Abser of pointer receiver type
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

## Interface values

Interface values = A tuple of a value and a concrete type


```go
var a Animal = Dog{"Buddy"} // The type here is `Dog` and the value is `Dog{"Buddy"}`

```


If the concrete value inside the interface is nil then *the method will be called with a nil receiver*


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

	var t *T
	i = t // Nil underlying value here
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
