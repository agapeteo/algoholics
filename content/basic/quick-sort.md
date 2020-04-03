+++
title = "Quick sort"
template = "level.html"
description = "Quick sort"
[extra]
    keywords = "quick sort, sorting, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Quick sort

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
Not stable

Very popular and really fast sorting algorithm

More details on [wikipedia on Quick sort](https://en.wikipedia.org/wiki/Quicksort)

2 approaches:

- [Lomuto](https://upload.wikimedia.org/wikipedia/commons/8/84/Lomuto_animated.gif)
- Hoare

Hoare's approach is used more often in practice than Lomuto's because it causes less swaps.
But Hoare's approach is very very error prone. It's easy to get into infinite loop or incorrectly treat duplicate values, especially last swap

see more details about [Hoare vs Lomuto](https://cs.stackexchange.com/questions/11458/quicksort-partitioning-hoare-vs-lomuto)


Other optimization:

- using insertion sort for small array ranges, f.e. less than 20
- choosing pivot as median in sort range
- 3-way partition - partition array range into 3 parts: less than pivot, equal to pivot and more than pivot




<div id="complexity"/>


## Complexity
`O(N log(N))`

Worst case - `O(N^2)`
 


<div id="practice"/>


## In practice
`Quick sort` is very popular and often used in practice for primitive types  
 



<div id="java"/>

## Java
Lomuto:
```java
package sort;

public class LomutoQuickSort {

    public static void sort(int[] array) {
        if (array == null || array.length < 2) return;

        sort(array, 0, array.length - 1);
    }

    private static void sort(int[] array, int lowIdx, int highIdx) {
        if (lowIdx >= highIdx) return;

        int partitionIdx = partition(array, lowIdx, highIdx);
        sort(array, lowIdx, partitionIdx - 1);
        sort(array, partitionIdx + 1, highIdx);
    }

    private static int partition(int[] array, int lowIdx, int highIdx) {
        int pivotValue = array[highIdx];
        int i = lowIdx;
        int j = i;

        while (j < highIdx) {
            if (array[j] <= pivotValue) {
                swap(array, i, j);
                i++;
            }
            j++;
        }
        swap(array, highIdx, i);
        return i;
    }

    private static void swap(int[] array, int i, int j) {
        int tmp = array[i];
        array[i] = array[j];
        array[j] = tmp;
    }
}
```

Hoare (simplified version from Sedgewick's book):
```java
package sort;

public class HoareQuickSort {

    public static void sort(int[] array) {
        sort(array, 0, array.length - 1);
    }

    private static void sort(int[] array, int lowIdx, int highIdx) {
        if (lowIdx >= highIdx) return;

        int partitionIdx = partition(array, lowIdx, highIdx);
        sort(array, lowIdx, partitionIdx - 1);
        sort(array, partitionIdx + 1, highIdx);
    }

    private static int partition(int[] array, int lowIdx, int highIdx) {
        int pivotValue = array[lowIdx];
        int leftIdx = lowIdx;
        int rightIdx = highIdx + 1;

        while (true) {
            do {
                leftIdx++;
                if (leftIdx == highIdx) break;
            } while (array[leftIdx] < pivotValue);

            do {
                rightIdx--;
                if (rightIdx == lowIdx) break;
            } while (array[rightIdx] > pivotValue);

            if (leftIdx >= rightIdx) break;

            swap(array, leftIdx, rightIdx);
        }
        swap(array, lowIdx, rightIdx); // put pivot value to final pivot index
        return rightIdx;
    }

    private static void swap(int[] array, int leftIdx, int rightIdx) {
        int tmpValue = array[leftIdx];
        array[leftIdx] = array[rightIdx];
        array[rightIdx] = tmpValue;
    }
}

```



<div id="go"/>

## Go
```go
package sort

type QuickSortingStrategy int

const (
	Lomuto QuickSortingStrategy = iota
	Hoare
)

var PickedQuickSortStrategy = Lomuto

func QuickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	quickSortRange(arr, 0, len(arr)-1)
}

func quickSortRange(arr []int, lowIdx, highIdx int) {
	if lowIdx >= highIdx {
		return
	}

	var partitionIdx = -1

	if PickedQuickSortStrategy == Lomuto {
		partitionIdx = partitionLomuto(arr, lowIdx, highIdx)
	}
	if PickedQuickSortStrategy == Hoare {
		partitionIdx = partitionHoare(arr, lowIdx, highIdx)
	}
	if partitionIdx == -1 {
		panic("PickedQuickSortStrategy not defined")
	}

	quickSortRange(arr, lowIdx, partitionIdx-1)
	quickSortRange(arr, partitionIdx+1, highIdx)
}

func partitionLomuto(arr []int, lowIdx, highIdx int) int {
	pivot := arr[highIdx]
	i := lowIdx

	for j := lowIdx; j < highIdx; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[highIdx], arr[i] = arr[i], arr[highIdx]
	return i
}

func partitionHoare(arr []int, lowIdx, highIdx int) int {
	pivot := arr[lowIdx]
	leftIdx := lowIdx
	rightIdx := highIdx + 1

	for {
		for {
			leftIdx++
			if leftIdx == highIdx || arr[leftIdx] >= pivot {
				break
			}
		}
		for {
			rightIdx--
			if  rightIdx == lowIdx || arr[rightIdx] <= pivot{
				break
			}
		}
		if leftIdx >= rightIdx {
			break
		}
		arr[leftIdx], arr[rightIdx] = arr[rightIdx], arr[leftIdx]
	}
	arr[rightIdx], arr[lowIdx] = arr[lowIdx], arr[rightIdx]
	return rightIdx
}

```

see also Go's implementation of quick sort in [sort package](https://golang.org/src/sort/sort.go)


<div id="python"/>

## Python
Lomuto:
```python
def partition(input_list, low_idx, high_idx):
    pivot_value = input_list[high_idx]
    i = low_idx

    for j in range(i, high_idx):
        if input_list[j] <= pivot_value:
            input_list[i], input_list[j] = input_list[j], input_list[i]
            i = i + 1
    input_list[high_idx], input_list[i] = input_list[i], input_list[high_idx]
    return i


def _sort_range(input_list, low_idx, high_idx):
    if low_idx >= high_idx:
        return

    partition_idx = partition(input_list, low_idx, high_idx)
    _sort_range(input_list, low_idx, partition_idx - 1)
    _sort_range(input_list, partition_idx + 1, high_idx)


def sort(input_list):
    if not input_list or len(input_list) < 2:
        return
    _sort_range(input_list, 0, len(input_list) - 1)
```

Hoare:
```python
def partition(input_list, low_idx, high_idx):
    pivot_value = input_list[low_idx]
    left_idx = low_idx
    right_idx = high_idx + 1

    while True:

        while True:
            left_idx = left_idx + 1
            if left_idx == high_idx or input_list[left_idx] >= pivot_value:
                break

        while True:
            right_idx = right_idx - 1
            if right_idx == low_idx or input_list[right_idx] <= pivot_value:
                break

        if left_idx >= right_idx:
            break

        input_list[left_idx], input_list[right_idx] = input_list[right_idx], input_list[left_idx]

    input_list[low_idx], input_list[right_idx] = input_list[right_idx], input_list[low_idx]
    return right_idx


def _sort_range(input_list, low_idx, high_idx):
    if low_idx >= high_idx:
        return

    partition_idx = partition(input_list, low_idx, high_idx)
    _sort_range(input_list, low_idx, partition_idx - 1)
    _sort_range(input_list, partition_idx + 1, high_idx)


def sort(input_list):
    if not input_list or len(input_list) < 2:
        return
    _sort_range(input_list, 0, len(input_list) - 1)
```

<div id="javascript"/>

## JavaScript
Lomuto:
```javascript
function sort(arr, lowIdx = 0, highIdx = arr.length - 1) {
  if (!arr || lowIdx >= highIdx || arr.length < 2) return

  const partitionIdx = partition(arr, lowIdx, highIdx)
  sort(arr, lowIdx, partitionIdx - 1)
  sort(arr, partitionIdx + 1, highIdx)
}

function partition(arr, lowIdx, highIdx) {
  const pivotValue = arr[highIdx]
  let i = lowIdx

  for (j = i; j < highIdx; j++) {
    if (arr[j] <= pivotValue) {
      swap(arr, i, j)
      i++
    }
  }
  swap(arr, highIdx, i)
  return i
}

function swap(arr, i, j) {
  const tmp = arr[i]
  arr[i] = arr[j]
  arr[j] = tmp
}
```

Hoare:
```javascript
function sort(arr, lowIdx = 0, highIdx = arr.length - 1) {
  if (!arr || lowIdx >= highIdx || arr.length < 2) return

  const partitionIdx = partition(arr, lowIdx, highIdx)
  sort(arr, lowIdx, partitionIdx - 1)
  sort(arr, partitionIdx + 1, highIdx)
}

function partition(arr, lowIdx, highIdx) {
  const pivotValue = arr[lowIdx]
  let leftIdx = lowIdx
  let rightIdx = highIdx + 1

  while(true) {
    do {
       leftIdx++;
       if (leftIdx == highIdx) {
         break;
       }
    } 
    while(arr[leftIdx] < pivotValue);

    do {
      rightIdx--;
      if (rightIdx == lowIdx) {
        break;
      }
    }
    while (arr[rightIdx] > pivotValue);

    if (leftIdx >= rightIdx) {
      break;
    }
    swap(arr, leftIdx, rightIdx)
  }

  swap(arr, lowIdx, rightIdx)
  return rightIdx
}

function swap(arr, i, j) {
  const tmp = arr[i]
  arr[i] = arr[j]
  arr[j] = tmp
}
```


<div id="kotlin"/>

## Kotlin
Lomuto:

```kotlin
package sort

class LomutoQuickSort {
    companion object {
        fun <T : Comparable<in T>> sort(array: Array<T>, lowIdx: Int = 0, highIdx: Int = array.size - 1) {
            if (array.size < 2 || lowIdx >= highIdx) return

            val partitionIdx = partition(array, lowIdx, highIdx)
            sort(array, lowIdx, partitionIdx - 1)
            sort(array, partitionIdx + 1, highIdx)
        }

        private fun <T : Comparable<in T>> partition(array: Array<T>, lowIdx: Int, highIdx: Int): Int {
            val pivot = array[highIdx]
            var i = lowIdx

            for (j in i until highIdx) {
                if (array[j] <= pivot) {
                    swap(array, i, j)
                    i++
                }
            }
            swap(array, highIdx, i)
            return i
        }

        private fun <T> swap(array: Array<T>, i: Int, j: Int) {
            val tmp = array[i]
            array[i] = array[j]
            array[j] = tmp
        }
    }
}

```

Hoare:
```kotlin
package sort

class HoareQuickSort {
    companion object {
        fun <T : Comparable<in T>> sort(array: Array<T>, lowIdx: Int = 0, highIdx: Int = array.size - 1) {
            if (array.size < 2 || lowIdx >= highIdx) return

            val partitionIdx = partition(array, lowIdx, highIdx)
            sort(array, lowIdx, partitionIdx - 1)
            sort(array, partitionIdx + 1, highIdx)
        }

        private fun <T : Comparable<in T>> partition(array: Array<T>, lowIdx: Int, highIdx: Int): Int {
            val pivotValue = array[lowIdx]
            var leftIdx = lowIdx
            var rightIdx = highIdx + 1

            while (true) {
                do {
                    leftIdx++
                    if (leftIdx == highIdx) break
                } while (array[leftIdx] < pivotValue)

                do {
                    rightIdx--
                    if (rightIdx == lowIdx) break
                } while (array[rightIdx] > pivotValue)

                if (leftIdx >= rightIdx) break

                swap(array, leftIdx, rightIdx)
            }

            swap(array, lowIdx, rightIdx)
            return rightIdx
        }

        private fun <T> swap(array: Array<T>, i: Int, j: Int) {
            val tmp = array[i]
            array[i] = array[j]
            array[j] = tmp
        }
    }
}
```



<div id="rust"/>

## Rust

Lomuto:

```rust
use rand::prelude::*;

pub fn sort<T: Ord>(slice: &mut [T]) {
    if slice.len() < 2 {
        return;
    }
    slice.shuffle(&mut rand::thread_rng());

    sort_range(slice, 0, slice.len() - 1);
}

fn sort_range<T: Ord>(slice: &mut [T], lo_idx: usize, hi_idx: usize) {
    if lo_idx >= hi_idx {
        return;
    }
    let partition_idx = partition(slice, lo_idx, hi_idx);
    if partition_idx > 0 {
        sort_range(slice, lo_idx, partition_idx - 1);
    }
    sort_range(slice, partition_idx + 1, hi_idx);
}

pub fn partition<T: Ord>(slice: &mut [T], lo_idx: usize, hi_idx: usize) -> usize {
    let mut i = lo_idx;
    for j in i..hi_idx {
        if slice[j] <= slice[hi_idx] {
            slice.swap(i, j);
            i += 1;
        }
    }
    slice.swap(hi_idx, i);
    i
}
```

Hoare: 
```rust
use rand::prelude::*;

pub fn sort<T: Ord>(slice: &mut [T]) {
    if slice.len() < 2 {
        return;
    }
    slice.shuffle(&mut rand::thread_rng());

    sort_range(slice, 0, slice.len() - 1);
}

fn sort_range<T: Ord>(slice: &mut [T], lo_idx: usize, hi_idx: usize) {
    if lo_idx >= hi_idx {
        return;
    }
    let partition_idx = partition(slice, lo_idx, hi_idx);
    if partition_idx > 0 {
        sort_range(slice, lo_idx, partition_idx - 1);
    }
    sort_range(slice, partition_idx + 1, hi_idx);
}

pub fn partition<T: Ord>(slice: &mut [T], lo_idx: usize, hi_idx: usize) -> usize {
    let mut left_idx = lo_idx;
    let mut right_idx = hi_idx + 1;

    loop {
        loop {
            left_idx += 1;
            if left_idx == hi_idx || slice[left_idx] >= slice[lo_idx] {
                break;
            }
        }
        loop {
            right_idx -= 1;
            if right_idx == lo_idx || slice[right_idx] <= slice[lo_idx] {
                break;
            }
        }
        if left_idx >= right_idx {
            break;
        }
        slice.swap(left_idx, right_idx);
    }
    slice.swap(lo_idx, right_idx);

    right_idx
}
```