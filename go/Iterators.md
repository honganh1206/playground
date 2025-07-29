---
id: Iterators
aliases: []
tags: []
---

[Docs](https://pkg.go.dev/iter)

[Article](https://bitfieldconsulting.com/posts/iterators)

A function that "yields" one result at a time instead of computing the whole set of result to return

When we write a `for` loop where the `range` expression is an iterator, _the loop will be executed once for each value returned_ by the iterator.

> Analogy: Think of a short-order cook who makes a dish for each customer on demand, rather than cooking a whole day's worth of meals in advance

## Why are iterators useful?

```go
type Item int

func Items() (items []Item) {
    ... // generate items
    return items
}

// Problem: We have to WAIT until the whole slice is generated
// and we have to allocate memory for the whole slice
// even though we only use one element at a time
for _, v := range Items() {
  fmt.Println("item", v)
}
```

The idea is to have something similar to a slice, but it only _produces one element at a time_

## How iterators look like

```go
// A function with a particular signature
// It takes a function named yield as an argument
// and this function will yield each successive value to the range loop
func (yield func(Item) bool)
```

When the iterator has completed yielding all available items, it returns, signaling "iteration complete"

```go
// iterateItems is an iterator
// When the loop terminates, yield() will return false
for v := range iterateItems {
  ...
}
```

Question: Why do we need to check if the boolean resullt of `yield` is `false` so we can stop? Why does it return a result at all?
Answer: There might be resources waiting to be clean up, and the `false` result signal the "Clean up, we are done."

## Single value iterator `iter.Seq`

```go
// Items() is not an iterator itself
// instead it RETURNS an interator
func Items() iter.Seq[Item] {
  return func(yield func(Item) bool) {
    items := []Item{1, 2, 3}
    for _, v := range items {
      if !yield(v) {
          return
      }
    }
  }
}

// We are ranging over not a function
// but the RESULT of invoking a function
for v := range Items() {
  fmt.Println("item", v)
}
```

## Two-value iterators `iter.Seq2`

What if we want the index too? We can change the signature

```go
// Each time round the loop,
// the yield function will generate a index-Item pair
func (int, Item) bool
```

## Dealing with errors

Some iterators are _infallible_: They will always produce _as many results as they're asked for_

But what if the iterator has to do something to generate each item and _that something could fail?_ We can yield an error instead of an index

```go
func Lines(file string) iter.Seq2[string, error]
```

## Composing iterators

Functions that take iterators. We can pass in a type argument

```go
func PrintAll[V any](seq iter.Seq[V]) {
    for v := range seq {
        fmt.Println(v)
    }
}
```

## When iterators beat [[Channels]]

Looping over a channel shows the same behavior when we loop over an iterator, but for that our program must be _concurrent_, which means dealing with more complexity

And what if the loop exits early? When we stop listening to a channel, whoever is trying to send something to a channel will just block it then, and at that point the sending goroutine will leak (hanging around taking up more and more memory)
