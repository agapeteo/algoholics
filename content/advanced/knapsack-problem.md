+++
title = "Knapsack problem"
template = "level.html"
description = "Knapsack problem"
[extra]
    keywords = "Knapsack problem, dynamic programming, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Knapsack problem

## Contents

---

- [Kotlin](#kotlin)

--- 


<div id="kotlin"/>

## Kotlin

```kotlin
package dynamic

import kotlin.math.max

object Knapsack {
    data class Item(val price: Int, val weight: Int)

    fun maxTotalPrice(inputItems: List<Item>, maxWeight: Int): Int {

        val items = withFirstZeroItem(inputItems)
        val choices = createTable(items.size, maxWeight + 1)

        for (curItemIdx in 1 until items.size) {
            for (curWeight in 0..maxWeight) {

                val curItem = items[curItemIdx]
                val prevBestPrice = choices[curItemIdx - 1][curWeight]

                val curBestPrice =
                        if (curItem.weight > curWeight) {
                            prevBestPrice
                        } else {
                            val prevBestPriceWithoutCurItemWeight = choices[curItemIdx - 1][curWeight - curItem.weight]
                            max(
                                    prevBestPrice,
                                    curItem.price + prevBestPriceWithoutCurItemWeight)
                        }
                choices[curItemIdx][curWeight] = curBestPrice
            }
        }
        return choices[items.size - 1][maxWeight]
    }

    private fun withFirstZeroItem(items: List<Item>): MutableList<Item> {
        val zeroItem = Item(0, 0)
        val result = mutableListOf(zeroItem)
        result.addAll(items)
        return result
    }

    private fun createTable(rowsCount: Int, columnsCount: Int): MutableList<MutableList<Int>> {
        val result = mutableListOf<MutableList<Int>>()

        repeat(rowsCount) {
            result.add(MutableList(columnsCount) { 0 })
        }
        return result
    }
}
```