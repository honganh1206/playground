---
id: Flow control statements
aliases: []
tags:
  - #study
  - #review
  - #programming
  - #golang
cssclasses:
  - center-images
---

- In Go, `while` is `for` but drop the semicolons

```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

- The forever loop omits the loop condition

```go
func main() {
	for {
	}
}

```

- **`if` with a short statement**: Execute some code before evaluating the condition

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // The math.Pow will be executed BEFORE evaluating the condition
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

- Variables declared inside an `if`’s short statement is _also available_ inside the `else` block

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
```

- **Defer**: Delay the execution of a function until the surrounding functions returns. Mostly used to perform clean-up actions

```go
func main() {
	// These lines will be executed when the main() function is about to finish
	// But in STACK ORDER
	defer fmt.Println("bye")
	defer fmt.Println("world")

	// The below lines are executed first
	fmt.Println("hello")
	fmt.Println("hi")
}


// output: hello hi world bye
```

[More on `defer`](https://go.dev/blog/defer-panic-and-recover)

[A good blog on why sometimes using `defer` keyword could be a trap](https://victoriametrics.com/blog/defer-in-go/#:~:text=In%20Go%2C%20defer%20is%20a,until%20the%20surrounding%20function%20finishes.&text=In%20this%20snippet%2C%20the%20defer,end%20of%20the%20main%20function.)
