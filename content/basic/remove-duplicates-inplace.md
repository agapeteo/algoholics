+++
title = "Remove duplicates inplace"
template = "level.html"
description = "Remove duplicates inplace"
[extra]
    keywords = "remove duplicates, deduplication, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Remove duplicates in-place

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

import java.util.Arrays;

public class InPlaceDuplicateRemoval {

    public static int remove(Object[] arr) {
        Arrays.sort(arr);
        int i = 0;
        for (int j = 1; j < arr.length; j++) {
            if (!arr[j].equals(arr[i])) {
                arr[i + 1] = arr[j];
                i++;
            }
        }
        return i;
    }
}
```


<div id="go"/>

## Go

```go
package arrays

import "sort"

func RemoveDuplicates(arr []string) int {
	sort.Strings(arr)
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return i
}

```


<div id="python"/>

## Python

```python
def inplace_deduplicate(in_list):
    in_list.sort()
    i = 0
    for j in range(1, len(in_list)):
        if in_list[j] != in_list[i]:
            in_list[i + 1] = in_list[j]
            i = i + 1
    return i
```


<div id="javascript"/>

## JavaScript

```javascript
const inPlaceRemoval = arr => {
  arr.sort()
  let i = 0
  for (j = 1; j < arr.length; j++) {
    if (arr[j] != arr[i]) {
      arr[i + 1] = arr[j]
      i++
    }
  }
  return i
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package arrays

object InPlaceDuplicateRemoval {
    fun <T> remove(arr: Array<T>): Int {
        arr.sort()
        var i = 0
        for (j in 1 until arr.size) {
            if (arr[j] != arr[i]) {
                arr[i + 1] = arr[j]
                i++
            }
        }
        return i
    }
}
```


<div id="rust"/>

## Rust
```rust
fn dedup<T: Ord>(slice: &mut [T]) -> usize {
    slice.sort();

    let mut i = 0;

    for j in i + 1..slice.len() {
        if slice[i] != slice[j] {
            i += 1;
            slice.swap(i, j);
        }
    }
    i + 1
}

#[cfg(test)]
mod tests {
    use super::dedup;

    #[test]
    fn all_unique() {
        let mut arr = [5, 4, 2, 1, 0];

        let last_idx = dedup(&mut arr);

        assert_eq!(last_idx, arr.len());
    }

    #[test]
    fn all_same() {
        let mut arr = ["a", "a", "a", "a"];

        let last_idx = dedup(&mut arr);

        assert_eq!(last_idx, 1);
    }

    #[test]
    fn some_dups_2() {
        let mut arr = ["a", "b", "a", "a"];

        let last_idx = dedup(&mut arr);

        assert_eq!(last_idx, 2);
    }

    #[test]
    fn some_dups_5() {
        let mut arr = ["a", "b", "a", "b", "c", "d", "d", "e"];

        let last_idx = dedup(&mut arr);

        assert_eq!(last_idx, 5);
    }
}
```
