# Validate single parentheses

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
Unlike [multiple parentheses](/validate-multiple-parentheses.html), single parentheses validation doesn't require `Stack` data structure.

Using simple counter is enough: 

- for open parentheses, like `(`, **increment counter**
- for close parentheses, like `)`, **decrement counter**

To avoid invalid cases like `)(` or `())(` check if `counter == 0`, if so - return `false`

After checking each character return `true` if `counter == 0`, otherwise return `false`


<div id="complexity"/>


## Complexity
`O(N)`
 



<div id="java"/>

## Java
```java
package arrays;

public class JavaSingleParenthesesValidator {
    private final char open;
    private final char close;
    private int counter;

    public JavaSingleParenthesesValidator(char open, char close) {
        this.open = open;
        this.close = close;
    }

    public boolean validate(String str) {
        if (str == null || str.isEmpty()) return false;

        char[] chars = str.toCharArray();

        for (int i = 0; i < chars.length; i++) {
            char curChar = chars[i];
            if (curChar == open) {
                counter++;
            } else if (curChar == close) {
                if (counter == 0) return false;

                counter--;
            }
        }
        return counter == 0;
    }
}
```

Test:
```java
package arrays;


import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class JavaSingleParenthesesValidatorTest {

    @Test
    public void validateCorrect() {
        // given
        JavaSingleParenthesesValidator validator = new JavaSingleParenthesesValidator('(', ')');

        // when
        boolean actual = validator.validate("to ((be)) or (not) to be");

        // then
        assertTrue(actual);
    }

    @Test
    public void validateIncorrect() {
        // given
        JavaSingleParenthesesValidator validator = new JavaSingleParenthesesValidator('(', ')');

        // when
        boolean actual = validator.validate("to ((be)(or) not");

        // then
        assertFalse(actual);
    }

    @Test
    public void startsClosedShouldFail() {
        // given
        JavaSingleParenthesesValidator validator = new JavaSingleParenthesesValidator('(', ')');

        // when
        boolean actual = validator.validate(")(");

        // then
        assertFalse(actual);
    }

    @Test
    public void notValidClosedPositionInTheMiddle() {
        // given
        JavaSingleParenthesesValidator validator = new JavaSingleParenthesesValidator('(', ')');

        // when
        boolean actual = validator.validate("to (be)) or (");

        // then
        assertFalse(actual);
    }

}
```



<div id="go"/>

## Go
```go
package arrays

func validate(open rune, close rune, str string) bool {
	counter := 0
	runes := []rune(str)
	for _, rune := range runes {
		if rune == open {
			counter++
		} else if rune == close {
			if counter == 0 {
				return false
			}
			counter--
		}
	}
	return counter == 0
}
```

Test:
```go
package arrays

import "testing"

func Test_validate(t *testing.T) {
	type args struct {
		open  rune
		close rune
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid case", args{rune('('), rune(')'), "to ((be)) or (not) to be"}, true},
		{"not valid case", args{rune('('), rune(')'), "to ((be)(or) not"}, false},
		{"should not start with closed", args{rune('('), rune(')'), ")("}, false},
		{"not valid closed at the end", args{rune('('), rune(')'), "to (be)) or ("}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.open, tt.args.close, tt.args.str); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
```


<div id="python"/>

## Python
```python
def validate(input_str, open_str="(", closed_str=")"):
    counter = 0
    for c in input_str:
        if c == open_str:
            counter = counter + 1
        elif c == closed_str:
            if counter == 0:
                return False
            counter = counter - 1
    return counter == 0


def test_validation():
    assert (validate("to ((be)) or (not) to be")), "should be valid"
    assert (not validate("to ((be)(or) not")), "should be not valid"
    assert (not validate(")(")), "should be not valid"
    assert (not validate("to (be)) or (")), "should be not valid"


if __name__ == "__main__":
    test_validation()

```


<div id="javascript"/>

## JavaScript
```javascript
function validate(open, closed, str) {
  let counter = 0
  for (i=0; i < str.length; i++) {
    let c = str.charAt(i)
    if (c == open) {
      counter++
    } else if (c == closed) {
      if (counter == 0) return false
      counter--
    }
  }
  return counter == 0
}

console.log(validate("(", ")", "to ((be)) or (not) to be")) // true
console.log(validate("(", ")", "to ((be)(or) not")) // false
console.log(validate("(", ")", ")(")) // false
console.log(validate("(", ")", "to (be)) or (")) // false

```


<div id="kotlin"/>

## Kotlin
```kotlin
package arrays

class SingleParenthesesValidator(val open: Char, val closed: Char) {

    fun validate(str: String): Boolean {
        var counter = 0
        if (str.isBlank()) return false;

        for (i in 0 until str.length) {
            val char = str[i]
            if (char == open) {
                counter++
            } else if (char == closed) {
                if (counter == 0) return false
                counter--
            }
        }
        return counter == 0
    }
}
```

Test:
```kotlin
package arrays

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

internal class SingleParenthesesValidatorTest {

    @Test
    fun validateCorrect() {
        // given
        val validator = SingleParenthesesValidator('(', ')')

        // when
        val actual = validator.validate("to ((be)) or (not) to be")

        // then
        assertTrue(actual)
    }

    @Test
    fun validateIncorrect() {
        // given
        val validator = SingleParenthesesValidator('(', ')')

        // when
        val actual = validator.validate("to ((be)(or) not")

        // then
        assertFalse(actual)
    }

    @Test
    fun startsClosedShouldFail() {
        // given
        val validator = SingleParenthesesValidator('(', ')')

        // when
        val actual = validator.validate(")(")

        // then
        assertFalse(actual)
    }

    @Test
    fun notValidClosedPositionInTheMiddle() {
        // given
        val validator = SingleParenthesesValidator('(', ')')

        // when
        val actual = validator.validate("to (be)) or (")

        // then
        assertFalse(actual)
    }
}

```