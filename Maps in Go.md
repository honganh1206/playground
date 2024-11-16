---
tags:
  - "#study"
  - "#review"
  - "#golang"
  - "#programming"
cssclasses:
  - center-images
---
A `nil` map has no keys. and keys cannot be added. We use `make` to initialize a map of a given type


```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

```

**Map literals** are similar to struct literals (initialization at the point of declaration) but keys are required

```go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// You can write it like this
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```


```go
// Mutating maps
func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // 48
	delete(m, "Answer") // Key not in map anymore
	fmt.Println("The value:", m["Answer"]) // default non-existent key to be 0

	v, ok := m["Answer"] // default boolean value to be false
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}


```



```go
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		if _, exist := m[w]; exist {
			m[w]++
		} else {
			m[w] = 1
		}
	}
	return m
}

```


