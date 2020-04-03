+++
title = "Binary Search"
template = "level.html"
description = "Binary Search"
[extra]
    keywords = "Binary Search, search, binary tree, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Binary search

## Contents

---

- [Binary search description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Binary search

Complexity `O log(N)` - Quick search in sorted collection 



<div id="practice"/>

## In practice
This algorithm is used very often to effectively find value in sorted collection.

If target element in sorted collection is not found, algorithm can:

- return `undefined` or `None` value, which is mostly applicable to dynamically typed languages like JS or Python  
- return `-1`
- return index when it stopped searching + 1 with minus sign

Last approach can be useful in some cases when you need to know location of values closest to target value (value you're looking for)
  
Another important question is "_What if there are multiple target values?_"

[Java's](#java) `Collections.binarySearch()` may return any index from found values. Same for `Rust`

[Go's `sort.Search`](#go) will return first index, bur you have to check if value by returned index actually exist.

See [Kotlin example](#kotlin) which shows both approach


There are few ways to find index in the middle:

- `lowIdx + highIdx / 2`
- `lowIdx + (highIdx - lowIdx) / 2`
- `(lowIdx + highIdx) >>> 1` (unsigned right shift) 

>Note: `lowIdx + (highIdx - lowIdx) / 2` 
>
>sometimes helps to avoid memory overflow
>
>`(low + high) >>> 1` prevents integer overflow




<div id="java"/>

## Java
> `java.util` already contains `Collections.binarySearch()` and `Arrays.binarySearch()`  binary search implementations
> this function expect that collection is sorted. If collection contain multiple target values - there is no guarantee which index of those will be returned

```java
public static <T> int binarySearch(List<? extends Comparable<? super T>> list, T value) {
        int lowIdx = 0;
        int highIdx = list.size() - 1;

        while (lowIdx <= highIdx) {
            int midIdx = lowIdx + (highIdx - lowIdx) / 2;

            Comparable<? super T> midValue = list.get(midIdx);
            if (midValue.equals(value)) {
                return midIdx;
            }
            if (midValue.compareTo(value) < 1) {
                lowIdx = midIdx + 1;
            } else {
                highIdx = midIdx - 1;
            }
        }
        return -(lowIdx + 1); // not found, also shows where it stopped with minus sign
    }
}
```


<div id="go"/>

## Go
using Go's `sort.Search()`
```go
func SearchFirstIndex(numbers []int, targetValue int) int {
	idx := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= targetValue
	})
	if idx < len(numbers) && numbers[idx] == targetValue {
		return idx
	}
	return -(idx + 1)
}
```



<div id="python"/>

## Python
regular approach:
```python
def binary_search(sequence, value):
    low_idx, high_idx = 0, len(sequence) - 1

    while low_idx <= high_idx:
        mid_idx = low_idx + (high_idx - low_idx) // 2

        if sequence[mid_idx] == value:
            return mid_idx
        if value > sequence[mid_idx]:
            low_idx = mid_idx + 1
        else:
            high_idx = mid_idx - 1
    return None


numbers = [1, 3, 5, 7, 9]
result = binary_search(numbers, 9)
print result  # 4
```

recursive approach:
```python
def binary_search_recursive(sequence, value, low_idx, high_idx):
    if low_idx > high_idx:
        return None

    mid_idx = low_idx + (high_idx - low_idx) // 2

    if sequence[mid_idx] == value:
        return mid_idx

    if value > sequence[mid_idx]:
        return binary_search_recursive(sequence, value, mid_idx + 1, high_idx)

    else:
        return binary_search_recursive(sequence, value, low_idx, mid_idx - 1)


numbers = [1, 3, 5, 7, 9]
result = binary_search_recursive(numbers, 3, 0, len(numbers) - 1)
print result  # 1
```


<div id="javascript"/>

## JavaScript
```javascript
function binarySearch(arr, value) {
    let lowIdx = 0;
    let highIdx = arr.length - 1;

    while (lowIdx <= highIdx) {
        const midIdx = Math.floor(lowIdx + (highIdx - lowIdx) / 2)

        if (value === arr[midIdx]) {
            return midIdx
        }
        if (value > arr[midIdx]) {
            lowIdx = midIdx + 1
        } else {
            highIdx = midIdx - 1
        }
    }
}

const numbers = [1, 3, 5, 7, 9];
console.log(binarySearch(numbers, 3)); // 1
```
> Note: Pay attention to `Max.floor()`
> 
>
> `console.log(3/2) // 1.5` 
>
> `numbers[3/2]` will return `undefined` without any error 
> 


<div id="kotlin"/>

## Kotlin
using build-in `binarySearch()` method for lists and arrays:
```kotlin
listOf(1, 3, 5, 7, 9).binarySearch(3) // 1 
```

iterative approach:
```kotlin
fun <T> binarySearch(list: List<Comparable<in T>>, value: T): Int {
    var lowIdx = 0
    var highIdx = list.size - 1

    while (lowIdx <= highIdx) {
        val midIdx = lowIdx + (highIdx - lowIdx) / 2
        when {
            list[midIdx] == value -> return midIdx
            list[midIdx] < value -> lowIdx = midIdx + 1
            list[midIdx] > value -> highIdx = midIdx - 1
        }
    }
    return -(lowIdx + 1)
}
```

recursive approach:
```kotlin
fun <T> binarySearchRecursive(list: List<Comparable<in T>>, value: T, lowIdx: Int = 0, highIdx: Int = list.size - 1): Int {
    val notFound = -(lowIdx + 1)
    if (lowIdx > highIdx) {
        return notFound
    }
    val midIdx = lowIdx + (highIdx - lowIdx) / 2
    return when {
        list[midIdx] == value -> midIdx
        list[midIdx] < value -> binarySearchRecursive(list, value, midIdx + 1, highIdx)
        list[midIdx] > value -> binarySearchRecursive(list, value, lowIdx, midIdx - 1)
        else -> notFound
    }
}
```

returning first index:
```kotlin
fun <T> binarySearchLowestIndex(list: List<Comparable<in T>>, value: T): Int {
    var lowIdx = -1
    var highIdx = list.size
    while (lowIdx + 1 < highIdx) {
        val midIdx = (lowIdx + highIdx) ushr 1 // shifting but to right is same as dividing by 2 
        if (list[midIdx] >= value) {
            highIdx = midIdx
        } else {
            lowIdx = midIdx
        }
    }
    return when (value) {
        list[highIdx] -> highIdx
        else -> -(highIdx + 1)
    }
}
```


<div id="rust"/>

## Rust
Rust had built-in binary search method for slices

```rust
#[test]
fn using_built_in() {
    let numbers = [0, 2, 4, 6, 8];

    assert_eq!(Ok(0), numbers.binary_search(&0));
    assert_eq!(Ok(3), numbers.binary_search(&6));

    let expected_missing = Err(2); // Err(2) means that 3 is missing, but would be inserted at idx 3
    assert_eq!(expected_missing, numbers.binary_search(&3));
}
```

See more details and other similar methods like `binary_search_by`, `binary_search_by_key` [here](https://doc.rust-lang.org/std/primitive.slice.html#method.binary_search)

Simple implementation in Rust:

```rust
fn binary_search<T>(slice: &[T], target: &T) -> Result<usize, usize>
    where T: Ord {
    let mut low_idx = 0;
    let mut high_idx = slice.len() - 1;

    while low_idx < high_idx {
        let mid_idx = (low_idx + high_idx) / 2;
        let mid_value = &slice[mid_idx];

        if *mid_value == *target {
            return Ok(mid_idx);
        }

        if *mid_value > *target {
            high_idx = mid_idx - 1;
        } else {
            low_idx = mid_idx + 1;
        }
    }
    Err(low_idx + 1)
}

#[test]
fn check_simple_binary_search() {
    let numbers = [0, 2, 4, 6, 8];

    assert_eq!(Ok(0), binary_search(&numbers, &0));
    assert_eq!(Ok(3), binary_search(&numbers, &6));

    let expected_missing = Err(2); // Err(2) means that 3 is missing, but would be inserted at idx 3
    assert_eq!(binary_search(&numbers, &3), expected_missing);
}
```