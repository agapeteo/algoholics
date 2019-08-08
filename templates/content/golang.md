# Go

### array
```go
i := [5]int{5, 4, 5, 2, 1}
s := [3]string{"Java", "Python", "Go"}
```

### dynamic array - slice
```go
var i []int
i := []int{5, 4, 5, 2, 1}
i := make([]int, 5) // len(i)=5
i := make([]int, 0, 5) // len(i)=0, cap(i)=5
i = append(i, 6) // adding new elemet
```

### associative array - map
```go
m = make(map[string]int)

v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

