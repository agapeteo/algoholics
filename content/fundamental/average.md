+++
title = "Average"
template = "level.html"
description = "Calculate average"
[extra]
    keywords = "Calculate average, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Calculate average

## Contents

---

- [Calculate average description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Calculate average

Complexity `O(N)` - requires iterating through all elements in collection 



<div id="practice"/>

## In practice
Used often when average number required


<div id="java"/>

## Java
```java
public static int average(List<Integer> numbers) {
        Integer sum = 0;
        for (Integer eachNumber : numbers){
            sum += eachNumber;
        }
        return sum/numbers.size();
    }
}
```


<div id="go"/>

## Go

```go
func Average(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum/len(numbers)
}
```


<div id="python"/>

## Python

```python
def average(numbers):
    sum_ = 0
    for num in numbers:
        sum_ += num
    return sum_/len(numbers)
```



<div id="javascript"/>

## JavaScript

```javascript
function average(arr) {
  var sum = 0;
  arr.forEach(n => {sum += n})
  return sum/arr.length
}
```

using `reduce()`:
```javascript
function average(arr) {
  const sum = arr.reduce((sum, n) => sum += n)
  return sum/arr.length
}
```


<div id="kotlin"/>

## Kotlin
Kotlin has `average()` method for arrays and lists for number types:
```kotlin
listOf(0, 200).average() // 100
```

custom way:
```kotlin
fun average(numbers: List<Int>): Int {
    var sum = 0
    for (n in numbers) {
        sum += n
    }
    return sum/numbers.size
}
```


<div id="rust"/>

## Rust

```rust
#[test]
fn average() {
    let arr = [1, 2];
    let average = arr.iter().sum::<i32>() as f32 / arr.len() as f32;

    assert_eq!(true, average.eq(&1.5));
}
```