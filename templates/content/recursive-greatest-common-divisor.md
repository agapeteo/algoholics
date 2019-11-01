# Greatest common divisor

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