# lazy

A lazy iterator for lazy programmers.

### Build
```go
r := lazy.Build(
    iterator.Generate(func() int {
        return rand.Int() % 10
    }), lazy.Filter(func(x int) bool {
        return x%2 == 0
    }), lazy.Map(func(v int) int {
        return v * v
    }), lazy.Limit[int](3))

for r.Next() {
    fmt.Print(r.Value(), " ")
}

// Output: 4 64 16
```

### Simple iteration on values
```go
for i := iterator.FromValues("a", "b", "c"); i.Next(); {
    fmt.Print(i.Value(), " ")
}

// Output: a b c
```

### Range, paginate and to slice
```go
ten := iterator.To(10) // 0, 1, ..., 9
secondPage := iterator.Paginate(ten, 2, 3) // second page, assuming three elements per page
fmt.Println(iterator.Slice(secondPage))

// Output: [3 4 5]
```


