+++
title = "Merge sort"
template = "level.html"
description = "Merge sort"
[extra]
    keywords = "merge sort, sorting, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Merge sort

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
Stable
more details on [wikipedia on Merge sort](https://en.wikipedia.org/wiki/Merge_sort)


<div id="complexity"/>


## Complexity
`O(N log(N))`
 


<div id="practice"/>


## In practice
`Merge sort` for sorting collection that doesn't fit memory  
 



<div id="java"/>

## Java
```java
package sort;

import java.util.ArrayList;
import java.util.List;


public class MergeSort {

    /**
     * Using Lists for better support of generics. Same algorithms applies to arrays
     **/
    public static <T extends Comparable<? super T>> void mergeSort(List<T> listToSort) {
        if (listToSort == null || listToSort.size() < 2) return;

        List<T> tmpList = new ArrayList<>(listToSort); // temporarily container to keep data for merge sort

        mergeSort(listToSort, tmpList, 0, listToSort.size() - 1); // high index inclusive
    }

    private static <T extends Comparable<? super T>> void mergeSort(List<T> listToSort, List<T> tmpList, int lowIdx, int highIdx) {
        if (lowIdx >= highIdx) return;

        int middleIdx = lowIdx + (highIdx - lowIdx) / 2;
        mergeSort(listToSort, tmpList, lowIdx, middleIdx); //  sort left side
        mergeSort(listToSort, tmpList, middleIdx + 1, highIdx); // sort right side

        merge(listToSort, tmpList, lowIdx, middleIdx, highIdx); // merge sorted from left and right sides
    }

    private static <T extends Comparable<? super T>> void merge(List<T> listToSort, List<T> tmpList, int lowIdx, int middleIdx, int highIdx) {
        for (int i = lowIdx; i <= highIdx; i++) {
            tmpList.set(i, listToSort.get(i));
        }

        int leftIdx = lowIdx;
        int rightIdx = middleIdx + 1;
        int curIdx = leftIdx;

        while (leftIdx <= middleIdx && rightIdx <= highIdx) {
            if (tmpList.get(leftIdx).compareTo(tmpList.get(rightIdx)) < 0) {
                listToSort.set(curIdx, tmpList.get(leftIdx));
                leftIdx++;
            } else {
                listToSort.set(curIdx, tmpList.get(rightIdx));
                rightIdx++;
            }
            curIdx++;
        }

        int remainingCount = middleIdx - leftIdx;
        for (int i = 0; i <= remainingCount; i++) {
            listToSort.set(curIdx + i, tmpList.get(leftIdx + i));
        }
    }
}

```



<div id="go"/>

## Go
```go
package sort

func MergeSort(inputSlice []int) {
	if len(inputSlice) < 2 {
		return
	}
	tmpSlice := make([]int, len(inputSlice), len(inputSlice))
	mergeSortRange(inputSlice, tmpSlice, 0, len(inputSlice)-1)
}

func mergeSortRange(inputSlice, tmpSlice []int, lowIdx, highIdx int) {
	if lowIdx >= highIdx {
		return
	}
	midIdx := lowIdx + (highIdx-lowIdx)/2
	mergeSortRange(inputSlice, tmpSlice, lowIdx, midIdx)
	mergeSortRange(inputSlice, tmpSlice, midIdx+1, highIdx)

	merge(inputSlice, tmpSlice, lowIdx, midIdx, highIdx)
}

func merge(inputSlice, tmpSlice []int, lowIdx, midIdx, highIdx int) {
	for i := lowIdx; i <= highIdx; i++ {
		tmpSlice[i] = inputSlice[i]
	}

	leftIdx := lowIdx
	rightIdx := midIdx + 1
	curIdx := leftIdx

	for leftIdx <= midIdx && rightIdx <= highIdx {
		if tmpSlice[leftIdx] <= tmpSlice[rightIdx] {
			inputSlice[curIdx] = tmpSlice[leftIdx]
			leftIdx++
		} else {
			inputSlice[curIdx] = tmpSlice[rightIdx]
			rightIdx++
		}
		curIdx++
	}

	remaining := midIdx - leftIdx
	for i := 0; i <= remaining; i++ {
		inputSlice[curIdx+i] = tmpSlice[leftIdx+i]
	}
}

```



<div id="python"/>

## Python
```python
def merge(input_list, tmp_list, low_idx, mid_idx, high_idx):
    for i in range(low_idx, high_idx + 1):
        tmp_list[i] = input_list[i]

    left_idx = low_idx
    right_idx = mid_idx + 1
    cur_idx = left_idx

    while left_idx <= mid_idx and right_idx <= high_idx:
        if tmp_list[left_idx] <= tmp_list[right_idx]:
            input_list[cur_idx] = tmp_list[left_idx]
            left_idx = left_idx + 1
        else:
            input_list[cur_idx] = tmp_list[right_idx]
            right_idx = right_idx + 1
        cur_idx = cur_idx + 1

    remaining = mid_idx - left_idx
    for i in range(0, remaining + 1):
        input_list[cur_idx + i] = tmp_list[left_idx + i]


def merge_sort_range(input_list, tmp_list, low_idx, high_idx):
    if low_idx >= high_idx:
        return

    mid_idx = low_idx + (high_idx - low_idx) / 2
    merge_sort_range(input_list, tmp_list, low_idx, mid_idx)
    merge_sort_range(input_list, tmp_list, mid_idx + 1, high_idx)

    merge(input_list, tmp_list, low_idx, mid_idx, high_idx)


def merge_sort(input_list):
    if len(input_list) < 2:
        return

    tmp_list = [None] * len(input_list)
    merge_sort_range(input_list, tmp_list, 0, len(input_list) - 1)
```


<div id="javascript"/>

## JavaScript
```javascript
function mergeSort(inputArr) {
  if (inputArr.length < 2) return inputArr

  const tmpArr = Array(inputArr.length).fill({})

  mergeSortRange(inputArr, tmpArr, 0, inputArr.length - 1)
}

function mergeSortRange(inputArr, tmpArr, lowIdx, highIdx) {
  if (lowIdx >= highIdx) return

  const midIdx = Math.floor(lowIdx + (highIdx - lowIdx)/2)  
  mergeSortRange(inputArr, tmpArr, lowIdx, midIdx)
  mergeSortRange(inputArr, tmpArr, midIdx + 1, highIdx)

  merge(inputArr, tmpArr, lowIdx, midIdx, highIdx)
}

function merge(inputArr, tmpArr, lowIdx, midIdx, highIdx) {
  for (i = lowIdx; i <= highIdx; i++) {
    tmpArr[i] = inputArr[i]
  }

  let leftIdx = lowIdx
  let rightIdx = midIdx + 1
  let curIdx = lowIdx

  while (leftIdx <= midIdx && rightIdx <= highIdx) {
    if (tmpArr[leftIdx] <= tmpArr[rightIdx]) {
      inputArr[curIdx] = tmpArr[leftIdx]
      leftIdx++
    } else {
      inputArr[curIdx] = tmpArr[rightIdx]
      rightIdx++
    }
    curIdx++
  }
  const remaining = midIdx - leftIdx
  for (i = 0; i <= remaining; i++) {
    inputArr[curIdx + i] = tmpArr[leftIdx + i]
  }
}
```


<div id="kotlin"/>

## Kotlin
```kotlin
package sort

import java.util.ArrayList


fun <T : Comparable<in T>> mergeSort(
        listToSort: MutableList<T>,
        tmpList: MutableList<T> = ArrayList(listToSort),
        lowIdx: Int = 0, highIdx: Int = listToSort.size - 1) {

    if (lowIdx >= highIdx || listToSort.size < 2) return

    val middleIdx = lowIdx + (highIdx - lowIdx) / 2
    mergeSort(listToSort, tmpList, lowIdx, middleIdx) //  sort left side
    mergeSort(listToSort, tmpList, middleIdx + 1, highIdx) // sort right side

    merge(listToSort, tmpList, lowIdx, middleIdx, highIdx) // merge sorted from left and right sides
}

private fun <T : Comparable<T>> merge(listToSort: MutableList<T>, tmpList: MutableList<T>, lowIdx: Int, middleIdx: Int, highIdx: Int) {
    for (i in lowIdx..highIdx) {
        tmpList[i] = listToSort[i]
    }
    
    var leftIdx = lowIdx
    var rightIdx = middleIdx + 1
    var curIdx = leftIdx

    while (leftIdx <= middleIdx && rightIdx <= highIdx) {
        if (tmpList[leftIdx] <= (tmpList[rightIdx])) {
            listToSort[curIdx] = tmpList[leftIdx]
            leftIdx++
        } else {
            listToSort[curIdx] = tmpList[rightIdx]
            rightIdx++
        }
        curIdx++
    }

    val remainingCount = middleIdx - leftIdx
    for (i in 0..remainingCount) {
        listToSort[curIdx + i] = tmpList[leftIdx + i]
    }
}

```


<div id="rust"/>

## Rust
```rust
struct Range {
    low_idx: usize,
    high_idx: usize,
}

impl Range {
    pub fn new(low_idx: usize, high_idx: usize) -> Self {
        Range { low_idx, high_idx }
    }
    pub fn len(&self) -> usize {
        self.high_idx - self.low_idx
    }
}

pub fn sort<T: Ord + Clone>(slice: &mut [T]) {
    if slice.len() < 2 {
        return
    }
    let mut helper_vec = Vec::new();
    helper_vec.extend_from_slice(slice);

    sort_range(slice, &Range::new(0, slice.len() - 1), &mut helper_vec);
}

fn sort_range<T: Ord + Clone>(slice: &mut [T], range: &Range, helper_slice: &mut [T]) {
    if range.low_idx >= range.high_idx {
        return;
    }

    let mid_idx = range.low_idx + (range.high_idx - range.low_idx) / 2;

    let left_range = &Range::new(range.low_idx, mid_idx);
    sort_range(slice, left_range, helper_slice);

    let right_range = &Range::new(mid_idx + 1, range.high_idx);
    sort_range(slice, right_range, helper_slice);

    merge(slice, left_range, right_range, helper_slice);
}

fn merge<T: Ord + Clone>(slice: &mut [T], left_range: &Range, right_range: &Range, helper_slice: &mut [T]) {
    helper_slice[left_range.low_idx..=right_range.high_idx]
        .clone_from_slice(&slice[left_range.low_idx..=right_range.high_idx]);

    let mut cur_idx = left_range.low_idx;
    let mut cur_left_idx = left_range.low_idx;
    let mut cur_right_idx = right_range.low_idx;

    while cur_left_idx <= left_range.high_idx && cur_right_idx <= right_range.high_idx {
        if helper_slice[cur_left_idx] < helper_slice[cur_right_idx] {
            slice[cur_idx] = helper_slice[cur_left_idx].clone();
            cur_left_idx += 1;
        } else {
            slice[cur_idx] = helper_slice[cur_right_idx].clone();
            cur_right_idx += 1;
        }
        cur_idx += 1;
    }

    if cur_right_idx <= right_range.high_idx {
        while cur_right_idx <= right_range.high_idx {
            slice[cur_idx] = helper_slice[cur_right_idx].clone();
            cur_right_idx += 1;
            cur_idx += 1;
        }
    } else {
        while cur_left_idx <= left_range.high_idx {
            slice[cur_idx] = helper_slice[cur_left_idx].clone();
            cur_left_idx += 1;
            cur_idx += 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::sort;
    use rand::prelude::*;

    fn sorted<T: Ord>(slice: &[T]) -> bool {
        for i in 0..(slice.len() -1) {
            if slice[i] > slice[i+1] {
                return false;
            }
        }
        true
    }

    #[test]
    fn already_sorted() {
        let mut arr = [1, 2, 3, 4, 5];
        sort(&mut arr);
        assert!(sorted(&arr));
    }

    #[test]
    fn already_sorted_reversed() {
        let mut arr = [5, 4, 3, 2, 1];
        sort(&mut arr);
        assert!(sorted(&arr));
    }

    #[test]
    fn with_dups() {
        let mut arr = [1, 4, 1, 3, 19, 2, 1, 19, 40];
        sort(&mut arr);
        assert!(sorted(&arr));
    }

    #[test]
    fn shuffled_0_100() {
        let mut rng = rand::thread_rng();
        let y: f64 = rng.gen();

        let mut arr: Vec<i32> = (1..100).collect();
        arr.shuffle(&mut rng);

        sort(&mut arr);

        assert!(sorted(&arr))
    }

    #[test]
    fn random_arr() {
        const TIMES: usize = 100;
        for _ in 0..TIMES {
            let mut arr: Vec<i32> = vec![];
            let mut rng = thread_rng();
            let rnd_vec_size: usize = rng.gen_range(1, 1_000);
            for _ in 0..rnd_vec_size {
                let rnd_num: i32 = rng.gen();
                arr.push(rnd_num);
            }
            sort(&mut arr);
            assert!(sorted(&arr));
        }
    }
}
```