# Validate multiple parentheses

## Contents

---

- [Description](#description)
- [Complexity](#complexity)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---

<div id="description"/>

## Description
This algorithms uses `Stack` data structure and works as follows:

- for open parenthesis, like `(` or `{` etc, **put** to `Stack`
- for close parenthesis, like `)` or `}` etc, **pop** from `Stack`, but open parenthesis value should correspond to close parenthesis

To avoid invalid cases like `)(` or `())(` check if stack is empty, if so - return `false`


After checking each character return `true` if `Stack` is empty, otherwise return `false`


<div id="complexity"/>


## Complexity
`O(N)`



<div id="java"/>

## Java

```java
package list;

import java.util.*;

public class JavaMultipleParenthesesValidator {
    private final char[] openArr;
    private final char[] closeArr;

    /**
     * Parentheses for @param openArr and @param closeArr should be in the same order. f.e '{', '[' and '}', ']'
     */
    public JavaMultipleParenthesesValidator(char[] openArr, char[] closeArr) {
        this.openArr = openArr;
        this.closeArr = closeArr;
    }

    public boolean validate(String str) {
        if (str == null || str.isEmpty()) return false;

        Deque<Character> stack = new LinkedList<>();
        Map<Character, Integer> openChars = mapFor(openArr);
        Map<Character, Integer> closeChars = mapFor(closeArr);

        char[] chars = str.toCharArray();

        for (int i = 0; i < chars.length; i++) {
            char curChar = chars[i];
            if (openChars.keySet().contains(curChar)) {
                stack.push(curChar);
            } else if (closeChars.keySet().contains(curChar)) {
                if (stack.isEmpty()) return false;

                Character topChar = stack.pop();
                // index for close char should correspond to open char in array of parentheses
                if (closeChars.get(curChar) != openChars.get(topChar)) {
                    return false;
                }
            }
        }
        return stack.isEmpty();
    }

    private Map<Character, Integer> mapFor(char[] chars) {
        Map<Character, Integer> result = new HashMap<>();

        for (int i = 0; i < chars.length; i++) {
            result.put(chars[i], i);
        }
        return result;
    }

}

```

Test:
```java
package list;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class JavaMultipleParenthesesValidatorTest {
    static char[] openChars = {'[', '{', '('};
    static char[] closeChars = {']', '}', ')'};

    @Test
    public void valid_same_parentheses() {
        // given
        JavaMultipleParenthesesValidator validator = new JavaMultipleParenthesesValidator(openChars, closeChars);

        // when
        boolean actual = validator.validate("to ((be)) or (not) to be");

        // then
        assertTrue(actual);
    }

    @Test
    public void valid_different_parentheses() {
        // given
        JavaMultipleParenthesesValidator validator = new JavaMultipleParenthesesValidator(openChars, closeChars);

        // when
        boolean actual = validator.validate("to ([be]) or {not} to be");

        // then
        assertTrue(actual);
    }

    @Test
    public void not_valid_different_parentheses() {
        // given
        JavaMultipleParenthesesValidator validator = new JavaMultipleParenthesesValidator(openChars, closeChars);

        // when
        boolean actual = validator.validate("to ( [be] {or} not }");

        // then
        assertFalse(actual);
    }

    @Test
    public void startsClosedShouldFail() {
        // given
        JavaMultipleParenthesesValidator validator = new JavaMultipleParenthesesValidator(openChars, closeChars);

        // when
        boolean actual = validator.validate(")(");

        // then
        assertFalse(actual);
    }

    @Test
    public void notValidClosedPositionInTheMiddle() {
        // given
        JavaMultipleParenthesesValidator validator = new JavaMultipleParenthesesValidator(openChars, closeChars);

        // when
        boolean actual = validator.validate("to (be) }{ or not");

        // then
        assertFalse(actual);
    }

}
```




<div id="go"/>

## Go
```go
package list

func validateMultiple(open []rune, close []rune, str string) bool {
	openChars := mapFor(open)
	closeChars := mapFor(close)
	var stack []rune

	runes := []rune(str)
	for _, c := range runes {
		_, isOpen := openChars[c]
		_, isClose := closeChars[c]

		if isOpen {
			stack = append(stack, c) // push
		} else if isClose {
			if len(stack) == 0 {
				return false
			}
			lastChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			if closeChars[c] != openChars[lastChar] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func mapFor(chars []rune) map[rune]int {
	m := make(map[rune]int)
	for i, c := range chars {
		m[c] = i
	}
	return m
}

```

Test:
```go
package list

import "testing"

func Test_validateMultiple(t *testing.T) {
	open := []rune{'(', '{', '['}
	close := []rune{')', '}', ']'}

	type args struct {
		open  []rune
		close []rune
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should be valid 1", args{open, close, "to ((be)) or (not) to be"}, true},
		{"should be valid 2", args{open, close, "to ([be]) or {not} to be"}, true},
		{"should not valid", args{open, close, "to ( [be] {or} not }"}, false},
		{"should not valid, starts with close ", args{open, close, ")("}, false},
		{"should not valid, not closing properly ", args{open, close, "to (be) }{ or not"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateMultiple(tt.args.open, tt.args.close, tt.args.str); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
```


<div id="python"/>

## Python


```python
def map_for(some_list):
    m = {}
    for i in range(len(some_list)):
        m[some_list[i]] = i
    return m


def validate(open_list, close_list, input_str):
    open_map = map_for(open_list)
    close_map = map_for(close_list)

    stack = []
    for each_char in input_str:
        if each_char in open_map:
            stack.append(each_char)
        elif each_char in close_map:
            if len(stack) == 0:
                return False
            top_char = stack.pop()
            if close_map[each_char] != open_map[top_char]:
                return False
    return len(stack) == 0


def test():
    open_list = ["{", "[", "("]
    close_list = ["}", "]", ")"]

    assert (validate(open_list, close_list, "to ((be)) or (not) to be")), "should be valid"
    assert (validate(open_list, close_list, "to ([be]) or {not} to be")), "should be valid"
    assert (not validate(open_list, close_list, "to ( [be] {or} not }")), "should not be valid"
    assert (not validate(open_list, close_list, ")(")), "should not be valid"
    assert (not validate(open_list, close_list, "to (be) }{ or not")), "should not be valid"


if __name__ == "__main__":
    test()

```


<div id="javascript"/>

## JavaScript
```javascript
function mapFor(arr) {
  const map = new Map()
  for (i=0; i < arr.length; i++) {
    map.set(arr[i], i)
  }
  return map
}

function validate(openArr, closedArr, str) {
  const openChars = mapFor(openArr)
  const closeChars = mapFor(closedArr)
  const stack = []

  for (i=0; i < str.length; i++) {
    const curChar = str[i]
    if (openChars.has(curChar)) {
      stack.push(curChar)
    } else if (closeChars.has(curChar)) {
        if (stack.length == 0) return false

        const topChar = stack.pop()
        if (closeChars.get(curChar) != openChars.get(topChar)) {
          return false
        }
    }
  }
  return stack.length == 0
}

const openArr = ["{", "[", "("]
const closeArr = ["}", "]", ")"]
console.log(validate(openArr, closeArr, "to ((be)) or (not) to be")) // true
console.log(validate(openArr, closeArr, "to ([be]) or {not} to be")) // true
console.log(validate(openArr, closeArr, "to ( [be] {or} not }")) // false
console.log(validate(openArr, closeArr, ")(")) // false
console.log(validate(openArr, closeArr, "to (be) }{ or not")) // false
```


<div id="kotlin"/>

## Kotlin
```kotlin
package list

import java.util.*

class MultipleParenthesesValidator(val openArr: Array<Char>, val closeArr: Array<Char>) {

    fun validate(str: String): Boolean {
        if (str.isBlank()) return false;

        val closeChars = mapFor(closeArr)
        val openChars = mapFor(openArr)
        val stack: Deque<Char> = LinkedList()

        for (i in 0 until str.length) {
            val curChar = str[i]
            if (openChars.contains(curChar)) {
                stack.push(curChar)
            } else if (closeChars.contains(curChar)) {
                if (stack.isEmpty()) return false

                val topChar = stack.pop()
                if (closeChars[curChar] != openChars[topChar]) {
                    return false
                }
            }
        }
        return stack.isEmpty()
    }

    private fun mapFor(chars: Array<Char>): Map<Char, Int> {
        val result = hashMapOf<Char, Int>()
        for ((idx, c) in chars.withIndex()) {
            result[c] = idx
        }

        return result
    }
}
```

Test:
```kotlin
package list

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

internal class MultipleParenthesesValidatorTest {
    private val openChars = arrayOf('[', '{', '(')
    private val closeChars = arrayOf(']', '}', ')')

    @Test
    fun `valid with same parentheses`() {
        // given
        val validator = MultipleParenthesesValidator(openChars, closeChars)

        // when
        val actual = validator.validate("to ((be)) or (not) to be")

        // then
        assertTrue(actual)
    }

    @Test
    fun `valid different parentheses`() {
        // given
        val validator = MultipleParenthesesValidator(openChars, closeChars)

        // when
        val actual = validator.validate("to ([be]) or {not} to be")

        // then
        assertTrue(actual)
    }

    @Test
    fun `not valid different parentheses`() {
        // given
        val validator = MultipleParenthesesValidator(openChars, closeChars)

        // when
        val actual = validator.validate("to ( [be] {or} not }")

        // then
        assertFalse(actual)
    }

    @Test
    fun `starts closed should fail`() {
        // given
        val validator = MultipleParenthesesValidator(openChars, closeChars)

        // when
        val actual = validator.validate(")(")

        // then
        assertFalse(actual)
    }

    @Test
    fun `not valid closed position in the middle`() {
        // given
        val validator = MultipleParenthesesValidator(openChars, closeChars)

        // when
        val actual = validator.validate("to (be) }{ or not")

        // then
        assertFalse(actual)
    }
}

```