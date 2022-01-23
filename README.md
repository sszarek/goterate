Goterate
========
An naive attempt to implement LINQ like experience of working with collections (arrays, slices and maps) in Go using generics.

Examples
========

Iterate over slice
------------------
```go
slice := []int {1,2,3,4}
iterable := NewSliceIterable(slice)
iterator := iterable.GetIterator()

for iterator.MoveNext() {
    fmt.Printf("%d", iterator.GetCurrent())
}

```