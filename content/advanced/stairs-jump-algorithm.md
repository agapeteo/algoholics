+++
title = "Stairs jump count"
template = "level.html"
description = "Stairs jump count"
[extra]
    keywords = "Stairs jump count, dynamic programming, recursion, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Stairs jump algorithm

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
Recursive approach:
```java
package recursion;

public class StairJump {

    public static int waysCount(int stairsCount, int[] jumpWays) {
        if (stairsCount < 0) return 0;

        if (stairsCount == 0) return 1;

        int totalWaysCount = 0;
        for (int curJump : jumpWays) {
            totalWaysCount += waysCount(stairsCount - curJump, jumpWays);
        }
        return totalWaysCount;
    }
}

```

with caching (dynamic):
```java
package dynamic;

public class StairJump {
    public static int waysCount(int stairsCount, int[] waysCount) {
        int[] cache = new int[stairsCount + 1];
        return waysCount(stairsCount, waysCount, cache);
    }

    private static int waysCount(int stairsCount, int[] waysCount, int[] cache) {
        if (stairsCount < 0) return 0;
        if (stairsCount == 0) return 1;

        int totalWays = 0;
        for (int curJump : waysCount) {
            int curJumpIdx = stairsCount - curJump;
            if (curJumpIdx < 0) continue;

            int cachedResult = cache[curJumpIdx];
            if (cachedResult == 0) {
                cachedResult = waysCount(curJumpIdx, waysCount, cache);
                cache[curJumpIdx] = cachedResult;
            }
            totalWays += cachedResult;
        }
        return totalWays;
    }
}

```


<div id="go"/>

## Go
```go
package recursion

func StairJumpCount(stairsCount int, jumps []int) int {
	if stairsCount < 0 {
		return 0
	}

	if stairsCount == 0 {
		return 1
	}

	var total int
	for _, curJumps := range jumps {
		total += StairJumpCount(stairsCount-curJumps, jumps)
	}
	return total
}
```

with caching (dynamic):
```go
package dynamic

func StairJumpCount(stairsCount int, jumps []int) int {
	cache := make([]int, stairsCount+1, stairsCount+1)
	return stairJumpCount(stairsCount, jumps, cache)
}

func stairJumpCount(stairsCount int, jumps []int, cache []int) int {
	if stairsCount < 0 {
		return 0
	}
	if stairsCount == 0 {
		return 1
	}

	var totalWays int
	for _, curJump := range jumps {
		curJumpIdx := stairsCount - curJump
		if curJumpIdx < 0 {
			continue
		}
		curWays := cache[curJumpIdx]

		if curWays == 0 {
			curWays = stairJumpCount(curJumpIdx, jumps, cache)
			cache[curJumpIdx] = curWays
		}
		totalWays += curWays
	}
	return totalWays
}

```


<div id="python"/>

## Python
```python
def count_ways(stairs_count, jump_ways):
    if stairs_count < 0:
        return 0
    if stairs_count == 0:
        return 1

    total_ways = 0
    for cur_jump in jump_ways:
        total_ways += count_ways(stairs_count - cur_jump, jump_ways)

    return total_ways
```

with caching (dynamic):
```python
def _count_ways(stairs_count, jump_ways, cache):
    if stairs_count < 0:
        return 0

    if stairs_count == 0:
        return 1

    total_ways = 0
    for cur_jump in jump_ways:
        cur_ways_idx = stairs_count - cur_jump
        if cur_ways_idx < 0:
            continue

        cur_ways = cache[cur_ways_idx]
        if not cur_ways:
            cur_ways = _count_ways(cur_ways_idx, jump_ways, cache)
            cache[cur_ways_idx] = cur_ways

        total_ways += cur_ways

    return total_ways


def count_ways(stairs_count, jump_ways):
    cache = [None] * (stairs_count + 1)
    return _count_ways(stairs_count, jump_ways, cache)
```


<div id="javascript"/>

## JavaScript
Recursive approach:
```javascript
function waysCount(stairsCount, jumpWays) {
    if (stairsCount < 0) return 0;
    if (stairsCount == 0) return 1;

    let totalCount = 0;
    for (let curJump of jumpWays) {
        totalCount += waysCount(stairsCount - curJump, jumpWays);
    }
    return totalCount
}
```

with caching (dynamic):
```javascript
function waysCount(stairsCount, jumpWays, cache) {
    cache = Array(stairsCount + 1).fill(0)
    return waysCountWithCache(stairsCount, jumpWays, cache)
}

function waysCountWithCache(stairsCount, jumpWays, cache) {
    if (stairsCount < 0) return 0;
    if (stairsCount == 0) return 1;

    let totalCount = 0;
    for (let curJump of jumpWays) {
        const curWaysIdx = stairsCount - curJump;
        if (curWaysIdx < 0) {
            continue;
        }
        let curWays = cache[curWaysIdx];
        if (!curWays) {
            curWays = waysCountWithCache(curWaysIdx, jumpWays, cache);
            cache[curWaysIdx] = curWays;
        }
        totalCount += curWays
    }
    return totalCount
}
```




<div id="kotlin"/>

## Kotlin
Recursive approach:
```kotlin
package recursion

object StairJump {

    fun waysCount(stairsCount: Int, jumpWays: IntArray): Int {
        if (stairsCount < 0) return 0

        if (stairsCount == 0) return 1

        var totalCount = 0;

        for (jumpStairsWay in jumpWays) {
            totalCount += waysCount(stairsCount - jumpStairsWay, jumpWays)
        }
        return totalCount;
    }
}
```

with caching results (dynamic):
```kotlin
package dynamic

object StairJump {

    fun waysCount(stairsCount: Int, jumpWays: IntArray, cache: Array<Int?> = arrayOfNulls<Int?>(stairsCount + 1)): Int {
        if (stairsCount < 0) return 0

        if (stairsCount == 0) return 1

        var totalWays = 0

        for (curJump in jumpWays) {
            val curWaysCountIdx = stairsCount - curJump

            if (curWaysCountIdx < 0) continue

            var curJumpWays = cache[curWaysCountIdx]
            if (curJumpWays == null) {
                curJumpWays = waysCount(curWaysCountIdx, jumpWays, cache)
                cache[curWaysCountIdx] = curJumpWays
            }
            totalWays += curJumpWays
        }
        return totalWays
    }
}
```

<div id="rust"/>

## Rust
Recursive approach:

```rust
fn ways(stairs: i32, jump_ways: &[i32]) -> usize {
    if stairs < 0 {
        return 0;
    }
    if stairs == 0 {
        return 1;
    }
    let mut total_ways = 0;
    for cur_jump in jump_ways {
        total_ways += ways(stairs - *cur_jump, jump_ways);
    }
    total_ways
}
```

Dynamic:
```rust
use std::collections::HashMap;

fn ways_dynamic(stairs: i32, jump_ways: &[i32]) -> usize {
    let mut cache: HashMap<i32, usize> = HashMap::new();
    ways_cached(stairs, jump_ways, &mut cache)
}

fn ways_cached(stairs: i32, jump_ways: &[i32], cache: &mut HashMap<i32, usize>) -> usize {
    if stairs < 0 {
        return 0;
    }
    if stairs == 0 {
        return 1;
    }
    let mut total_ways = 0;
    for cur_jump in jump_ways {
        let cur_stairs = stairs - *cur_jump;
        if let Some(count) = cache.get(&cur_stairs) {
            total_ways += *count;
        } else {
            cache.insert(cur_stairs, ways(cur_stairs, jump_ways));
            total_ways += cache[&cur_stairs];
        }
    }
    total_ways
}
```
