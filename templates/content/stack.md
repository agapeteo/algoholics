# Stack

## Contents
---

- [Stack description](#description)
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

- `push()` - _add elements to the top (end) of the stack_ 
- `pop()` - _delete and return element at the top (end) of the stack_

Both complexities are `O(1)`

> ``//TODO: description with images``


<div id="practice"/>

## In Practice 
> ``//TODO: ``


<div id="java"/>

## Java
Interface `Dequeue` (**_which extends `Queue`_**) has all methods required for `Stack`

See more details and list of methods [here at official documentation web page](https://docs.oracle.com/javase/8/docs/api/java/util/Deque.html)

Implementation examples:

- `LinkedList` 
- `ArrayDeque`  
- `LinkedBlockingDeque`

<br/>
See example of implementation [in linked list page for java](/linkedlist.html#java):



<div id="go"/>

## Go
using slices:
```go
stack := make([]int, 0)
stack = append(stack, 1) // same as push()
stack = append(stack, 2)
stack = append(stack, 3)

// same as pop()
lastIdx := len(stack) - 1
topElement := stack[lastIdx]
stack = stack[:lastIdx]

fmt.Printf("topElement: %v, stack: %v", topElement, stack) // topElement: 3, stack: [1 2]
```

<br/>

`Stack` struct using slices (`[]int`)
```go
package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		lastIdx := len(*s) - 1
		topElement := (*s)[lastIdx]
		*s = (*s)[:lastIdx]
		return topElement, true
	}
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topElement, _ := stack.Pop()

	fmt.Printf("topElement: %v, stack: %v", topElement, *stack) // topElement: 3, stack: [1 2]
}
```



<br/>
linked list example, which also implements stack look at [linked list page, Go](/linkedlist.html#go)

> Note. If you need to use linked list or stack/queue based linked list in production prefer using [Go's container/list] (https://golang.org/pkg/container/list/)


<div id="python"/>

## Python
using build-in list methods:

```python
stack = []

stack.append(1)  # same as `push()`, adds 1 to the top of stack (end of list)
stack.append(2)
stack.append(3)

x = stack.pop()  # deletes and returns element at the top of stack (end of list)

print x  # 3
print stack  # [1, 2]
```

<div id="javascript"/>

## JavaScript
using Arrays's methods:
```javascript
const stack = []
stack.push(1) // adds 1 to the top of stack (end of array)
stack.push(2)
stack.push(3)
const value = stack.pop() //deletes and returns element at the top of stack (end of array)

console.log(value) // 3
console.log(stack) // [1, 2]
```



<div id="kotlin"/>

## Kotlin
Kotlin can use `Dequeue` from [Java](#java)

Example of simplified version in Kotlin:
<script src="https://gist.github.com/agapeteo/15d80043e9b1454a7d57e102ea849ef4.js"></script>