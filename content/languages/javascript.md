+++
title = "JavaScript"
template = "level.html"
description = "data structures and algorithms in JavaScript programming languages"
[extra]
    keywords = "JavaScript, explain, tutorial, data structures, algorithms, interview"
+++

# JavaScript

## Contents
---

- [Array](#array)
- [Map - associative array](#map)
- [Set](#set)

---

<div id="array"/>

## JavaScript Array 
There two array implementation approaches in JavaScript:

- `Dense arrays` are like auto resizing arrays (dynamic array)
- `Sparse arrays` are more similar to `hash maps`

JavaScript engines choose array implementation in run-time in most cases

create array example:
```javascript
var arr = [0, 1, 2];
var anotherArray = new Array("a", "b", "c");
```

create empty array of specified length:
```javascript
const arrLength = 5;
const arr = Array(arrLength).fill({});
// or
const arr = Array(arrLength).fill(null);
```

Notable array methods:

method | description | returns | note
--- | --- | --- | ---
`push()` | add element(s) to the end  | new array length | same as `push()` in `Stack` and `enqueue()` in `Queue`
`pop()` | remove last element | removed element | same as `pop()` in `Stack`
`shift()` | remove first element | removed element | changes the length of the array  
`unshift()` | add element(s) to the beginning of array | added element | 
`slice()` | create slice of array based on start and end index | shallow copy of array | original array won't be changed. <br/> `["a", "b", "c"].slice(1) \\ [b, c]` <br/> `["a", "b", "c"].slice(1, 2) \\ [b]`  
`splice()` | in place changes the contents of an array by removing or replacing existing elements and/or adding new elements | array containing the deleted elements |  `arr = ["a", "c"];` <br/> `x = arr.splice(1, 0, "b")` <br/> `["a", "b", "c"]`  
`every()` | check if **all** values satisfy provided test function (predicate)  | Boolean |    
`some()` | check if **at least one** of values satisfy provided test function (predicate)  | Boolean |    
`sort()` | in place sort  | sorted array |  compares sequences of UTF-16 code units values  
`filter()` | filters by test function (predicate)  | new filtered array copy |  original won't be changed  
`reduce()` | applies **reduce** function on each element | result of reduction |  `const sum = [1, 2, 3].reduce((sum, n) => sum += n)`
`forEach()` | applies provided function on each element | undefined |  `[1,2,3].forEach(i => console.log(i))`
`map()` | creates a new array with the results of applying provided function on each element | new array  |  `[1,2,3].map(n => n * 2) \\ [2,4,6]`
 
  


Check more details about JavaScript arrays [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array)
 
 
  
<div id="map"/>
 
## Map - associative array
JavaScript `objects` historically used as a maps. 
It's implementation also depends on browser engine, 
but most cases it look like hash based algorithm, so `get()` and `set()` is `O(1)`

`EcmaScript 6` (sometimes called `ES2015`) and later versions have `Map` data structure, which similar to JS objects, but has additional features.
You can read about `Map` and `objects` difference [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map)    


#### basic Map usage
```javascript
const m = new Map();
m.set("one", 1); // set valur
m.get("one") // 1 - read value
m.has("one") // true - check if key exist
m.keys()

// using object 
const obj = {};
obj['one'] = 1;
obj.one; // 1
```
see more details about [JavaScript Map here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map)

<div id="set"/>

## Set
```javascript
const s = new Set([1, 2, 3, 4, 5]);

console.log(s.has(1)); // true
```
 
