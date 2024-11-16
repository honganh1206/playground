---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#golang"
cssclasses:
  - center-images
---
## Receivers

Go does not have classes, but you can define *methods on types*

Methods are *functions with a special receiver argument*


```go
type Vertex struct {
	X, Y float64
}

// Defining methods on type Vertex
// v is the receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}


```

We can declare methods on non-struct types or pointers (more common than value receiver)

```go
type MyFloat float64

// Non-struct receiver
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

```


## Methods and pointer indirection

While functions with a pointer argument must take a pointer, methods with pointer receivers can take *either a value or pointer as the receiver*

```go
// Can work with value as the receiver as well
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p) // {60 80} &{96 72}
}

// The reverse direction works as well
// Methods can take either value receivers or pointer receivers
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Funcs must take arguments of a specific type
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```


> [!tip] Choosing a value or pointer receiver?
> There are 2 reasons to use a pointer receiver over a value receiver:
> 1. Method can *directly* modify the value its receiver points to.
> 2. We avoid copying the value on each method call especially when working with a large struct
