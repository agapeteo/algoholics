# Bubble sort

## Contents

---

- [Complexity](#complexity)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---

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
        for i in range(last_unsorted - 1):
            if arr[i + 1] < arr[i]:
                arr[i], arr[i + 1] = arr[i + 1], arr[i]
                is_sorted = False
        last_unsorted = last_unsorted - 1


arr = [9, 1, 4, 7]
bubble_sort(arr)

print arr
```


<div id="javascript"/>

## JavaScript
```javascript
function bubbleSort(arr) {
  let sorted = false
  let lastUnsorted = arr.length - 1
  while(!sorted) {
    sorted = true
    for (i = 0; i < lastUnsorted; i++) {
      if (arr[i + 1] < arr[i]) {
        tmp = arr[i]
        arr[i] = arr[i + 1]
        arr[i + 1] = tmp
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