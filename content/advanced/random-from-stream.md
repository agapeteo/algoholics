+++
title = "Random element from stream"
template = "level.html"
description = "Random element from stream"
[extra]
    keywords = "random, stream, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Random element from stream

## Contents
---

- [Kotlin](#kotlin)

---


<div id="kotlin"/>

## Kotlin

```kotlin
package random

import java.util.*
import java.util.stream.Stream

object RandomFromStream {

    fun <T> random(stream: Stream<T>): T? {
        val rnd = Random()

        var result: T? = null

        var i = 0
        stream.forEach { element ->
            i += 1
            if (i == 1) {
                result = element
                return@forEach
            }

            val rndIdx = 1 + rnd.nextInt(i)
            if (rndIdx == i) {
                result = element
            }
        }
        return result
    }
}
```
