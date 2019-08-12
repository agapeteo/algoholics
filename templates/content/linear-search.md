# Linear search

## Contents

---

- [Linear search description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="description" />

## Linear search

Complexity O(N) - requires iterating through all elements in collection, usually unsorted array or list 



<div id="practice"/>

## In practice
Examples include: 

- check if elements exist in unsorted array
- print all odd values 
- finding index of element


<div id="java"/>

## Java
#### finding index of element
```java
int indexOf(int value, List<Integer> list) {
    for (int idx = 0; idx < list.size(); idx++) {
        if (list.get(idx) == value) return idx;
    }
    return -1; // not found
}
```
> Interface `List` already has method `int indexOf(Object o);`
>

#### print elements using forEach:
```java
list.forEach(System.out::println);
```

#### filter elements using streams:
```java
list.stream().filter(i -> i > 0).collect(Collectors.toList());
```


<div id="go"/>

## Go

#### find index for element:
```go
func idxFor(value int, numbers []int) int {
	for idx, element := range numbers {
		if element == value {
			return idx
		}
	}
	return -1
}
```


<div id="python"/>

## Python

#### find index for element
```python
def idx_for(value, numbers):
    for idx in range(len(numbers)):
        if numbers[idx] == value:
            return idx
    return -1
```



<div id="javascript"/>

## JavaScript

```javascript
function idxFor(value, arr) {
	for (i = 0; i < arr.length; i++) {
		if (arr[i] == value) {
			return i;
		}
	}
	return -1;
}

```

#### using `Array.findIndex`:
```javascript
[1, 5, -3, 11].findIndex(element => element === -3) // 2
```

#### print elements greater than 0 using `forEach`
```javascript
[-2, 5, 0, 11, -9, 2].forEach(element => {
	if (element > 0) {
		console.log(element);
	}
});
```
> Note: there is no way to `break` or `return` from `forEach` cycle (except throwing exception)! 
> 

> The difference between `forEach` and `map`:
>
> - `forEach` applies provided function on each element (f.e print) 
> - `map` returns new copy of array with applied function on each element
> 
> for example `[1, 2, 3].map(i => i * 2)` will return array with values `[2, 4, 6]`
> 
>See more details [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/forEach) and [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/map) 
>



<div id="kotlin"/>

## Kotlin

#### print all number greater than 0
```kotlin
listOf(0, -2, 4, 12).filter { it > 0 }.forEach { println(it) }
```

### index of element:
```kotlin
fun idxOf(value: Int, list: List<Int>): Int {
    for ((idx, element) in list.withIndex()) {
        if (element == value) return idx
    }
    return -1
}
```

> Note: as in `Java` there is already available `indexOf()` method in `List`
> 





