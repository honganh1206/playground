Tags: #review #programming #golang 


```go
type Stringer interface {
    String() string
}

```

The `Stringer` interface is implemented by the `fmt` package as *a type that can describe itself as a string*


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

Exercise


```go
import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	parts := make([]string, len(ip))
	for index, byteVal := range ip {
		parts[index] = fmt.Sprintf("%d", byteVal)
	}
	
	return strings.Join(parts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


```
