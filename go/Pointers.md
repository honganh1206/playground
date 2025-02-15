---
tags:
  - "#study"
  - "#review"
  - "#golang"
cssclasses:
  - center-images
---
Pointers in Go does NOT have pointer arithmetic

```go
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)	// read the memory address
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```