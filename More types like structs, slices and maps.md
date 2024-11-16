---
tags:
  - "#study"
  - "#review"
  - "#golang"
  - "#programming"
cssclasses:
  - center-images
---
## [[Pointers]]

## [[Structs]]


## Arrays in Go

In Go, the type `[n]T` is an array of `n` values of type `T`

## [[Slices]]

## Range


```go
func main() {
	// The index and a copy of the element at that index
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	pow := make([]int, 10)
	for i := range pow {
		// Bitwise operation - Shift 1 to the left by i positions
		// == 2^i
		pow[i] = 1 << uint(i) 
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	// Init slice of length dy
	s := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		// Init slices in a slice
		s[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[y][x] = uint8((x+y)/2) + uint8(x^y) + uint8(x*y)
		}
	}
	return s
}

```

## [[Maps in Go]]

## [[Functions in Go]]

