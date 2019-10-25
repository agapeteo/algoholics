# Recursive string permutations

## Contents

---

- [Complexity](#complexity)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="complexity"/>


## Complexity
`O(N!)`

 

<div id="java"/>

## Java
```java
package recursion;

import java.util.ArrayList;
import java.util.List;

public class StringPermutation {

    public static List<String> permute(String str) {
        return permute("", str, new ArrayList<>());
    }

    private static List<String> permute(String prefix, String str, List<String> permutations) {
        if (!str.isEmpty()) {
            for (int i = 0; i < str.length(); i++) {
                String newPrefix = String.format("%s%s", prefix, str.charAt(i));
                String strBeforeChar = str.substring(0, i);

                String strAfterChar = "";
                boolean isLastChar = i == str.length() - 1;
                if (!isLastChar) {
                    strAfterChar = str.substring(i + 1);
                }
                String newString = String.format("%s%s", strBeforeChar, strAfterChar);
                permute(newPrefix, newString, permutations);
            }
        } else {
            permutations.add(prefix);
        }
        return permutations;
    }
}

```


<div id="go"/>

## Go
```go
package recursion

func PermuteStr(str string) []string {
	perms := make([]string, 0)

	return permuteStr([]rune(""), []rune(str), &perms)
}

func permuteStr(prefix, str []rune, perms *[]string) []string {
	if len(str) > 0 {
		for idx, char := range str {
			newPrefix := append(prefix, char)
			strBeforeRune := str[:idx]
			strAfterRune := str[idx+1:]
			newStr := append(strAfterRune, strBeforeRune...)

			permuteStr(newPrefix, newStr, perms)
		}
	} else {
		*perms = append(*perms, string(prefix))
	}
	return *perms
}
```



<div id="python"/>

## Python
```python
def _string_permutation(prefix, in_str, permutations):
    if in_str:
        for char_idx in range(len(in_str)):
            new_prefix = prefix + in_str[char_idx]
            str_before = in_str[:char_idx]
            str_after = ""
            if char_idx != len(in_str) - 1:
                str_after = in_str[char_idx + 1:]
            new_str = str_before + str_after

            _string_permutation(new_prefix, new_str, permutations)
    else:
        permutations.append(prefix)
    return permutations


def string_permutation(in_str):
    return _string_permutation("", in_str, [])
```

<div id="javascript"/>

## JavaScript
```javascript
function permute(str, prefix, permutations) {
    if (!prefix) prefix = "";
    if (!permutations) permutations = [];

    if (str) {
        for (let i = 0; i < str.length; i++) {
            const newPrefix = `${prefix}${str.charAt(i)}`;
            const strBeforeChar = str.substring(0, i);
            let strAfterChar = "";
            if (i != str.length - 1) {
                strAfterChar = str.substring(i + 1);
            }
            const newStr = `${strBeforeChar}${strAfterChar}`;

            permute(newStr, newPrefix, permutations);
        }
    } else {
        permutations.push(prefix)
    }
    return permutations;
}
```


<div id="kotlin"/>

## Kotlin
```kotlin
package recursion

object StringPermutation {

    fun permute(str: String, prefix: String = "", curPermutations: MutableList<String> = mutableListOf()): List<String> {
        if (str.isNotEmpty()) {
            for (charIdx in str.indices) {
                val newPrefix = "${prefix}${str[charIdx]}"
                val strBeforeCurChar = str.substring(0, charIdx)
                val isLastIdx = charIdx == str.length - 1
                val strAfterCurChar = if (isLastIdx) "" else str.substring(charIdx + 1)
                val newStr = "${strBeforeCurChar}${strAfterCurChar}"

                permute(newStr, newPrefix, curPermutations)
            }
        } else {
            curPermutations.add(prefix)
        }
        return curPermutations
    }
}
```