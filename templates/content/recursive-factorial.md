# Recursive factorial

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