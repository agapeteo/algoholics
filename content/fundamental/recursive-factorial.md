+++
title = "Recursive factorial"
template = "level.html"
description = "Recursive factorial"
[extra]
    keywords = "Recursive factorial, recursion, factorial, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Recursive factorial

## Contents

---

- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---



<div id="java"/>

## Java
```java
static int recursiveFactorial(int n) {
    if (n == 1) return 1;
    return n * recursiveFactorial(n - 1);
}
```



<div id="go"/>

## Go
```go
func recursiveFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}
```



<div id="python"/>

## Python
```python
def recursive_factorial(n):
    if n == 1:
        return 1
    return n * recursive_factorial(n - 1)
```


<div id="javascript"/>

## JavaScript
```javascript
function recursiveFactorial(n) {
  if (n == 1) return 1
  return n * recursiveFactorial(n - 1)
}
```


<div id="kotlin"/>

## Kotlin
```kotlin
fun recursiveFactorial(n: Int): Int {
    if (n == 1) return 1;
    return n * recursiveFactorial(n - 1);
}
```

<div id="rust"/>

## Rust
```rust
fn factorial(n: usize) -> usize {
    if n <= 1 {
        return 1;
    }
    n * factorial(n - 1)
}

#[test]
fn test() {
    assert_eq!(factorial(1), 1);
    assert_eq!(factorial(2), 2);
    assert_eq!(factorial(3), 6);
    assert_eq!(factorial(5), 120);
}
```