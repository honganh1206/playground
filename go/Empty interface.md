Tags: #review #golang #programming 


An *empty interface* specifies zero methods. It can hold values of **any** type and is used to handle values of unknown type.

Think of this as a super easy-going person? This person accepts value of any type because **every** type satisfies this person.


```go

func main() {
	var i interface{}
	describe(i) // 	(<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

