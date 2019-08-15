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
queue = queue[1:]

fmt.Printf("firstElement: %v, queue: %v", firstElement, queue) // firstElement: 1, queue: [2 3]
```
<br/>

creating `Queue` struct based on slices (`[]int`):
```go
package main

import (
	"fmt"
)

type Queue []int

func (q *Queue) Enqueue(element int) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		firstElement := (*q)[0]
		*q = (*q)[1:]
		return firstElement, true
	}
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	firstElement, _ := queue.Dequeue()

	fmt.Printf("firstElement: %v, queue: %v", firstElement, *queue) // firstElement: 1, queue: [2 3]
}
```

<br/>

See `Queue` struct using linked list in [linked list page for Go](/linkedlist.html#go)

> Note. If you need to use linked list or stack/queue based linked list in production prefer using [Go's container/list] (https://golang.org/pkg/container/list/)


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