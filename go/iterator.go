package main

import (
	"fmt"
	"iter"
)

type Item int

// This function is an iterator itself
func Items() iter.Seq[Item] {
	// The function yield working as the argument
	// will yield each successive Item for the range loop
	return func(yield func(Item) bool) {
		items := []Item{1, 2, 3}
		for _, v := range items {
			if !yield(v) {
				// Signal complete yielding all items if return false
				return
			}
		}
	}
}

func Items2() iter.Seq2[int, Item] {
	return func(yield func(int, Item) bool) {
		items := []Item{1, 2, 3}
		for i, v := range items {
			if !yield(i, v) {
				// Signal complete yielding all items if return false
				return
			}
		}
	}
}

func PrintAll[V any](seq iter.Seq[V]) {
	for v := range seq {
		fmt.Println(v)
	}
}

func main() {
	// for v := range Items() {
	// 	fmt.Println("item", v)
	// }
	// for i, v := range Items2() {
	// 	fmt.Println("item", i, "is", v)
	// }
	PrintAll(Items())

}
