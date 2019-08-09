# Go

## Contents
---

- [Array](#array)
- [Slice - dynamic array](#dynamic-array)
- [Map - associative array](#map)
- [Set](#set)
- [Working with files](#files)

---



<div id="array" />

## Array
```go
i := [5]int{5, 4, 5, 2, 1}
s := [3]string{"Java", "Python", "Go"}
```


<div id="dynamic-array" />

## Slice - Dynamic array
```go
var i []int
i := []int{5, 4, 5, 2, 1}
i := make([]int, 5) // len(i)=5
i := make([]int, 0, 5) // len(i)=0, cap(i)=5
i = append(i, 6) // adding new elemet
```


<div id="map"/>

## Map - associative array 
```go
m = make(map[string]int)

v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

<div id="set"/>

## Set
There is no build-in `Set` data structure in `Golang`. But keys in `map` are used as `set`

create:

```go
m := make(map[int]struct{})
```

set value:
```go
m[someValue] = struct{}() // empty struct will not consume memory
```

check if exist:
```go
_, ok := m[someValue]
return ok
```

<div id="files" />

## Files

check if file exist: 
```go
f := "/path/to/file"
if _, err := os.Stat(f); os.IsNotExist(err) {  
    log.Fatalf("%s does not exist", f)
}
```

read file content
```go
contentBytes, err := ioutil.ReadFile(filePath)
if err != nil {
    log.Fatalf ("can't read file %s error: %v", filePath, err)
}
```
