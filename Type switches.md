Tags: #review #programming #golang 

A construct that allows several type assertions in series.

A type switch is similar to a regular switch statement but the cases are types + we compare the values against the types of the value held by the given interface value.


```go
func do(i interface{}) {
	switch v := i.(type) { // Of either int type or string type and holds the value of i
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

```
