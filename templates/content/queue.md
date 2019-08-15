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
interface `Queue` with multiple implementations, like: 

- `LinkedList` 
- `LinkedBlockingQueue`
- `ConcurrentLinkedQueue`
- `ConcurrentLinkedDeque`
- `DelayQueue`
- `LinkedTransferQueue`
- `PriorityQueue`
- `SynchronousQueue`


Simplified implementation example:
<script src="https://gist.github.com/agapeteo/40e8764a5da58bb54b9a8501feedd725.js"></script>


<div id="go"/>

## Go
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

struct using slices:

> `TODO:`

struct using linked list:

> `TODO:`


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

Example of simplified version in Kotlin:
<script src="https://gist.github.com/agapeteo/15d80043e9b1454a7d57e102ea849ef4.js"></script>