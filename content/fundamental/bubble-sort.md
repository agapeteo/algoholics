+++
title = "Bubble sort"
template = "level.html"
description = "Bubble sort"
[extra]
    keywords = "bubble sort, sorting, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Bubble sort

## Contents

---

- [Description](#description)
- [Complexity](#complexity)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---

<div id="description"/>

## Description
more details on [wikipedia on Bubble sort](https://en.wikipedia.org/wiki/Bubble_sort)


<div id="complexity"/>


## Complexity
`O(n^2)`
 


<div id="practice"/>


## In practice
`Bubble sort` used only for learning, it's not used in practice. 

For small arrays [insert sort](/insert-sort.html) preferred 



<div id="java"/>

## Java
```java
static void bubbleSort(int[] array) {
    boolean sorted = false;
    int lastUnsorted = array.length - 1;
    while (!sorted) {
        sorted = true;
        for (int i = 0; i < lastUnsorted; i++) {
            if (array[i + 1] < array[i]) {
                int tmp = array[i];
                array[i] = array[i + 1];
                array[i + 1] = tmp;

                sorted = false; // reset sorted
            }
        }
        lastUnsorted--;
    }
}
```



<div id="go"/>

## Go
```go
func bubbleSort(arr []int) {
	sorted := false
	lastUnsorted := len(arr) - 1
	for !sorted {
		sorted = true
		for i := 0; i < lastUnsorted; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		lastUnsorted--
	}
}
```



<div id="python"/>

## Python
```python
def bubble_sort(arr):
    is_sorted = False
    last_unsorted = len(arr) - 1
    while not is_sorted:
        is_sorted = True
        for i in range(last_unsorted):
            if arr[i + 1] < arr[i]:
                arr[i], arr[i + 1] = arr[i + 1], arr[i]
                is_sorted = False
        last_unsorted -= 1
```


<div id="javascript"/>

## JavaScript
```javascript
function bubbleSort(arr) {
    let sorted = false;
    let lastUnsorted = arr.length - 1;
    while(!sorted) {
        sorted = true;
        for (let i = 0; i < lastUnsorted; i++) {
            if (arr[i + 1] < arr[i]) {
                let tmp = arr[i];
                arr[i] = arr[i + 1];
                arr[i + 1] = tmp;
                sorted = false
            }
        }
        lastUnsorted--
    }
}
```


<div id="kotlin"/>

## Kotlin
```kotlin
fun bubbleSort(array: IntArray) {
    fun swap(i: Int, j: Int) {
        val tmp = array[i]
        array[i] = array[j]
        array[j] = tmp
    }
    var sorted = false
    var lastUnsorted = array.size - 1
    while (!sorted) {
        sorted = true
        for (i in 0 until lastUnsorted - 1) {
            if (array[i + 1] < array[i]) {
                swap(i, i + 1)
                sorted = false
            }
        }
        lastUnsorted--
    }
}
```

<div id="rust"/>

## Rust
```rust
fn sort<T: PartialOrd>(slice: &mut [T]) {
    if slice.len() < 2 {
        return;
    }
    let mut last_idx = &slice.len() - 1;

    while last_idx > 0 {
        let mut max_idx = 0;

        for i in 1..=last_idx {
            if slice[i] > slice[max_idx] {
                max_idx = i;
            }
        }
        slice.swap(last_idx, max_idx);
        last_idx -= 1;
    }
}
```
