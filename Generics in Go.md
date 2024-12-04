Tags: #review #golang #programming 

Go functions can work on multiple types with **type parameters**

```go
// Type parameter appear between the []
// s is a slice of type T and x is the value of the same type
// comparable keyword is a constraint to use the == and != on values of the type
func Index[T comparable](s []T, x T) int
```

## Generic types

Go also supports generic types. A type can be parameterized with a type parameter to implement generic data structures


```go
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
```
