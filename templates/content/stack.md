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
interface `Dequeue` (**_extends `Queue`_**)

Implementation examples:

- `ArrayDeque`  
- `LinkedList` 
- `ArrayDeque`  
- `LinkedBlockingDeque`

Example of simple implementation of LinkedList in Java [here](/linkedlist.html#java)


Simplified implementation example:
<script src="https://gist.github.com/agapeteo/40e8764a5da58bb54b9a8501feedd725.js"></script>


<div id="go"/>

## Go
using slices:
```go
stack := make([]int, 0)
stack = append(stack, 1) // same as push()
stack = append(stack, 2)
stack = append(stack, 2)

// same as pop()
lastIdx := len(stack) - 1
topElement := stack[lastIdx]
stack = stack[:lastIdx]

fmt.Printf("topElement: %v, stack: %v", topElement, stack) // topElement: 2, stack: [1 2]
```
struct using slices:

> `TODO:`

struct using linked list:

> `TODO:`



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