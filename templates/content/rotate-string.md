# Rotate string

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
package strings;

public class StringRoller {
    private final char[] chars;

    public enum Direction {Left, Right}

    public StringRoller(String str) {
        this.chars = str.toCharArray();
    }

    public String rotate(Direction direction, int count) {
        switch (direction) {
            case Left:
                for (int k = 0; k < count; k++) {
                    rotateLeft();
                }
                return new String(chars);
            case Right:
                for (int k = 0; k < count; k++) {
                    rotateRight();
                }
                return new String(chars);
            default:
                throw new UnsupportedOperationException();
        }
    }

    private void rotateLeft() {
        char first = chars[0];

        for (int i = 0; i < chars.length - 1; i++) {
            chars[i] = chars[i + 1];
        }
        chars[chars.length - 1] = first;
    }

    private void rotateRight() {
        char last = chars[chars.length - 1];

        for (int i = chars.length - 1; i > 0; i--) {
            chars[i] = chars[i - 1];
        }
        chars[0] = last;
    }
}

```


<div id="go"/>

## Go

```go
package arrays

func RollRight(in string, count int) string {
	runes := []rune(in)
	size := len(runes)

	for c := 0; c < count; c++ {
		var last = runes[size-1]
		for i := size - 1; i > 0; i-- {
			runes[i] = runes[i-1]
		}
		runes[0] = last
	}
	return string(runes)
}

func RollLeft(in string, count int) string {
	runes := []rune(in)
	size := len(runes)

	for c := 0; c < count; c++ {
		var first = runes[0]
		for i := 0; i < size-1; i++ {
			runes[i] = runes[i+1]
		}
		runes[size-1] = first
	}
	return string(runes)
}

```


<div id="python"/>

## Python

```python
def rotate_left(in_str, count):
    chars = list(in_str)

    for k in range(count):
        first = chars[0]
        for i in range(0, len(chars) - 1):
            chars[i] = chars[i + 1]
        chars[len(chars) - 1] = first
    return ''.join(chars)


def rotate_right(in_str, count):
    chars = list(in_str)

    for k in range(count):
        last = chars[len(chars) - 1]
        for i in range(len(chars) - 1, 0, -1):
            chars[i] = chars[i - 1]
        chars[0] = last
    return ''.join(chars)
```


<div id="javascript"/>

## JavaScript

```javascript
const rotateLeft = (str, count) => {
  const chars = Array.from(str)
  for (k = 0; k < count; k++) {
    const first = chars[0]

    for (i = 0; i < chars.length - 1; i++) {
      chars[i] = chars[i+1]
    }
    chars[chars.length - 1] = first
  }
  return chars.join("")
}

const rotateRight = (str, count) => {
  const chars = Array.from(str)
  for (k = 0; k < count; k++) {
    const last = chars[str.length - 1]
 
    for (i = chars.length - 1; i > 0; i--) {
      chars[i] = chars[i-1]
    }
    chars[0] = last
  }
  return chars.join("")
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package arrays

class StringRoller(private val inputString: String) {

    enum class Direction { Left, Right }

    fun rotate(rotationCount: Int = 1, direction: Direction): String {
        return when (direction) {
            Direction.Left -> rotateLeft(rotationCount)
            Direction.Right -> rotateRight(rotationCount)
        }
    }

    private fun rotateRight(rotationCount: Int): String {
        val chars = inputString.toCharArray()
        repeat(rotationCount) {
            val last = chars[chars.size - 1]
            for (i in chars.size - 1 downTo 1) {
                chars[i] = chars[i - 1]
            }
            chars[0] = last
        }
        return String(chars)
    }

    private fun rotateLeft(rotationCount: Int): String {
        val chars = inputString.toCharArray()
        repeat(rotationCount) {
            val first = chars[0]
            for (i in 0 until chars.size - 1) {
                chars[i] = chars[i + 1]
            }
            chars[chars.size - 1] = first
        }
        return String(chars)
    }
}
```