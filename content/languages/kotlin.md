+++
title = "Kotlin"
template = "level.html"
description = "data structures and algorithms in Kotlin programming languages"
[extra]
    keywords = "Kotlin, explain, tutorial, data structures, algorithms, interview"
+++

# Kotlin

## Contents
---

- [Array](#array)
- [List](#dynamic-array)
- [Map](#map)
- [Set](#set)

---

> Kotlin can any data structures from Java. <br/>See more detailed Kotlin collections overview [here](https://kotlinlang.org/docs/reference/collections-overview.html)

<div id="array" />

## Array
```kotlin
val array = arrayOf(1, 2, 3, 4, 5) // intArrayOf() prefered for int types
println(Arrays.toString(array))
println(array.first())
println(array.last())
println(array.indexOf(1))
println(array[0])
println(array.get(0)) 
```


<div id="dynamic-array" />

## List
```kotlin
val aList = listOf(1, 2, 3) // immutable by default
val bList = mutableListOf("a", "b", "c") 
```

<div id="map" />

## Map
```kotlin
val m = mapOf("one" to 1, "two" to 2, "three" to 3, "four" to 1) // immutable by default
val mutableMap = mutableMapOf("one" to 1, "two" to 2, "three" to 3, "four" to 1)

```


<div id="set" />

## Set
```kotlin
val s = setOf("a", "b", "c")
```

