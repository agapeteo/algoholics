# Insertion sort

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
more details on [wikipedia on Insertion sort](https://en.wikipedia.org/wiki/Insertion_sort)


<div id="complexity"/>


## Complexity
`O(n^2)`
 


<div id="practice"/>


## In practice
Most popular among quadratic sort logarithms in practice



<div id="java"/>

## Java
```java
static void insertionSort(int[] array) {
    for (int i = 1; i < array.length; i++) {
        int j = i;
        while (j > 0 && array[j] < array[j - 1]) {
            int tmp = array[j];
            array[j] = array[j - 1];
            array[j - 1] = tmp;
            j--;
        }
    }
}
```



<div id="go"/>

## Go
```go
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
```



<div id="python"/>

## Python
```python
def insertion_sort(arr):
    for i in range(len(arr)):
        j = i
        while j > 0 and arr[j] < arr[j - 1]:
            arr[j], arr[j - 1] = arr[j - 1], arr[j]
            j = j - 1


arr = [9, 1, 4, 7]
insertion_sort(arr)

print arr
```


<div id="javascript"/>

## JavaScript
```javascript
function insertionSort(arr) {
  for (i = 1; i < arr.length; i++) {
    let j = i
    while (j > 0 && arr[j] < arr[j - 1]) {
      let tmp = arr[j]
      arr[j] = arr[j - 1]
      arr[j - 1] = tmp
      j--
    }
  }
}
```


<div id="kotlin"/>

## Kotlin
```kotlin
fun insertionSort(array: IntArray) {
    fun swap(i: Int, j: Int) {
        val tmp = array[i]
        array[i] = array[j]
        array[j] = tmp
    }
    for (i in 1 until array.size) {
        var j = i
        while (j > 0 && array[j] < array[j - 1]) {
            swap(j, j - 1)
            j--
        }
    }
}
```