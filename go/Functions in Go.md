---
id: Functions in Go
aliases: []
tags:
  -  #study
  -  #review
  -  #programming
  -  #golang
cssclasses:
  - center-images
---

Functions are values and we can pass them in a similar way we pass values

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4) // default param values for func param
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	fmt.Println(compute(hypot)) // 5
	fmt.Println(compute(math.Pow)) // 81
}

```

Go functions may be **closures** (function value outside of function body)

```go
func adder() func(int) int {
	sum := 0 // Closure
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/*
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
*/


```

```go
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		res := a // Save the current number to return
		a, b = b, a+b
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

```

## Variadic Function

```go
func someFunction(args ...interface{}) {
    // args is a slice of interface{} values
    // You can pass any number of arguments of any type
}

// Usage
someFunction(1, "hello", true)
```

In this case, it allows the function to accept any number of arguments of any type.
