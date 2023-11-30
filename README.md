# lazy

A lazy iterator for lazy programmers with emphasis on clarity, and decent performance.

### Pick Your Source
```go
// {1}
lazy.FromFunc(func() (int, bool) {
    return 1, false
})

// 1, 1, ...
lazy.Generate(func() int {
    return 1
})

lazy.FromValues("a", "b") // {"a", "b"}
lazy.FromSlice([]string{"a", "b", "c"}) // {"a", "b", "c"}

lazy.To(10) // [0, 10)
lazy.From(5) // [5, inf)
lazy.From(1).To(3) // {1, 2}
lazy.From(0).To(10).By(50) // {0, 50}
```

### Chain Operations
Chain operations pointer free, interface free and gluten free:
```go
import "github.com/gavraz/lazy"

it := lazy.Generate(func() int {
    return rand.Int() % 100
}).Filter(func(x int) bool {
    return x%2 == 1
}).Map(func(v int) int {
    return v * v
}).Limit(10)

for v, ok := it.Next(); ok; v, ok = it.Next() {
	fmt.Println(v)
}
// Output: 49 25 1 1 25 1 81 9 25 49
```
Current Operations:
* Filter
* Map
* Paginate
* Limit
* Discard

### Iterate Easily
```go
it = lazy.(...).Easy() 
for it.Next() {
    // it.Value()
}
```
Or, turn into a slice:
```go
s := it.Limit(10).Slice()
fmt.Println(cap(s)) // A best effort is made to allocate the exact size: cap(s) = 10
```