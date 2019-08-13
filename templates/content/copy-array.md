# Copy array

## Contents

---

- [Copy array description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="description" />

## Copy array

Complexity `O(N)` - requires iterating through all elements in collection



<div id="practice"/>

## In practice
Used very often. Dynamic array creates new array and copied data into new array for growing and shrinking


<div id="java"/>

## Java
There few way to copy array in Java:

- `Object.clone()` - to create new same copy of array 
- `Arrays.copyOf()`  
- `Arrays.copyOfRange()`
- `System.arraycopy()` 
- using streams API: `String[] newArray = Arrays.stream(srcArray).toArray(String[]::new)`

> `Arrays.copyOf()` and `Arrays.copyOfRange()` are using `System.arraycopy()` inside.
> 
>`System.arraycopy()` is very fast native implementation  which is used very often in core Java


System.arraycopy() signature:

`arraycopy(Object srcArray, int srcPos, Object destArray, int destPos, int length)`

- **`srcArray`** - source array
- **`srcPos`** - source initial position
- **`destArray`** - destination array 
- **`destPos`** - destination initial position
- **`length`** - how many elements to copy from source to destination

Example for creating twice bigger array and copy data into new array:

```java
int[] srcArr = {1, 2, 3};
int[] destArr = new int[srcArr.length * 2];
System.arraycopy(srcArr, 0, destArr, 0, srcArr.length); // [1, 2, 3, 0, 0, 0]
```

Brute-force approach (same size):
```java
int[] srcArr = {1, 2, 3};
int[] destArr = new int[srcArr.length];

for (int i = 0; i < srcArr.length; i++) {
    destArr[i] = srcArr[i];
}
System.out.println(Arrays.toString(destArr));
```



<div id="go"/>

## Go
copy into new slice which is twice bigger:
```go
srcArr := []int{0, 1, 2}
destArr := make([]int, len(srcArr) * 2)
copy(destArr, srcArr) // pay attention: destination is first argument

fmt.Printf("new array: %v", destArr) // new array: [0 1 2 0 0 0]
```

Brute-force approach (same size):
```go
srcArr := []int{0, 1, 2}
destArr := make([]int, len(srcArr))

for idx, v := range srcArr {
    destArr[idx] = v
}

fmt.Printf("new array: %v", destArr)

```




<div id="python"/>

## Python
There few ways to copy list in Python:

- `dest_list = list(src_list)`
- `dest_list = src_list[:]`
- `dest_list = src_list.copy()` (since Python 3.3 )
- for copy and deep copy:
 
```python
import copy
dest_list = copy.deepcopy(src_list)
```

Brute-force approach (same size):
```python
src_list = [1, 2, 3]
dst_list = []

for n in src_list:
    dst_list.append(n)

print dst_list
```


<div id="javascript"/>

## JavaScript
There few ways to copy list in JavaScript:

using `slice()`:
```javascript
const srcArr = [1, 2, 3]
const destArr = srcArr.slice()
```

using `Array.from()`:
```javascript
const srcArr = [1, 2, 3]
const destArr = Array.from(srcArr)
```

ES6 way `[...]`:
```javascript
const srcArr = [1, 2, 3]
const destArr = [...srcArr]
```

deep copy using JSON:
```javascript
const srcArr = [[1, 2], 3]
const destArr = JSON.parse(JSON.stringify(srcArr))
```

Brute-force approach (same size):
```javascript
srcArr = [1, 2, 3]
destArr = []
srcArr.forEach(n => {destArr.push(n)})

console.log(destArr)
```


<div id="kotlin"/>

## Kotlin
Can use Java's `System.arraycopy()`

Few other simple ways:

- `val destArr = arrayOf(srcArr)`
- `val destArr = srcArr.copyOf()`
- `val destArr = srcArr.clone()` 


Brute-force approach (same size):
```kotlin
val srcArr = arrayOf(1, 2, 3)
val destArr = IntArray(srcArr.size)

for ((idx, value) in srcArr.withIndex()) {
    destArr[idx] = srcArr[idx]
}
println("destArr is: ${Arrays.toString(destArr)}")
```

> You can use `arrayOfNulls<T>(size)` for any custom type of arrays, 
>
>For example: `arrayOfNulls<String>(10)` 





