+++
title = "Coin change count"
template = "level.html"
description = "Coin change count"
[extra]
    keywords = "Coin change count, dynamic programming, recursion, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Coin change count

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
package dynamic;

import java.util.List;

public class Coins {
    public static int changeWays(int amount, List<Integer> denoms) {
        return changeWays(amount, denoms, 0);
    }

    private static int changeWays(int amount, List<Integer> denoms, int denomIdx) {
        if (denomIdx >= denoms.size() - 1) {
            return 1;
        }

        int totalWays = 0;
        int curDenomAmount = denoms.get(denomIdx);

        for (int i = 0; i * curDenomAmount <= amount; i++) {
            int remainingAmount = amount - i * curDenomAmount;
            totalWays += changeWays(remainingAmount, denoms, denomIdx + 1);
        }

        return totalWays;
    }
}
```


<div id="go"/>

## Go
```go
package dynamic

func CoinChangeCount(amount int, denoms []int) int {
	return coinChangeCount(amount, denoms, 0)
}

func coinChangeCount(amount int, denoms []int, denomIdx int) int {
	if denomIdx >= len(denoms)-1 {
		return 1
	}

	totalWays := 0
	curDenomAmount := denoms[denomIdx]

	for i := 0; i*curDenomAmount <= amount; i++ {
		remainingAmount := amount - i*curDenomAmount
		totalWays += coinChangeCount(remainingAmount, denoms, denomIdx+1)
	}

	return totalWays
}
```


<div id="python"/>

## Python
```python
def coin_change_count(amount, denoms, denom_idx=0):
    if denom_idx >= len(denoms) - 1:
        return 1

    total_ways = 0
    cur_demom_amount = denoms[denom_idx]

    i = 0
    while i * cur_demom_amount <= amount:
        remaining_amount = amount - i * cur_demom_amount
        total_ways += coin_change_count(remaining_amount, denoms, denom_idx + 1)
        i += 1

    return total_ways
```



<div id="javascript"/>

## JavaScript

```javascript
function coinChangeCount(amount, denoms) {
    return _coinChangeCount(amount, denoms, 0)
}

function _coinChangeCount(amount, denoms, denomIdx) {
    if (denomIdx >= denoms.length - 1) {
        return 1
    }

    let totalWays = 0;
    const denomAmount = denoms[denomIdx];

    for (let i = 0; i * denomAmount <= amount; i++) {
        const remainingAmount = amount - i * denomAmount;
        totalWays += _coinChangeCount(remainingAmount, denoms, denomIdx + 1);
    }
    return totalWays
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package dynamic

object Coins {
    fun changeWays(amount: Int, denoms: List<Int>, denomIdx: Int = 0): Int {

        if (denomIdx >= denoms.size - 1) {
            return 1
        }

        val denomAmount = denoms[denomIdx]
        var ways = 0

        var i = 0
        while (i * denomAmount <= amount) {
            val remainingAmount = amount - i * denomAmount
            ways += changeWays(remainingAmount, denoms, denomIdx + 1)
            i++
        }
        return ways
    }
}
```



<div id="rust"/>

## Rust

```rust
fn coin_change_count(amount: usize, denoms: &[usize]) -> usize {
    fn count(amount: usize, denoms: &[usize], denom_idx: usize) -> usize {
        if denom_idx >= denoms.len() - 1 {
            return 1;
        }
        let mut total_ways = 0;
        let cur_denom_amount = denoms[denom_idx];

        let mut i = 0;
        while i * cur_denom_amount <= amount {
            let remaining_amount = amount - i * cur_denom_amount;
            total_ways += count(remaining_amount, denoms, denom_idx + 1);
            i += 1;
        }
        total_ways
    }
    count(amount, denoms, 0)
}

#[test]
fn test() {
    assert_eq!(coin_change_count(100, &[50, 20]), 3);
    assert_eq!(coin_change_count(5, &[1, 2, 3]), 12);
}
```
