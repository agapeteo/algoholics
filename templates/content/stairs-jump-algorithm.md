# Stairs jump algorithm

## Contents

---

- [Complexity](#complexity)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [Kotlin](#kotlin)

---


<div id="complexity"/>


## Complexity
`TODO`

 

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

caching result approach (dynamic programming):
```kotlin
```