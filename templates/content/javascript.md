# JavaScript

## Contents
---

- [Array - dynamic array](#array)
- [Map - associative array](#map)
- [Set](#set)

---

<div id="array"/>

## Array - dynamic array
JavaScript arrays are `dynamic arrays`.
It's implementation (`dense` or `sparse` arrays) depends on browser and may be defined in runtime.

- `Dense arrays` are like auto resizing arrays
- `Sparse arrays` are similar to `hash maps`


create array:
```javascript
var arr = [0, 1, 2];
var anotherArray = new Array("a", "b", "c");
```

Check more details about JS arrays [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array)
 
 
  
<div id="map"/>
 
## Map - associative array
JavaScript `objects` historically used as a maps. 
It's implementation also depends on browser engine, 
but most cases it look like hash based algorithm, so `get()` and `set()` is `O(1)`

`EcmaScript 6` (sometimes called `ES2015`) and later versions have `Map` data structure, which similar to JS objects, but has additional features.
You can read about `Map` and `objects` difference [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map#Objects_and_maps_compared)    


#### basic Map usage
```javascript
const m = new Map();
m.set("one", 1);
m.get("one") // 1

// using object 
const obj = {};
obj['one'] = 1;
obj.one; // 1
```


<div id="set"/>

## Set
```javascript
const s = new Set([1, 2, 3, 4, 5]);

console.log(s.has(1)); // true
```
 


