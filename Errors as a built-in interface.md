Tags: #review #programming #golang 

An interesting thing about Go is that it deals with errors as values with the `error` type.

The `error` type is also a built-in interface similar to Stringer


```go
type error interface {
    Error() string
}
```

Functions often returns an `error`-typed valued and our code should handle errors by testing whether it is equal to `nil`


```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
