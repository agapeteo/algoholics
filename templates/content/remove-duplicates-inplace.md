# Remove duplicates in-place

## Contents

---

- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="java"/>

## Java
```java
package arrays;

import java.util.Arrays;

public class InPlaceDuplicateRemoval {

    public static int remove(Object[] arr) {
        Arrays.sort(arr);
        int i = 0;
        for (int j = 1; j < arr.length; j++) {
            if (!arr[j].equals(arr[i])) {
                arr[i + 1] = arr[j];
                i++;
            }
        }
        return i;
    }
}
```


<div id="go"/>

## Go

```go
package arrays

import "sort"

func RemoveDuplicates(arr []string) int {
	sort.Strings(arr)
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return i
}

```


<div id="python"/>

## Python

```python
def inplace_deduplicate(in_list):
    in_list.sort()
    i = 0
    for j in range(1, len(in_list)):
        if in_list[j] != in_list[i]:
            in_list[i + 1] = in_list[j]
            i = i + 1
    return i
```


<div id="javascript"/>

## JavaScript

```javascript
const inPlaceRemoval = arr => {
  arr.sort()
  let i = 0
  for (j = 1; j < arr.length; j++) {
    if (arr[j] != arr[i]) {
      arr[i + 1] = arr[j]
      i++
    }
  }
  return i
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package arrays

object InPlaceDuplicateRemoval {
    fun <T> remove(arr: Array<T>): Int {
        arr.sort()
        var i = 0
        for (j in 1 until arr.size) {
            if (arr[j] != arr[i]) {
                arr[i + 1] = arr[j]
                i++
            }
        }
        return i
    }
}
```