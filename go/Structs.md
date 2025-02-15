---
tags:
  - "#study"
  - "#review"
  - "#golang"
  - "#programming"
cssclasses:
  - center-images
---

Collections of fields

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

Structs fields can be accessed through a struct pointer


```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	(*p).X = 1e9 // p.X also works 
	fmt.Println(v)
}

```

**Struct Literals** allow us to *initialize a struct at the point of declaration*. This is written using a struct type followed by curly braces `{}`


```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

```