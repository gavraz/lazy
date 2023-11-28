# lazy

A lazy iterator for lazy programmers.

### Concatenate operations and iterate easily
```go
import "github.com/gavraz/lazy"

it := lazy.Generate(func() int {
    return rand.Int() % 100
}).Filter(func(x int) bool {
    return x%2 == 1
}).Map(func(v int) int {
    return v * v
}).Limit(10).Easy()

for it.Next() {
    fmt.Println(it.Value())
}
// Output: 49 25 1 1 25 1 81 9 25 49
```

### Instantiate from values
```go
for i := lazy.FromValues("a", "b", "c").Easy(); i.Next(); {
    fmt.Print(i.Value(), " ")
}
// Output: a b c
```

### Range, paginate and slice
```go
ten := lazy.To(10) // 0, 1, ..., 9
secondPage := ten.Paginate(2, 3)
fmt.Println(secondPage.Slice())
// Output: [3 4 5]
```