+++
title = "Sieve of eratosthenes"
template = "level.html"
description = "Find prime numbers using sieve of eratosthenes"
[extra]
    keywords = "prime numbers, sieve of eratosthenes, eratosthenes, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Sieve of Eratosthenes

## Contents

---

- [Find prime numbers using sieve of Eratosthenes approach](#description)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Sieve of Eratosthenes

Find prime numbers using sieve of Eratosthenes 



<div id="java"/>

## Java
```java
package math;

import java.util.ArrayList;
import java.util.List;

public class EratosthenesSievePrime {

    public static List<Integer> primes(int limit) {
        boolean[] arr = new boolean[limit];

        for (int i = 2; i < Math.sqrt(limit); i++) {
            if (arr[i]) continue;
            for (int j = i + i; j < limit; j += i) {
                arr[j] = true;
            }
        }

        List<Integer> result = new ArrayList<>();
        for (int k = 2; k < limit; k++) {
            if (!arr[k]) {
                result.add(k);
            }
        }
        return result;
    }
}
```


<div id="go"/>

## Go

```go
package math

func Primes(lim int) []int {
	s := make([]bool, lim)

	for i := 2; i*i < lim; i++ {
		if s[i] {
			continue
		}
		for j := i + i; j < lim; j += i {
			s[j] = true // true == number is not prime
		}
	}
	return extractPrimes(s)
}

func extractPrimes(s []bool) []int {
	res := make([]int, 0)
	for num, notPrime := range s {
		if num == 0 || num == 1 {
			continue
		}
		if !notPrime {
			res = append(res, num)
		}
	}
	return res
}
```


<div id="python"/>

## Python

```python
import math


def primes(limit):
    sieve_list = [False] * limit
    for i in range(2, math.floor(math.sqrt(limit))):
        if sieve_list[i]:
            continue

        for j in range(i + i, limit, i):
            sieve_list[j] = True

    result = []
    for k in range(2, limit):
        if not sieve_list[k]:
            result.append(k)
    return result
```



<div id="javascript"/>

## JavaScript

```javascript
const primes = limit => {
    const arr = Array(limit).fill(false);
    for (let i = 2; i * i < limit; i++) {
        if (arr[i]) continue;

        for (let j = i + i; j < limit; j += i) {
            arr[j] = true;
        }
    }
    const result = [];
    for (let k = 2; k < limit; k++) {
        if (!arr[k]) {
            result.push(k);
        }
    }
    return result;
};
```



<div id="kotlin"/>

## Kotlin

```kotlin
package math

import kotlin.math.sqrt

object EratosthenesSievePrimes {

    fun primes(limit: Int): List<Int> {
        val arr = BooleanArray(limit)

        for (i in 2 until sqrt(limit.toDouble()).toInt()) {
            if (arr[i]) continue

            var j = i + i
            while (j < limit) {
                arr[j] = true
                j += i
            }
        }
        val result = mutableListOf<Int>()
        for ((num, notPrime) in arr.withIndex()) {
            if (num == 0 || num == 1) continue
            if (!notPrime) {
                result.add(num)
            }
        }
        return result
    }
}
```


<div id="rust"/>

# Rust
```rust
use std::collections::HashSet;

fn primes(num: usize) -> Vec<usize> {
    let mut not_primes: HashSet<usize> = HashSet::new();
    let mut result: Vec<usize> = Vec::new();

    for i in 2..=num {
        if not_primes.contains(&i) {
            continue
        }
        result.push(i);

        let mut j = i;
        while j <= num {
            j += i;
            not_primes.insert(j);
        }
    }
    result
}

#[test]
fn test() {
    let primes = primes(100);
    assert_eq!(primes, 
               [2, 3, 5, 7, 
                   11, 13, 17, 19, 
                   23, 29, 
                   31, 37, 
                   41, 43, 47, 
                   53, 59, 
                   61, 67, 
                   71, 73, 79, 
                   83, 89, 
                   97]);
}
```
