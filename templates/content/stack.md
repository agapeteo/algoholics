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
stack = stack[:lastIdx] // WARNING! potential memory leak

fmt.Printf("topElement: %v, stack: %v", topElement, stack) // topElement: 3, stack: [1 2]
```

<br/>

`Stack` struct based on `container/list`:
```go
package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	elements *list.List
}

func NewStack() *Stack {
	s := new(Stack)
	s.elements = list.New()
	return s
}

func (s *Stack) Push(e interface{}) {
	s.elements.PushBack(e)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.elements.Len() == 0 {
		return nil, true
	}
	e := s.elements.Back().Value
	s.elements.Remove(s.elements.Back())
	return e, false
}

func (s *Stack) Elements() []interface{} {
	c := make([]interface{}, s.elements.Len())

	for i, e := 0, s.elements.Front(); e != nil; i, e = i+1, e.Next() {
		c[i] = e.Value
	}
	return c
}

func (s *Stack) Len() int {
	return s.elements.Len()
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, _ := s.Pop()

	fmt.Printf("value %v, all elements: %v", v, s.Elements()) // value 3, all elements: [1 2]
}
```

<br/>
See also deque implementation [based on container.list with improved performance](https://github.com/juju/utils/blob/master/deque/deque.go)


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

<br/>
Also see Python's [collections.deque](https://docs.python.org/3/library/collections.html#collections.deque) which is implemented based on `Linked list` and has methods for `Stack` and `Queue`


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

See example of simple `LinkedList` in Kotlin which implements `Stack` and `Queue` in [Linked List page](/linkedlist.html#kotlin)