+++
title = "Copy Array"
template = "level.html"
description = "Copy Array"
[extra]
    keywords = "copy array, clone, array, memory, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Copy array

## Contents

---

- [Copy array description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Copy array

Complexity `O(N)` - requires iterating through all elements in collection



<div id="practice"/>

## In practice
Used very often. Dynamic array creates new array and copied data into new array for growing and shrinking


<div id="java"/>

## Java
There few way to copy array in Java:

- `Object.clone()` - to create new same copy of array 
- `Arrays.copyOf()`  
- `Arrays.copyOfRange()`
- `System.arraycopy()` 
- using streams API: `String[] newArray = Arrays.stream(srcArray).toArray(String[]::new)`

> `Arrays.copyOf()` and `Arrays.copyOfRange()` are using `System.arraycopy()` inside.
> 
>`System.arraycopy()` is very fast native implementation  which is used very often in core Java


System.arraycopy() signature:

`arraycopy(Object srcArray, int srcPos, Object destArray, int destPos, int length)`

- **`srcArray`** - source array
- **`srcPos`** - source initial position
- **`destArray`** - destination array 
- **`destPos`** - destination initial position
- **`length`** - how many elements to copy from source to destination

Example for creating twice bigger array and copy data into new array:

```java
int[] srcArr = {1, 2, 3};
int[] destArr = new int[srcArr.length * 2];
System.arraycopy(srcArr, 0, destArr, 0, srcArr.length); // [1, 2, 3, 0, 0, 0]
```

Brute-force approach (same size):
```java
int[] srcArr = {1, 2, 3};
int[] destArr = new int[srcArr.length];

for (int i = 0; i < srcArr.length; i++) {
    destArr[i] = srcArr[i];
}
System.out.println(Arrays.toString(destArr));
```



<div id="go"/>

## Go
copy into new slice which is twice bigger:
```go
srcArr := []int{0, 1, 2}
destArr := make([]int, len(srcArr) * 2)
copy(destArr, srcArr) // pay attention: destination is first argument

fmt.Printf("new array: %v", destArr) // new array: [0 1 2 0 0 0]
```

Brute-force approach (same size):
```go
srcArr := []int{0, 1, 2}
destArr := make([]int, len(srcArr))

for idx, v := range srcArr {
    destArr[idx] = v
}

fmt.Printf("new array: %v", destArr)

```




<div id="python"/>

## Python
There few ways to copy list in Python:

- `dest_list = list(src_list)`
- `dest_list = src_list[:]`
- `dest_list = src_list.copy()` (since Python 3.3 )
- for copy and deep copy:
 
```python
import copy
dest_list = copy.deepcopy(src_list)
```

Brute-force approach (same size):
```python
src_list = [1, 2, 3]
dst_list = []

for n in src_list:
    dst_list.append(n)

print dst_list
```


<div id="javascript"/>

## JavaScript
There few ways to copy list in JavaScript:

using `slice()`:
```javascript
const srcArr = [1, 2, 3]
const destArr = srcArr.slice()
```

using `Array.from()`:
```javascript
const srcArr = [1, 2, 3]
const destArr = Array.from(srcArr)
```

ES6 way `[...]`:
```javascript
const srcArr = [1, 2, 3]
const destArr = [...srcArr]
```

deep copy using JSON:
```javascript
const srcArr = [[1, 2], 3]
const destArr = JSON.parse(JSON.stringify(srcArr))
```

Brute-force approach (same size):
```javascript
srcArr = [1, 2, 3]
destArr = []
srcArr.forEach(n => {destArr.push(n)})

console.log(destArr)
```


<div id="kotlin"/>

## Kotlin
Can use Java's `System.arraycopy()`

Few other simple ways:

- `val destArr = arrayOf(srcArr)`
- `val destArr = srcArr.copyOf()`
- `val destArr = srcArr.clone()` 


Brute-force approach (same size):
```kotlin
val srcArr = arrayOf(1, 2, 3)
val destArr = IntArray(srcArr.size)

for ((idx, value) in srcArr.withIndex()) {
    destArr[idx] = srcArr[idx]
}
println("destArr is: ${Arrays.toString(destArr)}")
```

> You can use `arrayOfNulls<T>(size)` for any custom type of arrays, 
>
>For example: `arrayOfNulls<String>(10)` 



<div id="rust"/>

## Rust

```rust
#[test]
fn copy_copyable_array() {
    // if array contain Copy type, assigning to another variable will copy array.
    // i32 is copyable
    let mut a = [1, 2, 3];
    let b = a;

    assert_eq!(a, b);
    a[0] = 0;

    assert_eq!(a, [0, 2, 3]);
    assert_eq!(b, [1, 2, 3]);
}

#[test]
fn clone_array() {
    #[derive(PartialEq, Debug, Clone)]
    struct S(i32);

    let mut array_a = [S(1), S(2)];
    // let mut array_b = array_a; // -> moves, doesn't copy (to copy struct S should implement Copy)
    let array_b = array_a.clone();

    println!("array_a: {:?}", array_a);
    println!("array_b: {:?}", array_b);
    assert_eq!(array_a, array_b);

    array_a[0] = S(0);

    println!("array_a: {:?}", array_a);
    println!("array_b: {:?}", array_b);

    assert_eq!(array_a, [S(0), S(2)]);
    assert_eq!(array_b, [S(1), S(2)]);
}

#[test]
fn copy_vec() {
    let src = vec![1, 2, 3];

    let dest_a = src.clone();
    assert_eq!(src, dest_a);

    let mut dest_b: Vec<i32> = vec![];
    dest_b.clone_from(&src);
    assert_eq!(src, dest_b);

    let mut dest_c: Vec<i32> = vec![];
    dest_c.append(&mut dest_b);
    assert_eq!(src, dest_c);
}

#[test]
fn iter_collect_slow() {
    let src = vec![1, 2, 3];
    let dest: Vec<_> = src.iter().copied().collect();
    assert_eq!(src, dest);
}

#[test]
fn copy_vec_brute_force_slow() {
    let src = vec![1, 2, 3];
    let mut dest = Vec::with_capacity(src.len());

    for x in src.iter() {
        dest.push(*x);
    }
    assert_eq!(src, dest);
}
```