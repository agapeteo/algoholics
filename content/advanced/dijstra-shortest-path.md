+++
title = "Shortest path"
template = "level.html"
description = "Shortest path in weighted graph (Dijstra)"
[extra]
    keywords = "shortest path, dijstra, graph, weighted graph, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Dijstra's shortest path

## Contents
---

- [Description](#description)
- [Kotlin](#kotlin)

---


<div id="description"/>

## Description
Dijstra's shortest path algorithms require edge weights to be >= 0 (not negative). 

This algorithm uses greedy approach and it is actually optimal.

There are multiple ways to implement this algorithms. Most common 2 approaches:
- using simple queue (`O(V^2)`)
- using priority queue (`O(E+V log V)`) which is better for sparse graphs 



<div id="kotlin"/>

## Kotlin

Using priority queue:

```kotlin
package graph

import java.util.*
import kotlin.collections.HashSet

data class EdgeTo<V>(val to: V, val weight: Double)

class WeightedGraph<V> {

    val edges: MutableMap<V, HashSet<EdgeTo<V>>> = mutableMapOf()

    fun addEdge(from: V, to: V, weight: Double) {
        edges.computeIfAbsent(from) { HashSet() }.add(EdgeTo(to, weight))
    }

    fun addBothEdges(from: V, to: V, weight: Double) {
        addEdge(from, to, weight)
        addEdge(to, from, weight)
    }
}

class ShortestPath<V>(private val graph: WeightedGraph<V>) {

    private data class VertexRank<V>(val vertex: V, var curWeight: Double, var prev: V?, var visited: Boolean = false)

    private val ranksComparator = kotlin.Comparator { stat1: VertexRank<V>, stat2: VertexRank<V> ->
        when {
            stat1.curWeight == stat2.curWeight -> 0
            stat1.curWeight < stat2.curWeight -> -1
            else -> 1
        }
    }

    fun minPath(start: V, end: V): List<V> {
        val ranks = mutableMapOf<V, VertexRank<V>>()
        val minQueue = PriorityQueue(ranksComparator)

        val startRank = VertexRank(start, 0.0, null)
        ranks[start] = startRank

        minQueue.add(startRank)

        while (minQueue.isNotEmpty()) {
            val curRank = minQueue.poll()

            if (curRank.visited) continue

            curRank.visited = true

            val neighborEdges = graph.edges[curRank.vertex] ?: continue

            for (neighborEdge in neighborEdges) {
                var neighborRank = ranks[neighborEdge.to]

                if (neighborRank != null) {
                    if (ranks[neighborRank.vertex]!!.visited) continue

                    if (curRank.curWeight + neighborEdge.weight < neighborRank.curWeight) {
                        neighborRank.curWeight = curRank.curWeight + neighborEdge.weight
                        neighborRank.prev = curRank.vertex
                    }
                } else {
                    neighborRank = VertexRank(neighborEdge.to, curRank.curWeight + neighborEdge.weight, curRank.vertex)
                    ranks[neighborRank.vertex] = neighborRank
                }
                minQueue.add(neighborRank)
            }
        }
        return extractPath(end, ranks)
    }

    private fun extractPath(end: V, ranks: Map<V, VertexRank<V>>): List<V> {
        val result = mutableListOf(end)

        var prev = ranks[end]?.prev
        while (prev != null) {
            result.add(prev)
            prev = ranks[prev]?.prev
        }
        result.reverse()
        return result
    }
}
```
