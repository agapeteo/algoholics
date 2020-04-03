+++
title = "Graph"
template = "level.html"
description = "Graph data structure"
[extra]
    keywords = "graph, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Graph

## Contents

---

- [Description](#description)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Description

> `TODO description` 



<div id="java"/>

## Java


```java
package graph;

import java.util.*;

public class Graph<E> {
    private final Map<E, Set<E>> edges;

    public Graph() {
        edges = new HashMap<>();
    }

    void addBothEdges(E from, E to) {
        addEdge(from, to);
        addEdge(to, from);
    }

    private void addEdge(E from, E to) {
        Set<E> currentEdges = edges.get(from);
        if (currentEdges == null) {
            currentEdges = new HashSet<>();
            edges.put(from, currentEdges);
        }
        currentEdges.add(to);
    }

    public List<E> elementsDfs() {
        List<E> resultList = new ArrayList<>();
        Set<E> visited = new HashSet<>();
        edges.keySet().forEach(v -> elementsDfs(v, resultList, visited));
        return resultList;
    }

    private void elementsDfs(E vertex, List<E> resultList, Set<E> visited) {
        if (visited.contains(vertex)) return;

        resultList.add(vertex);
        visited.add(vertex);

        edges.get(vertex).forEach(v -> elementsDfs(v, resultList, visited));
    }

    public List<E> elementsBfs() {
        List<E> resultList = new ArrayList<>();
        Set<E> visited = new HashSet<>();
        edges.keySet().forEach(v -> elementsBfs(v, resultList, visited));
        return resultList;
    }

    private void elementsBfs(E vertex, List<E> resultList, Set<E> visited) {
        Deque<E> queue = new LinkedList<>();
        queue.add(vertex);

        while (!queue.isEmpty()) {
            E eachVertex = queue.pollFirst();

            if (visited.contains(eachVertex)) continue;

            resultList.add(eachVertex);
            visited.add(eachVertex);

            edges.get(eachVertex).forEach(v -> queue.add(v));
        }
    }

    public boolean connected(E start, E target) {
        Set<E> visited = new HashSet<>();
        return connectedFromNode(start, target, visited);
    }

    private boolean connectedFromNode(E start, E target, Set<E> visited) {
        if (visited.contains(start)) return false;

        if (start == target) return true;

        visited.add(start);

        for (E vertex: edges.get(start)) {
            if (connectedFromNode(vertex, target, visited)) return true;
        }

        return false;
    }
}

```


<div id="go"/>

## Go

```go
package graph

type Graph struct {
	edges map[interface{}][]interface{}
}

func NewGraph() *Graph {
	g := new(Graph)
	g.edges = make(map[interface{}][]interface{})
	return g
}

func (g Graph) AddEdge(from interface{}, to interface{}) {
	edges := g.edges[from]
	newEdges := append(edges, to)
	g.edges[from] = newEdges
}

func (g Graph) AddBothEdges(from interface{}, to interface{}) {
	g.AddEdge(from, to)
	g.AddEdge(to, from)
}

func (g Graph) ElementsDfs() []interface{} {
	visited := make(map[interface{}]struct{})
	result := make([]interface{}, 0)

	for k, _ := range g.edges {
		g.appendElementsDfs(k, &visited, &result)
	}
	return result
}

func (g Graph) IsConnected(from interface{}, to interface{}) bool {
	visited := make(map[interface{}]struct{})
	return g.isConnectedFrom(from, to, visited)
}

func (g Graph) isConnectedFrom(from interface{}, to interface{}, visited map[interface{}]struct{}) bool {
	_, isVisited := visited[from];
	if isVisited {
		return false
	}

	if from == to {
		return true
	}

	visited[from] = struct{}{}

	for _, conn := range g.edges[from] {
		if g.isConnectedFrom(conn, to, visited) {
			return true
		}
	}
	return false
}

func (g Graph) ElementsBfs() []interface{} {
	visited := make(map[interface{}]struct{})
	result := make([]interface{}, 0)

	for k, _ := range g.edges {
		queue := make([]interface{}, 0)
		queue = append(queue, k)

		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]

			_, isVisited := visited[k];
			if isVisited {
				continue
			}
			result = append(result, first)
			visited[first] = struct{}{}

			for _, e := range g.edges[first] {
				queue = append(queue, e)
			}
		}
	}
	return result
}

func (g Graph) appendElementsDfs(from interface{}, visited *map[interface{}]struct{}, list *[]interface{}) {
	if _, isVisited := (*visited)[from]; isVisited {
		return
	}
	*list = append(*list, from)
	(*visited)[from] = struct{}{}

	for _, v := range g.edges[from] {
		g.appendElementsDfs(v, visited, list)
	}
}
```



<div id="python"/>

## Python

```python
class Graph:
    def __init__(self):
        self.edges = {}

    def add_edge(self, input_from, input_to):
        connections = self.edges.get(input_from)

        if not connections:
            connections = set()
        connections.add(input_to)

        self.edges[input_from] = connections

    def add_both_edges(self, input_from, input_to):
        self.add_edge(input_from, input_to)
        self.add_edge(input_to, input_from)

    def elements_dfs(self):
        result_list = []
        visited = set()

        for element in self.edges.keys():
            self.add_dfs_elements(element, result_list, visited)

        return result_list

    def add_dfs_elements(self, from_element, result_list, visited):
        if from_element in visited:
            return

        result_list.append(from_element)
        visited.add(from_element)

        for e in self.edges[from_element]:
            self.add_dfs_elements(e, result_list, visited)

    def elements_bfs(self):
        visited = set()
        result_list = []

        for k in self.edges.keys():
            queue = [k]
            while len(queue) > 0:
                first_element = queue.pop(0)
                if first_element in visited:
                    continue
                result_list.append(first_element)
                visited.add(first_element)

                for vertex in self.edges[first_element]:
                    queue.append(vertex)
        return result_list

    def is_connected(self, from_element, to_element):
        return self._is_connected(from_element, to_element, set())

    def _is_connected(self, from_element, to_element, visited):
        if from_element in visited:
            return False

        if from_element == to_element:
            return True

        visited.add(from_element)

        for vertex in self.edges[from_element]:
            if self._is_connected(vertex, to_element, visited):
                return True

        return False
```



<div id="javascript"/>

## JavaScript
```javascript
class Graph {
  constructor() {
    this.edges = new Map()
    this.elementsDfs=this.elementsDfs.bind(this)
  }

  addEdge(from_element, to_element) {
    let connections = this.edges[from_element]
    if (!connections) {
      connections = []
    }
    connections.push(to_element)

    this.edges[from_element] = connections
  }

  addBothEdges(from_element, to_element) {
    this.addEdge(from_element, to_element)
    this.addEdge(to_element, from_element)
  }

  elementsDfs() {
    const visited = new Set()
    const result = []
   
    for (const key in this.edges) {       
      this.addElementsDfsFrom(key, visited, result)
    } 
    return result
  }

  addElementsDfsFrom(from_element, visited, result) {    
    if (visited.has(from_element)) return

    result.push(from_element)
    visited.add(from_element)

    for (const k of this.edges[from_element]) {
      this.addElementsDfsFrom(k, visited, result)
    }
  }

  elementsBfs() {
    const visited = new Set()
    const result = []
   
    for (const key in this.edges) {       
      const queue = []
      queue.push(key)
      while (queue.length > 0) {
        first = queue.shift()
        if (visited.has(first)) continue
        result.push(first)
        visited.add(first)
        for (const k of this.edges[first]) {
          queue.push(k)
        }
      }
    } 
    return result
  }

  isConnected(fromElement, toElement) {
    const visited = new Set()
    return this.isConnectedFrom(fromElement, toElement, visited)
  }

  isConnectedFrom(fromElement, toElement, visited) {
    if (visited.has(fromElement)) return false

    if (fromElement == toElement) return true
    
    visited.add(fromElement)

    for (vertex of this.edges[fromElement]) {      
      if (this.isConnectedFrom(vertex, toElement, visited)) {
        return true
      }
    }
    return false
  }
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package graph

import java.util.*
import kotlin.collections.HashSet


class Graph<E> {

    val edges: MutableMap<E, HashSet<E>> = mutableMapOf()

    fun addEdge(from: E, to: E) {
        edges.computeIfAbsent(from) { HashSet() }.add(to)
    }

    fun addBothEdges(from: E, to: E) {
        addEdge(from, to)
        addEdge(to, from)
    }

    fun elementsDfs(): List<E> {
        val visited = mutableSetOf<E>()

        fun addToResult(element: E, list: MutableList<E>) {
            if (visited.contains(element)) return

            visited.add(element)
            list.add(element)
            edges[element]?.forEach { addToResult(it, list) }
        }

        val result = mutableListOf<E>()
        edges.keys.forEach { addToResult(it, result) }
        return result
    }

    fun elementsBfs(): List<E> {
        val result = mutableListOf<E>()
        val visited = mutableSetOf<E>()

        edges.keys.forEach { topElement ->
            val queue: Deque<E> = LinkedList()

            queue.add(topElement)
            while (queue.isNotEmpty()) {
                val element = queue.removeFirst()
                if (!visited.contains(element)) {
                    visited.add(element)
                    result.add(element)
                    edges[element]?.forEach { queue.add(it) }
                }
            }
        }
        return result
    }

    fun connected(from: E, to: E): Boolean { // using bfs
        if (from == to) return true

        val alreadyVisited = mutableSetOf<E>().apply { add(from) }
        val queue: Deque<E> = LinkedList()

        edges[from]?.forEach { queue.add(it) }

        while (queue.isNotEmpty()) {
            val element = queue.removeFirst()

            if (element == to) return true

            if (!alreadyVisited.contains(element)) {
                alreadyVisited.add(element)
                edges[element]?.forEach { queue.add(it) }
            }
        }
        return false
    }
}

```



<div id="rust"/>

## Rust

```rust
use std::collections::{HashMap, HashSet, VecDeque};
use std::hash::Hash;

// expecting all edges references in the same lifetime scope
pub struct Graph<'a, T: Eq + Hash> {
    edges: HashMap<&'a T, HashSet<&'a T>>
}

impl<'a, T: Eq + Hash> Graph<'a, T> {
    pub fn new() -> Self {
        let edges: HashMap<&'a T, HashSet<&'a T>> = HashMap::new();
        Graph { edges }
    }

    pub fn add_both_edges(&mut self, from: &'a T, to: &'a T) {
        self.add_edge(from, to);
        self.add_edge(to, from);
    }

    pub fn add_edge(&mut self, from: &'a T, to: &'a T) {
        let connections_maybe = self.edges.get_mut(&from);
        match connections_maybe {
            None => {
                let mut new_connections: HashSet<&'a T> = HashSet::new();
                new_connections.insert(to);
                self.edges.insert(from, new_connections);
            }
            Some(connections) => {
                connections.insert(to);
            }
        }
    }

    pub fn collect_elements_dfs(&self) -> Vec<&'a T> {
        let mut visited: HashSet<&'a T> = HashSet::new();
        let mut elements: Vec<&'a T> = Vec::new();

        for k in self.edges.keys() {
            self.collect_dfs(k, &mut visited, &mut elements);
        }

        elements
    }

    pub fn collect_dfs(&self, elem: &'a T, visited: &mut HashSet<&'a T>, elements: &mut Vec<&'a T>) {
        if visited.contains(&elem) {
            return;
        }
        elements.push(elem);
        visited.insert(elem);

        let connection_maybe = self.edges.get(elem);
        if let Some(connections) = connection_maybe {
            for e in connections {
                self.collect_dfs(e, visited, elements);
            }
        }
    }

    pub fn collect_elements_bfs(&self) -> Vec<&T> {
        let mut elements: Vec<&'a T> = Vec::new();
        let mut visited: HashSet<&'a T> = HashSet::new();

        for elem in self.edges.keys() {
            let mut queue: VecDeque<&T> = VecDeque::new();
            queue.push_front(elem);
            while let Some(cur_element) = queue.pop_front() {
                if visited.contains(cur_element) {
                    continue;
                }
                elements.push(cur_element);
                visited.insert(cur_element);
                if let Some(connections) = self.edges.get(cur_element) {
                    for other_element in connections {
                        queue.push_back(other_element);
                    }
                }
            }
        }
        elements
    }

    pub fn connected(&self, from: &'a T, to: &'a T) -> bool {
        let mut visited: HashSet<&T> = HashSet::new();
        let mut queue: VecDeque<&T> = VecDeque::new();

        queue.push_front(from);

        while let Some(cur_element) = queue.pop_front() {
            if visited.contains(cur_element) {
                continue;
            }

            if *to == *cur_element {
                return true;
            }

            visited.insert(cur_element);

            if let Some(connections) = self.edges.get(cur_element) {
                for other_element in connections {
                    queue.push_back(other_element);
                }
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn collect_dfs() {
        let mut graph: Graph<&str> = Graph::new();
        let alex = "Alex";
        let bob = "Bob";
        let michael = "Michael";
        let john = "John";
        let sara = "Sara";

        graph.add_both_edges(&alex, &bob);
        graph.add_edge(&bob, &michael);
        graph.add_edge(&bob, &john);
        graph.add_edge(&john, &sara);
        graph.add_edge(&sara, &john);

        let mut elements = graph.collect_elements_dfs();
        elements.sort();

        assert_eq!(elements, vec![&alex, &bob, &john, &michael, &sara]);
    }

    #[test]
    fn collect_bfs() {
        let mut graph: Graph<&str> = Graph::new();
        let alex = "Alex";
        let bob = "Bob";
        let michael = "Michael";
        let john = "John";
        let sara = "Sara";

        graph.add_both_edges(&alex, &bob);
        graph.add_edge(&bob, &michael);
        graph.add_edge(&bob, &john);
        graph.add_edge(&john, &sara);
        graph.add_edge(&sara, &john);

        let mut elements = graph.collect_elements_bfs();
        elements.sort();

        assert_eq!(elements, vec![&alex, &bob, &john, &michael, &sara]);
    }

    #[test]
    fn check_connected() {
        let mut graph: Graph<&str> = Graph::new();
        let alex = "Alex";
        let bob = "Bob";
        let michael = "Michael";
        let john = "John";
        let sara = "Sara";

        assert_eq!(graph.connected(&alex, &bob), false);

        graph.add_both_edges(&alex, &bob);


        assert_eq!(graph.connected(&alex, &bob), true);

        graph.add_edge(&bob, &michael);
        assert_eq!(graph.connected(&alex, &michael), true);

        graph.add_both_edges(&john, &sara);
        assert_eq!(graph.connected(&john, &sara), true);
        assert_eq!(graph.connected(&alex, &sara), false);

        graph.add_edge(&michael, &john);
        assert_eq!(graph.connected(&alex, &sara), true);
    }
}
```