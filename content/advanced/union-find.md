+++
title = "Disjoint set"
template = "level.html"
description = "Disjoint Set (union find)"
[extra]
    keywords = "Disjoint Set, union find, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Disjoint Set (union find)

## Contents

---

- [Kotlin](#kotlin)

--- 


<div id="kotlin"/>

## Kotlin

interface:
```kotlin
package graph

interface DisjointSet<E> {
    fun union(left: E, right: E);

    fun find(elem: E): E  // result root element of set

    fun connected(left: E, right: E): Boolean
}
```

implementation:
```kotlin
package graph

class SimpleDisjointSet<E>(
        private val sizes: MutableMap<E, Int> = mutableMapOf(),
        private val map: MutableMap<E, E> = mutableMapOf()) : DisjointSet<E> {

    override fun union(left: E, right: E) {
        if (handleSame(left, right)) return

        val curLeft = map[left]
        val curRight = map[right]

        if (curLeft == null && curRight == null) {
            joinBoth(left, right)
            return
        }

        if (curLeft == null || curRight == null) {
            joinMissing(left, right)
            return
        }

        merge(left, right)
    }

    private fun joinBoth(left: E, right: E) {
        map[left] = left
        map[right] = left
        sizes[left] = 2
    }

    private fun joinMissing(left: E, right: E) {
        val missing = if (map[left] == null) left else right
        val existing = if (left == missing) right else left

        val curRoot = rootFor(existing)
        map[missing] = curRoot
        sizes[curRoot] = sizes[curRoot]!! + 1
    }

    private fun merge(left: E, right: E) {
        val leftRoot = rootFor(left)
        val rightRoot = rootFor(right)

        if (leftRoot == rightRoot) return

        val leftRootSize = sizes[leftRoot]!!
        val rightRootSize = sizes[rightRoot]!!
        val smallerRoot = if (leftRootSize < rightRootSize) leftRoot else rightRoot
        val biggerRoot = if (leftRoot == smallerRoot) rightRoot else leftRoot

        map[smallerRoot] = biggerRoot
        sizes[biggerRoot] = sizes[biggerRoot]!! + sizes[smallerRoot]!!
    }

    private fun rootFor(elem: E): E {
        val value = map[elem]

        if (value == elem) return elem

        return rootFor(value!!)
    }

    private fun handleSame(left: E, right: E): Boolean {
        if (left != right) {
            return false
        }
        if (left == null) {
            map[left] = left
            sizes[left] = 1
        }
        return true
    }

    override fun find(elem: E): E {
        return rootFor(elem)
    }

    override fun connected(left: E, right: E): Boolean {
        return find(left) == find(right)
    }
}
```