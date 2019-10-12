# Topological sort

## Contents

---

- [Kotlin](#kotlin)

---


<div id="kotlin"/>

## Kotlin

```kotlin
package graph

object GandzhaTopologySort {
    fun <T> topologySort(vertices: List<T>, dependencies: List<Pair<T, T>>): List<T> {
        val graph = buildReverseGraph(vertices, dependencies)
        val dependenciesCount = countDependencies(graph)
        return toList(dependenciesCount)
    }

    private fun <T> countDependencies(graph: Map<T, List<T>>): Map<T, Int> {
        val result = mutableMapOf<T, Int>()

        for (vertex in graph.keys) {
            countFor(graph, vertex, mutableSetOf(), result)
        }
        return result
    }

    private fun <T> buildReverseGraph(vertices: List<T>, dependencies: List<Pair<T, T>>): Map<T, List<T>> {
        val result = mutableMapOf<T, MutableList<T>>()
        for (v in vertices) {
            result[v] = mutableListOf()
        }
        for (pair in dependencies) {
            result[pair.first]?.add(pair.second)
        }
        return result
    }

    private fun <T> countFor(graph: Map<T, List<T>>, vertex: T, visiting: MutableSet<T>, result: MutableMap<T, Int>): Int {
        visiting.add(vertex)
        var count = 0

        val dependencies = graph[vertex] ?: emptyList()
        for (dependency in dependencies) {
            require(!visiting.contains(dependency)) { "cycled dependencies detected" }
            count++
            count += result[dependency] ?: countFor(graph, dependency, visiting, result)
        }
        result[vertex] = count
        return count
    }

    private fun <T> toList(dependenciesCount: Map<T, Int>): List<T> {
        return dependenciesCount.toList()
                .sortedBy { it.second }
                .map { it.first }
    }
}
```
