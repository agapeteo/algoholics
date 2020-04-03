+++
title = "Queue"
template = "level.html"
description = "Queue"
[extra]
    keywords = "queue, deque, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Queue

## Contents
---

- [Queue description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---

<div id="description"/>

## Description
Core methods:

- `enqueue()` - _add elements to the end of the queue_ 
- `dequeue()` - _delete and return element from the beginning of the queue_

> ``//TODO: description with images``


<div id="practice"/>

## In Practice 
> ``//TODO: ``



<div id="java"/>

## Java
See Java's `Queue` description and methods [here at official documentation](https://docs.oracle.com/javase/8/docs/api/java/util/Queue.html)

interface `Queue` with multiple implementations, like: 

- `LinkedList` 
- `LinkedBlockingQueue`
- `ConcurrentLinkedQueue`
- `ConcurrentLinkedDeque`
- `DelayQueue`
- `LinkedTransferQueue`
- `PriorityQueue`
- `SynchronousQueue`

<br/>
See example of implementation [in linked list page for java](/linkedlist.html#java):



<div id="go"/>

## Go
using slice:
```go
queue := make([]int, 0)
queue = append(queue, 1) // same as enqueue()
queue = append(queue, 2)
queue = append(queue, 3)

// same as dequeue()
firstElement := queue[0]
queue = queue[1:] // WARNING: potential memory leak

fmt.Printf("firstElement: %v, queue: %v", firstElement, queue) // firstElement: 1, queue: [2 3]
```
<br/>

`Queue` struct based on `container/list`:
```go
package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	elements *list.List
}

func NewQueue() *Queue {
	q := new(Queue)
	q.elements = list.New()
	return q
}

func (q *Queue) Enqueue(e interface{}) {
	q.elements.PushBack(e)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if q.elements.Len() == 0 {
		return nil, true
	}
	e := q.elements.Front().Value
	q.elements.Remove(q.elements.Front())
	return e, false
}

func (q *Queue) Elements() []interface{} {
	c := make([]interface{}, q.elements.Len())

	for i, e := 0, q.elements.Front(); e != nil; i, e = i+1, e.Next() {
		c[i] = e.Value
	}
	return c
}

func (q *Queue) Len() int {
	return q.elements.Len()
}

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Dequeue()

	fmt.Printf("value %v, all elements: %v", v, q.Elements()) // value 1, all elements: [2 3]
}
```

<br/>

Other queue implementations:

- [deque implementation based on container.list with improved performance](https://github.com/juju/utils/blob/master/deque/deque.go)
- [based on slices no memory leaks](https://github.com/phf/go-queue/blob/master/queue/queue.go) 


<div id="python"/>

## Python

using build-in list methods:
```python
queue = []

queue.append(1)  # same as `enqueue()`, adds 1 to the end of queue
queue.append(2)
queue.append(3)

x = queue.pop(0)  # same as 'dequeue()', deletes and returns element at the beginning of queue

print x  # 1
print queue  # [2, 3]
```

<br/>

Also see Python's [collections.deque](https://docs.python.org/3/library/collections.html#collections.deque) which is implemented based on `Linked list` and has methods for `Stack` and `Queue`

<div id="javascript"/>

## JavaScript
using Arrays's methods:
```javascript
const queue = []
queue.push(1) // adds 1 to the end of queue
queue.push(2)
queue.push(3)
const value = queue.shift() //deletes and returns element at the beginning of queue

console.log(value) // 1
console.log(queue) // [2, 3]
```

<div id="kotlin"/>

## Kotlin
Kotlin can use `Queue` from [Java](#java)

See example of simple `LinkedList` in Kotlin which implements `Stack` and `Queue` in [Linked List page](/linkedlist.html#kotlin)


## Rust
See example [Queue using VecDeque](/languages/rust#queue)