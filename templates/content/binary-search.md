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

---


<div id="description" />

## Binary search

Complexity `O log(N)` - Quick search in sorted collection of data 



<div id="practice"/>

## In practice
Used very often to effectively find value in sorted collection 

If target element in sorted collection is not found, algorithm can:

- return `undefined` or `None` value, which is mostly applicable to dynamically typed languages like JS or Python  
- return `-1`
- return index when it stopped searching + 1 with minus sign

Last approach can be useful in some cases to understand location of values closest to target value 
  
Another important question is "_What if there multiple target values?_"

[Java's](#java) `Collections.binarySearch()`, for example, will not have guarantees on which index will be returned

[Go's `sort.Search`](#go) will return first index, bur you have to check if value by returned index actually exist.

See [Kotlin example](#kotlin) which shows both approach


There are two way to find index in the middle:

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
  var lowIdx = 0
  var highIdx = arr.length - 1

  while (lowIdx <= highIdx) {    
    var midIdx = Math.floor(lowIdx + (highIdx - lowIdx) / 2)
    
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

const numbers = [1, 3, 5, 7, 9]
console.log(binarySearch(numbers, 3)) // 1 
console.log(binarySearch(numbers, 4)) // undefined
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
        var midIdx = lowIdx + (highIdx - lowIdx) / 2
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
