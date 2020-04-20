+++
title = "Radix sort"
template = "level.html"
description = "Radix sort"
[extra]
    keywords = "radix sort, sorting, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Radix sort

## Contents
---

- [Description](#description)
- [Kotlin](#kotlin)

---


<div id="description"/>

## Description
Very old algorithm used to sort punch cards. Idea is to sort number by each digit level (from lowest to highest or opposite way).
Complexity is `O(N*d)`, where `d` is number of digits in numbers.



<div id="kotlin"/>

## Kotlin

```kotlin
package sort

import kotlin.math.pow

object RadixSort {
    private const val BASE = 10

    fun sort(list: List<Int>): List<Int> {
        var result = list

        if (list.size < 2) return result

        val digitsCount = countDigits(list.max()!!)

        for (digitLevel in 0 until digitsCount) {
            result = sortAtLevel(result, digitLevel)
        }
        return result
    }

    private fun countDigits(maxNumber: Int): Int {
        var digits = 0
        var number = maxNumber

        while (number != 0) {
            number /= 10
            digits += 1
        }
        return digits
    }

    private fun sortAtLevel(numbers: List<Int>, digitLevel: Int): List<Int> {
        var buckets = tenEmptyDigitBuckets()

        for (number in numbers) {
            val bucketIdx = (number / tenInPow(digitLevel)) % BASE

            buckets[bucketIdx].add(number)
        }

        val result = mutableListOf<Int>()
        for (bucket in buckets) {
            result.addAll(bucket)
        }
        return result
    }

    private fun tenInPow(digitLevel: Int): Int {
        return (BASE.toDouble()).pow(digitLevel).toInt()
    }

    private fun tenEmptyDigitBuckets(): List<MutableList<Int>> {
        var buckets = mutableListOf<MutableList<Int>>()
        for (i in 0..BASE) {
            buckets.add(mutableListOf())
        }
        return buckets
    }
}
```
