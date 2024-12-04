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
// A type that implements the Abser interface
type MyFloat float64
// A struct that implements the Abser interface
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

[[Interface values]]

[[Empty interface]]

[[Type assertion]]

[[Type switches]]

[[The Stringer Interface]]

[[Errors as a built-in interface]]

[[Readers]]

[[Images]]


