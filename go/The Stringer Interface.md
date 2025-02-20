---
id: The Stringer Interface
aliases: []
tags: []
---

Tags: #review #programming #golang

```go
type Stringer interface {
    String() string
}

```

The `Stringer` interface is implemented by the `fmt` package as _a type that can describe itself as a string_

Implementing the `Stringer` interface allows values to be _automatically converted into strings_ when using with `fmt.Print()` or `fmt.Printf()`. It is a common pattern in Go for _providing string representations for custom types_

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z) // Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
}
```
