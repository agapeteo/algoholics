# Find maximum

## Contents

---

- [Find maximum description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="description" />

## Find maximum

Complexity `O(N)` - requires iterating through all elements in collection 



<div id="practice"/>

## In practice
Usually most of languages have build-in function which implements finding maximum


<div id="java"/>

## Java
`java.util.Collections.max()` returns max value in collection

simple old school implementation:
```java
static int max(List<Integer> numbers) {
        int max = Integer.MIN_VALUE;
        for (int i = 0; i < numbers.size(); i++) {
            int curNumber = numbers.get(i);
            if (curNumber > max) {
                max = curNumber;
            }
        }
        return max;
    }
```

using streams:
```java
int max = numbers.stream().max(Integer::compareTo).get();
```

using Collections.max:
```java
Collections.max(numbers);
```

<div id="go"/>

## Go

#### find index for element:
```go
func Max(numbers []int) int {
	var max int
	for idx, n := range numbers {
		if idx == 0 || n < max {
			max = n
		}
	}
	return max
}
```


<div id="python"/>

## Python

using build-in function `max()`:
```python
print max([1, 5, -3, 11]) 
```

implementing own max_in_list:
```python
def max_in_list(numbers):
    maximum = numbers[0]
    for num in numbers:
        if num > maximum:
            maximum = num
    return maximum
```



<div id="javascript"/>

## JavaScript

using JavaScript's `Math.max()`:
```javascript
const numbers = [1, 22, 33, 4]
console.log(Math.max(...numbers))
```

custom max:
```javascript
function max(arr) {
	var max = arr[0]
    max.forEach(n => {
      if (n > max) {
      	max = n
      }
    })
  	return max
}
```

<div id="kotlin"/>

## Kotlin

Kotlin can use Java's `Collections.max()`

Custom implementation:
```kotlin
fun max(numbers: List<Int>): Int {
    var max = Int.MIN_VALUE
    for (n in numbers) {
        if (n > max) {
            max = n
        }
    }
    return max
}
```





