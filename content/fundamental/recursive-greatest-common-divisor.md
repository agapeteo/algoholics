+++
title = "Greatest common divisor recursive"
template = "level.html"
description = "Greatest common divisor, GCD, recursive"
[extra]
    keywords = "Greatest common divisor, GCD, recursion, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Greatest common divisor

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
package recursion;

public class GreatestCommonDivisor {
    public static int gcd(int a, int b) {
        if (b == 0) return a;
        return gcd(b, a % b);
    }
}
```



<div id="go"/>

## Go
```go
package recursion

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
```



<div id="python"/>

## Python
```python
def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)

```


<div id="javascript"/>

## JavaScript
```javascript
const gcd = (a, b) => {
    if (b == 0) {
        return a
    }
    return gcd(b, a % b)
}
```


<div id="kotlin"/>

## Kotlin
```kotlin
package recursion

fun gcd(a: Int, b: Int): Int {
    if (b == 0) return a
    return gcd(b, a % b)
}
```

<div id="rust"/>

## Rust
```rust
fn gcd(a: usize, b: usize) -> usize {
    if b == 0 {
        return a;
    }
    gcd(b, a % b)
}

#[test]
fn test() {
    assert_eq!(gcd(12, 6), 6);
    assert_eq!(gcd(12, 8), 4);
    assert_eq!(gcd(9, 11), 1);
}
```