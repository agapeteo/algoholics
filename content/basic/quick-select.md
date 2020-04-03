+++
title = "Quick select"
template = "level.html"
description = "Quick select"
[extra]
    keywords = "quick select, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Quick select

## Contents

---

- [Description](#description)
- [Complexity](#complexity)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---

<div id="description"/>

## Description
Finds smallest or biggest k element



<div id="complexity"/>


## Complexity
Time - `O(N)`

Space - `O(1)`
 

<div id="practice"/>




<div id="java"/>

## Java

```java
package sort;

import java.util.List;

public class QuickSelect {

    public static <T extends Comparable<T>> T top(int topIdx, List<T> list) {
        return top(topIdx, list, 0, list.size() - 1);
    }

    private static <T extends Comparable<T>> T top(int topIdx, List<T> list, int lowIdx, int hiIdx) {
        T pivotValue = list.get(hiIdx);
        int i = lowIdx;
        for (int j = i; j < hiIdx; j++) {
            if (list.get(j).compareTo(pivotValue) <= 0) {
                swap(i, j, list);
                i++;
            }
        }
        swap(i, hiIdx, list);

        if (i != topIdx) {
            if (topIdx < i) {
                return top(topIdx, list, lowIdx, i - 1);
            }
            return top(topIdx, list, i + 1, hiIdx);
        }
        return list.get(i);
    }

    private static <T extends Comparable<T>> void swap(int i, int j, List<T> list) {
        T tmp = list.get(i);
        list.set(i, list.get(j));
        list.set(j, tmp);
    }
}
```



<div id="go"/>

## Go
```go
package sort

type Comparable interface {
	Compare(comparable Comparable) int
}

func QuickSelect(topIdx int, s []Comparable, lowIdx int, hiIdx int) Comparable {
	pivotValue := s[hiIdx]
	i := lowIdx

	for j := i; j < hiIdx; j++ {
		if s[j].Compare(pivotValue) <= 0 {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[hiIdx] = s[hiIdx], s[i]

	if topIdx != i {
		if topIdx < i {
			return QuickSelect(topIdx, s, lowIdx, i-1)
		}
		return QuickSelect(topIdx, s, i+1, hiIdx)
	}
	return s[i]
}
```


<div id="python"/>

## Python

```python
def _quick_select(top_idx, in_list, low_idx, hi_idx):
    pivot_value = in_list[hi_idx]
    i = low_idx

    for j in range(i, hi_idx):
        if in_list[j] <= pivot_value:
            in_list[i], in_list[j] = in_list[j], in_list[i]
            i += 1
    in_list[i], in_list[hi_idx] = in_list[hi_idx], in_list[i]

    if top_idx != i:
        if top_idx < i:
            return _quick_select(top_idx, in_list, low_idx, i - 1)
        return _quick_select(top_idx, in_list, i + 1, hi_idx)

    return in_list[i]


def quick_select(top_idx, in_list):
    return _quick_select(top_idx, in_list, 0, len(in_list) - 1)
```


<div id="javascript"/>

## JavaScript

```javascript
function quickSelect(topIdx, array) {
    return select(topIdx, array, 0, array.length - 1)
}

function select(topIdx, array, lowIdx, hiIdx) {
    const pivotValue = array[hiIdx];
    let i = lowIdx;

    for (let j = i; j < hiIdx; j++) {
        if (array[j] <= pivotValue) {
            [array[i], array[j]] = [array[j], array[i]];
            i++;
        }
    }
    [array[i], array[hiIdx]] = [array[hiIdx], array[i]];

    if (topIdx !== i) {
        if (topIdx < i) {
            return select(topIdx, array, lowIdx, i - 1);
        }
        return select(topIdx, array, i + 1, hiIdx);
    }
    return array[i];
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package sort

object QuickSelect {
    fun <T : Comparable<T>> top(topIdx: Int, list: MutableList<T>, lowIdx: Int = 0, hiIdx: Int = list.size - 1): T {
        fun swap(i: Int, j: Int) {
            val tmp = list[i]
            list[i] = list[j]
            list[j] = tmp
        }

        var i = lowIdx
        val pivotValue = list[hiIdx]
        for (j in i until hiIdx) {
            if (list[j] <= pivotValue) {
                swap(i, j)
                i++
            }
        }
        swap(hiIdx, i)

        if (i != topIdx) {
            if (topIdx < i) {
                return top(topIdx, list, lowIdx, i - 1) // go left
            }
            return top(topIdx, list, i + 1, topIdx) // go right
        }
        return list[i]
    }
}
```



<div id="kotlin"/>

## Rust
```rust
fn top<T: Ord>(top_idx: usize, slice: &mut [T]) -> &T {
    top_in_range(top_idx, slice, 0, slice.len() - 1)
}

fn top_in_range<T: Ord>(top_idx: usize, slice: &mut [T], low_idx: usize, hi_idx: usize) -> &T {
    let mut i = low_idx;

    for j in i..hi_idx {
        if slice[j] <= slice[hi_idx] {
            slice.swap(i, j);
            i += 1;
        }
    }
    slice.swap(i, hi_idx);

    if i != top_idx {
        return if top_idx > i {
            top_in_range(top_idx, slice, i + 1, hi_idx)
        } else {
            top_in_range(top_idx, slice, low_idx, i - 1)
        };
    }
    &slice[i]
}

#[cfg(test)]
mod test {
    use super::top;

    #[test]
    fn random_number() {
        let mut arr = [2, 1, 44, 5, 12];

        assert_eq!(top(0, &mut arr), &1);
        assert_eq!(top(1, &mut arr), &2);
        assert_eq!(top(2, &mut arr), &5);
        assert_eq!(top(3, &mut arr), &12);
        assert_eq!(top(4, &mut arr), &44);
    }

    #[test]
    fn ascending() {
        const LIMIT: usize = 100;
        let mut arr: Vec<usize> = (0..LIMIT).collect();

        for i in 0..LIMIT {
            assert_eq!(top(i, &mut arr), &i);
        }
    }

    #[test]
    fn descending() {
        const LIMIT: usize = 100;
        let mut arr: Vec<usize> = (0..LIMIT).rev().collect();

        for i in 0..LIMIT {
            assert_eq!(top(i, &mut arr), &i);
        }
    }
}
```
