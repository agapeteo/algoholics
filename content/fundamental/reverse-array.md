+++
title = "Reverse Array"
template = "level.html"
description = "Reverse Array"
[extra]
    keywords = "reverse array, array, memory, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Reverse array

## Contents

---

- [Reverse array description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Reverse array

Complexity `O(N)` - requires iterating through all (or half) elements in collection

Space complexity is `O(1)` for in-place approach and `O(N)` for new copy of array



<div id="practice"/>

## In practice
Used for reversing arrays and lists or strings


<div id="java"/>

## Java
Use `Collections.reverse()` to reverse any collection

> To convert array to List use `Arrays.toList()`

Using new array (`O(N)` space complexity):
```java
int[] srcArr = {1, 2, 3};
int[] reversedArr = new int[srcArr.length];

for (int i = 0; i < srcArr.length; i++) {
    reversedArr[srcArr.length - i - 1] = srcArr[i];
}
System.out.println(Arrays.toString(reversedArr));
```


In-place reverse (`O(1)` space complexity):
```java
int[] arr = {1, 2, 3};

for (int i = 0; i <= arr.length/2; i++) {
    int tmp = arr[i];
    arr[i] = arr[arr.length - i - 1];
    arr[arr.length - i - 1] = tmp;
}
System.out.println(Arrays.toString(arr));
```

> Note: to reverse String use: 
>
>`new StringBuilder("someString").reverse().toString()` 
>
> or convert string to char array and than reverse array, for example:
>
> `char[] chars = "someString".toCharArray();` 

<div id="go"/>

## Go
approach #1:
```go
numbers := []int{0, 1, 2, 3, 4}
for i := 0; i < len(numbers)/2; i++ {
    opposite := len(numbers) - 1 - i
    numbers[i], numbers[opposite] = numbers[opposite], numbers[i]
}
fmt.Printf("numbers: %v", numbers)
```

approach #2:
```go
numbers := []int{0, 1, 2}
for leftIdx, rightIdx := 0, len(numbers)-1; leftIdx < rightIdx; leftIdx, rightIdx = leftIdx+1, rightIdx-1 {
    numbers[leftIdx], numbers[rightIdx] = numbers[rightIdx], numbers[leftIdx]
}
fmt.Printf("numbers: %v", numbers)
```


<div id="python"/>

## Python
using `reverse()` (in-place):
```python
numbers = [1, 22, 3]
numbers.reverse()
print numbers  # [3, 22, 1]
```

using slices:
```python
numbers = [1, 22, 3]
numbers_reversed = numbers[::-1]
print numbers_reversed  # [3, 22, 1]
```

straight forward approach (in-place):
```python
def reverse_inplace(any_list):
    for i in range(len(any_list) / 2):
        opposite_idx = len(any_list) - i - 1
        any_list[i], any_list[opposite_idx] = any_list[opposite_idx], any_list[i]


numbers = [1, 22, 3]
reverse_inplace(numbers)
print numbers  # [3, 22, 1]
```

<div id="javascript"/>

## JavaScript
Use `reverse()` method for arrays:
```javascript
console.log([1, 2, 3].reverse()) // [3, 2, 1]
```

implementation example:
```javascript
function reverse(arr) {
  for (i = 0; i <= arr.length/2; i++) {
    tmp = arr[i]
    arr[i] = arr[arr.length - i - 1]
    arr[arr.length - i - 1] = tmp
  }
}
const numbers = [1, 22, 3]
reverse(numbers)

console.log(numbers) // [3, 22, 1]
```



<div id="kotlin"/>

## Kotlin
using `reversed()`:
```kotlin
arrayOf(1, 2, 3).reversed() // [3, 2, 1]
```

using `asReversed()` read-only view:
```kotlin
val reversedView = listOf(0, 1).asReversed()
println("${reversedView.first()}, ${reversedView.last()}") //  1, 0
```

implementing in-place reverse:
```kotlin
val numbers = arrayOf(1, 22, 3)
    
for (idx in 0..numbers.size / 2) {
    val oppositeIdx = numbers.size - idx - 1
    val tmp = numbers[idx]
    numbers[idx] = numbers[oppositeIdx]
    numbers[oppositeIdx] = tmp
}
println("reversed numbers: ${Arrays.toString(numbers)}") // [3, 22, 1]
```



<div id="rust"/>

## Rust

```rust
#[test]
fn reverse_array() {
    let mut arr = [1, 2, 3];
    arr.reverse(); // same for Vec
    assert_eq!([3, 2, 1], arr);
}

#[test]
fn reverse_array_brute_force_slow() {
    let mut arr = [1, 2, 3];
    let mut i = 0;
    let mut j = arr.len() - 1;

    while i < j {
        // fast and efficient swap, just swap internal pointer, no extra variables or copies
        arr.swap(i, j);
        i += 1;
        j -= 1;
    }
    assert_eq!([3, 2, 1], arr);
}
```