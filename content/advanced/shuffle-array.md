+++
title = "Shuffle array"
template = "level.html"
description = "Shuffle array"
[extra]
    keywords = "random, shuffle, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Shuffle array

## Contents
---

- [Kotlin](#kotlin)

---



<div id="kotlin"/>

## Kotlin

```kotlin
package random

import java.util.*

object Shuffle {

    private fun <T> swap(list: MutableList<T>, i: Int, j: Int) {
        val tmp = list[i]
        list[i] = list[j]
        list[j] = tmp
    }

    fun <T> shuffle(list: MutableList<T>) {
        val rnd = Random()
        var j = list.size - 1

        while (j > 0) {
            val rndIdx = rnd.nextInt(j + 1)

            swap(list, j, rndIdx)

            j -= 1
        }
    }
}
```
