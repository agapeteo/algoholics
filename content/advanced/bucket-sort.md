+++
title = "Bucket sort"
template = "level.html"
description = "Bucket sort"
[extra]
    keywords = "bucket sort, sorting, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Bucket sort

## Contents
---

- [Description](#description)
- [Kotlin](#kotlin)

---


<div id="description"/>

## Description
Unlike [counting sort](/basic/counting-sort) or [radix sort](/advanced/radix-sort) numbers with floating point can also be used.

Idea is to put numbers into buckets by intervals (`(number - minNumber) * bucketsCount / range`), sort buckets and collect buckets into result list 

Average complexity is` O(N)`, but worst case is `O(N^2)` when numbers have bad distribution.

This algorithm can be used together with other simple sorting algorithms like [insertion sort](/fundamental/insertion-sort) for sorting small buckets



<div id="kotlin"/>

## Kotlin

```kotlin
package sort

import kotlin.math.sqrt

object BucketSort {

    fun sort(numbers: MutableList<Double>): List<Double> {
        if (numbers.size < 2) return numbers

        val minNumber = numbers.min()!!
        val range = numbers.max()!! - minNumber
        val bucketsCount = sqrt(range).toInt()
        val buckets = emptyBuckets(bucketsCount)

        for (number in numbers) {
            val curBucketIdx = (number - minNumber) * bucketsCount / range
            buckets[curBucketIdx.toInt()].add(number)
        }

        val result = mutableListOf<Double>()

        for (bucket in buckets) {
            val sortedBucket = BucketSort.sort(bucket)
            result.addAll(sortedBucket)
        }
        return result
    }

    private fun emptyBuckets(count: Int): List<MutableList<Double>> {
        var buckets = mutableListOf<MutableList<Double>>()
        for (i in 0..count) {
            buckets.add(mutableListOf())
        }
        return buckets
    }
}
```
