![Master](https://github.com/sszarek/goterate/actions/workflows/go.yaml/badge.svg)
[![codecov](https://codecov.io/gh/sszarek/goterate/branch/master/graph/badge.svg?token=7VUBJK8QOX)](https://codecov.io/gh/sszarek/goterate)

# Goterate
An naive attempt to implement LINQ like experience of working with collections (arrays, slices and maps) in Go using generics.

## Implemented methods
- `Filter`
- `Take`

# Examples

## Iterate over slice
```go
slice := []int {1,2,3}
iterable := NewSliceIterable(slice)
iterator := iterable.GetIterator()

for iterator.MoveNext() {
    fmt.Printf("%d\n", iterator.GetCurrent())
}
```

The above statement will produce following output:
```
1
2
3
``` 

## Filter slice
```go
slice := []int{-1,-2, -3, 1, 2, 3}

iterable := NewSliceIterable(slice).Filter(func(t int) bool { return t > 0 })
iterator := iterable.GetIterator()

for iterator.MoveNext() {
    fmt.Printf("%d", iterator.GetCurrent())
}
```

The code will create `iterable` which will filter out numbers less than zero during iteration. Note that original slice will not be modified.
Following output will be produced:
```
1
2
3
```