+++
title = "Go"
template = "level.html"
description = "data structures and algorithms in Go programming languages"
[extra]
    keywords = "go, golang, explain, tutorial, data structures, algorithms, interview"
+++

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
a := [5]int{5, 4, 5, 2, 1}
b := [...]int{5, 4, 5, 2, 1} 
c := b // creates copy of array a (not a reference!)

s := [...]string{"Java", "Python", "Go"}
```


<div id="dynamic-array" />

## Slice - Dynamic array
create slice:
```go
var i []int

i := []int{5, 4, 5, 2, 1}

i := make([]int, 5) // len(i)=5, cap(i)=5

i := make([]int, 0, 5) // len(i)=0, cap(i)=5

```

append to slice:
```go
var i []int
i = append(i, 1)
```

join slices
```go
a := []int{1, 2}
b := []int{2, 4}
a = append(a, b...) // a - [1, 2, 3, 4]
```
slices can overlap underlying arrays:
```go
a := []int{1, 2, 3, 4}
b := a[2:]

b[0] = 300 // referencing to underlying array of a (a also updated by changes in slice b)
fmt.Printf("a: %v, len: %v cap: %v \n", a, len(a), cap(a)) // a: [1 2 300 4], len: 4 cap: 4
fmt.Printf("b: %v, len: %v cap: %v \n", b, len(b), cap(b)) // b: [300 4], len: 2 cap: 2

b = append(b, 500, 600, 700)
b[1] = 400 // slice b is not referencing to underlying array of slice a anymore (doesn't change slice a, even it was overlapping before)

fmt.Println("after append and b modifications:")
fmt.Printf("a: %v, len: %v cap: %v \n", a, len(a), cap(a)) // a: [1 2 300 4], len: 4 cap: 4
fmt.Printf("b: %v, len: %v cap: %v \n", b, len(b), cap(b)) // b: [300 400 500 600 700], len: 5 cap: 6
```

More info about slices:

- [a tour to go slices playground](https://tour.golang.org/moretypes/7)
- [slice internals](https://blog.golang.org/go-slices-usage-and-internals)
- [how append() work inside](https://blog.golang.org/slices)
- [slice tricks](https://github.com/golang/go/wiki/SliceTricks)
- [slice source code](https://golang.org/src/runtime/slice.go#L11)

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
