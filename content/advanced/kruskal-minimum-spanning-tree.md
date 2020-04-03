+++
title = "Kruskal minimum spanning tree"
template = "level.html"
description = "Kruskal minimum spanning tree"
[extra]
    keywords = "Kruskal minimum spanning tree, MST, graph, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Minimum Spanning Tree

## Contents

---

- [Minimum Spanning Tree](#minimum-spanning-tree)
  - [Contents](#contents)
  - [Java](#java)

---


<div id="java"/>

## Java

```java
package io.algoholics.graph.mst;

import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

public class DisjointSet<T> {
    private final Map<T, T> parents;
    private final Map<T, Integer> heights;
    private int count;

    public DisjointSet() {
        this(null);
    }

    public DisjointSet(Collection<T> elements) {
        if (null == elements || elements.isEmpty()) {
            parents = new HashMap<>();
            heights = new HashMap<>();
            count = 0;
        } else {
            parents = elements.stream().collect(Collectors.toMap(p -> p, p -> p));
            heights = elements.stream().collect(Collectors.toMap(p -> p, p -> 0));
            count = elements.size();
        }
    }

    public int getCount() {
        return count;
    }

    public T find(T element) {
        if (!parents.containsKey(element)) {
            parents.put(element, element);
            heights.put(element, 0);
            count++;
            return element;
        }

        while (element != parents.get(element)) {
            parents.put(element, parents.get(parents.get(element)));
            element = parents.get(element);
        }
        return element;
    }

    public boolean connected(T element1, T element2) {
        return find(element1) == find(element2);
    }

    public boolean union(T element1, T element2) {
        final T rootElement1 = find(element1);
        final T rootElement2 = find(element2);

        if (rootElement1 == rootElement2)
            return false;

        if (heights.get(rootElement1) < heights.get(rootElement2)) {
            parents.put(rootElement1, rootElement2);
        } else if (heights.get(rootElement1) > heights.get(rootElement2)) {
            parents.put(rootElement2, rootElement1);
        } else {
            parents.put(rootElement2, rootElement1);
            heights.merge(rootElement1, 1, Integer::sum);
        }

        count--;
        return true;
    }
}
```

```java
package io.algoholics.graph.mst;

import java.util.Objects;
import java.util.StringJoiner;

public class Edge<T> {
    private final T source;
    private final T target;
    private final double weight;

    public Edge(T source, T target, double weight) {
        this.source = source;
        this.target = target;
        this.weight = weight;
    }

    public T getSource() {
        return source;
    }

    public T getTarget() {
        return target;
    }

    public double getWeight() {
        return weight;
    }

    public T getOther(T vertex) {
        if (source.equals(vertex)) {
            return target;
        }
        if (target.equals(vertex)) {
            return source;
        }
        throw new IllegalArgumentException(String.format("No vertex %s in the edge %s", vertex, this));
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Edge<?> edge = (Edge<?>) o;
        return Double.compare(edge.weight, weight) == 0 &&
                Objects.equals(source, edge.source) &&
                Objects.equals(target, edge.target);
    }

    @Override
    public int hashCode() {
        return Objects.hash(source, target, weight);
    }

    @Override
    public String toString() {
        return new StringJoiner(", ", Edge.class.getSimpleName() + "[", "]")
                .add("source=" + source)
                .add("target=" + target)
                .add("weight=" + weight)
                .toString();
    }
}
```

```java
package io.algoholics.graph.mst;

import java.util.*;
import java.util.stream.Collectors;

public class Graph<T> {
    private final Map<T, List<Edge<T>>> adjacencyList;
    private int edgesCount;

    public Graph(final Map<T, List<Edge<T>>> adjacencyList) {
        this.adjacencyList = adjacencyList;
        this.edgesCount = getEdges().size();
    }

    public Graph() {
        this(new HashMap<>());
    }

    public void addEdge(final Edge<T> edge) {
        adjacencyList.merge(edge.getSource(), new ArrayList<>(Collections.singletonList(edge)), (oldVal, newVal) -> {
            oldVal.addAll(newVal);
            return oldVal;
        });
        adjacencyList.merge(edge.getTarget(), new ArrayList<>(Collections.singletonList(edge)), (oldVal, newVal) -> {
            oldVal.addAll(newVal);
            return oldVal;
        });
        edgesCount++;
    }

    public List<Edge<T>> getEdges() {
        return adjacencyList.entrySet().stream().flatMap(p -> p.getValue().stream()).distinct().collect(Collectors.toList());
    }

    public List<Edge<T>> getEdges(T vertex) {
        return new ArrayList<>(adjacencyList.get(vertex));
    }

    public List<T> getVertices() {
        return new ArrayList<>(adjacencyList.keySet());
    }
}
```

```java
package io.algoholics.graph.mst;

import java.util.Comparator;
import java.util.HashSet;
import java.util.PriorityQueue;
import java.util.Set;

public class KuskalMst {

    private static final class EdgeComparator<T> implements Comparator<Edge<T>> {

        @Override
        public int compare(Edge<T> t1, Edge<T> t2) {
            return Double.compare(t1.getWeight(), t2.getWeight());
        }
    }

    public <T> Set<Edge<T>> mst(final Graph<T> graph) {
        final Set<Edge<T>> mst = new HashSet<>();
        final DisjointSet<T> connected = new DisjointSet<>(graph.getVertices());
        final int vertexCount = connected.getCount();
        final PriorityQueue<Edge<T>> pq = new PriorityQueue<>(new EdgeComparator<>());

        pq.addAll(graph.getEdges());

        while (!pq.isEmpty() && mst.size() < vertexCount - 1) {
            final Edge<T> edge = pq.poll();
            if (connected.union(edge.getSource(), edge.getTarget())) {
                mst.add(edge);
            }
        }

        return mst;
    }
}
```
