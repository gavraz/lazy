# lazy

### Build
```go
r := lazy.Build(
    iterator.Generate(func() int {
        return rand.Int() % 10
    }), lazy.Filter(func(x int) bool {
        return x%2 == 0
    }), lazy.Map(func(v int) int {
        return v * v
    }), lazy.Limit[int](3),
)
```
```go
for r.Next() {
    fmt.Print(r.Value(), ", ")
}
```
```
Output: 4, 64, 16,
```

### Ranges
```go
for i := iterator.To(10); i.Next(); {
    fmt.Print(i.Value(), " ")
}
```
```
Output: 0 1 2 3 4 5 6 7 8 9
```
