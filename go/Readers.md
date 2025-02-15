Tags: #review #golang #programming 

The `io.Reader` interface represents **the read end** of a stream of data. This interface has a `Read()` method


```go
// Receive byte slice with data and return number of bytes. Return io.EOF when streams end
func (T) Read(b []byte) (n int, err error)

// Sample implementation
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	// read 8 bytes each iteration
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
/*
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"

n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"

n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = "" 

*/

// This method satisfies the contract of the Reader interface by filling the param slice with letter A
func (r MyReader) Read(b []byte) (n int, err error) {
	// Crucial to populate the b slice
	for i := range b {
		b[i] = 'A'
	}
	
	return len(b), nil
}
```

 - A good practice: An `io.Reader` that wraps another `io.Reader` and modifies the stream in the same way.

## Exercise

```go
// Rot13 reader implementation
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	// Read from the underlying reader
	n, err = r.r.Read(b)
	if n > 0 {
		// Apply ROT13 transformation
		for i := 0; i < n; i++ {
			b[i] = rot13(b[i])
		}
	}
	return n, err
}

// rot13 applies the ROT13 cipher to a single byte
func rot13(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return 'A' + (c-'A'+13)%26
	} else if c >= 'a' && c <= 'z' {
		return 'a' + (c-'a'+13)%26
	}
	return c
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // "You cracked the code!" in ROT13
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // Decodes and prints: "You cracked the code!"
}


```
