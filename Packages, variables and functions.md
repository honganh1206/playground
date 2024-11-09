- A name is **exported** if it *begins with a capital letter*
- A function can return **any**  number of results

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

- Go’s return values may be named, and if so *they are treated as variables defined at the top of the function*

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Return the named return values aka Naked return
}

func main() {
	fmt.Println(split(17))
}

```

- `var` statement declares **a list of variables** with the type at the end. `var` could be at package or function level

```go
package main

import "fmt"

var c, python, java bool // package level

func main() {
	var i int // function level
	fmt.Println(i, c, python, java)
}

```

- A variable declaration can include initializers. With initializers, the type can be omitted

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```

- The short assignment `:=` can be used instead of a `var` declaration with implicit type (Function scope only)

```go
package main

import "fmt"

var a int = 1 // this is fine

a := 1 // this will lead to a compile error

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java, a)
}

```

- Some of Go’s basic types like `byte` and `rune` are **aliases** for types like `uint8` and `int32` respectively
- Variables *declared without an explicit initial value* are given their zero/default value

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
}
```

- The expression `T(v)` converts the value `v` to type `T`

```go
i := 42
f := float64(i)
u := uint(f)
```

- When the right-hand side of the declaration is typed, the new variable is of that same type

```go
var i int
j := i // j is an int
```

- Constants cannot be declared using `:=`
- Numeric constants are **high-precision** values. An untyped constant takes the type needed by its context

```cs
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Big)) // This leads to overflow as int can store at max 64-bit integer
}

```