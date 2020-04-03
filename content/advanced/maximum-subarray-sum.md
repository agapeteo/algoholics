+++
title = "Maximum subarray sum, Kadane's algorithm&"
template = "level.html"
description = "Maximum subarray sum, Kadane's algorithm"
[extra]
    keywords = "Maximum subarray sum, kadane's algorithm, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Maximum subarray sum

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
package arrays;

import java.util.List;

public class Kadane {
    public static int maxSum(List<Integer> list) {
        int result = 0;
        int curSum = 0;

        for (int i = 0; i < list.size(); i++) {
            if (curSum + list.get(i) > 0) {
                curSum += list.get(i);
                if (curSum > result) {
                    result = curSum;
                }
            } else {
                curSum = 0;
            }
        }
        return result;
    }
}
```

<div id="go"/>

## Go

```go
package arrays

func maxSum(s []int) int {
	result := 0
	curSum := 0

	for i, _ := range s {
		if curSum+s[i] > 0 {
			curSum += s[i]
			if curSum > result {
				result = curSum
			}
		} else {
			curSum = 0
		}
	}
	return result
}
```


<div id="python"/>

## Python
```python
# From wiki page


def max_subarray(numbers):
    best_sum = 0  # or: float('-inf')
    best_start = best_end = 0  # or: None
    current_sum = 0
    for current_end, x in enumerate(numbers):
        if current_sum <= 0:
            # Start a new sequence at the current element
            current_start = current_end
            current_sum = x
        else:
            # Extend the existing sequence with the current element
            current_sum += x

        if current_sum > best_sum:
            best_sum = current_sum
            best_start = current_start
            best_end = current_end + 1  # the +1 is to make 'best_end' exclusive

    return best_sum, best_start, best_end

```


<div id="javascript"/>

## JavaScript

```javascript
function maxSum(arr) {
    let result = 0;
    let curSum = 0;

    for (let i = 0; i < arr.length; i++) {
        if (curSum + arr[i] > 0) {
            curSum += arr[i];

            if (curSum > result) {
                result = curSum;
            }
        } else {
            curSum = 0;
        }
    }
    return result;
}
```


<div id="kotlin"/>

## Kotlin

```kotlin
package arrays

fun maxSum(list: List<Int>): Int {
    var result = 0
    var curSum = 0

    for (i in list.indices) {
        if (curSum + list[i] > 0) {
            curSum += list[i]
            if (curSum > result) {
                result = curSum
            }
        } else {
            curSum = 0
        }
    }
    return result
}
```


<div id="rust"/>

## Rust

```rust
fn max_sum(slice: &[isize]) -> isize {
    let mut result = 0;
    let mut cur_sum = 0;

    for i in 0..slice.len() {
        if cur_sum + slice[i] > 0 {
            cur_sum += slice[i];
            if cur_sum > result {
                result = cur_sum;
            }
        } else {
            cur_sum = 0;
        }
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        let arr = [-1, 2, 3, -7, 3, 4];
        assert_eq!(max_sum(&arr), 7);
    }

    #[test]
    fn test_2() {
        let arr = [34, -50, 42, 14, -5, 86];
        assert_eq!(max_sum(&arr), 137);
    }

    #[test]
    fn all_negative() {
        let arr = [-3, -2, -1, -4];
        assert_eq!(max_sum(&arr), 0);
    }

    #[test]
    fn all_positive() {
        let arr = [1, 1, 0, 1, 1];
        assert_eq!(max_sum(&arr), 4);
    }
}
```

