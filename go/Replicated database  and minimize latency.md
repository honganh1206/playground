
```go
// Ask all replicated databases and return the first response to arrive
func Query(conns []Conn, query string) Result {
	ch := make(chan Result, len(conns)) // Buffered
	for _, conn := range conns {
		go func(c Conn) {
			ch <- c.DoQuery(query)
		}(conn)
	}
	return <-ch
}
```