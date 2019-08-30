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