package main

func fanIn(c1, c2 <-chan string) <-chan string {
	merged := make(chan string)

	go func() {
		defer close(merged)
		for {
			select {
			case msg1, ok := <-c1:
				if !ok {
					c1 = nil
					continue
				}
			}
		}
	}()
}

func main() {
}
