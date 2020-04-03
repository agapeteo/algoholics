+++
title = "Python"
template = "level.html"
description = "data structures and algorithms in Python programming languages"
[extra]
    keywords = "python, explain, tutorial, data structures, algorithms, interview"
+++

# Python

## Contents
---

- [Tuple - array](#array)
- [List - dynamic array](#dynamic-array)
- [Dictionary - map, associative array](#map)
- [Set](#set)

---
<div id="array" />

## Tuple - array `( )`
create tuple: 
```python
t = ("a", "b", "c")
```

tuple packing and unpacking:
```python
t = 1,2,3
a, b, c = t
```


<div id="dynamic-array" />

## List - dynamic array `[ ]`
Python's lists can be used as `Stack` and `Queue`

create list:
```python
l = [1, 2, 3]
```

create empty list of specified size:
```python
list_size = 12
some_list = [None] * list_size
```

List of notable methods:

method | description | returns | note | __________example__________ | big O
--- | --- | --- | --- | --- | ---
`append(element)` | add element to the end | `None` |  | `[0, 1].append(2)` | `O(1)`
`extend(iterable)` | joins other list in place | `None` | same as `[1,2] + [3,4] ` <br/> but `extend` is in place | `[0, 1].extend([2, 3])` | `O(k)`
`insert (idx, element)` | inserts element at given index | `None` | in place | `[1, 3].insert(1, 2)`  <br/> inserts 2 between 1 and 3 | `O(N)`
`pop(idx)` | removes element and index and returns that value | removed value |  default idx is `-1` (last) | `[1, 2].pop()` | `O(N)` <br/> but `O(1)` for last element 
`remove(value)`| removes value from list| `None` | raises error if value not found | `[1, 2].remove[2]` | `O(N)`
`reverse()`| reverses list | `None` | in place | `[0, 1].reverse()` | `O(N)`
`sort()` | sorts list <br/> in place  | `None` | optionally specify sort key and/or reverse order, defaults are: (`key=None, reverse=False`).<br/> Algorithm used: [Timsort](https://en.wikipedia.org/wiki/Timsort) | `[0, 2, 1].sort()` | `O(log(N))`
`index(value)` | finds fist index of element | found index |  raises `ValueError` if not found | `[0, 1].index(1)` | `O(N)`
`count(element)` | returns the number of times given element appears in list | number of repeats | | `["a", "b", "a"].count("a")` | `O(N)`


See other build-in functions like `len()`, `max()`, `cmp()`, `filter()`, `map()`, `any()`, `all()`, `range()` etc [here at official Python documentation](https://docs.python.org/3/library/functions.html)


<div id="map" />

## Dictionary - map, associative array `{ key: value }`
create dictionary: 
```python
d = {"one": 1, "two": 2}  # create dict with initialized data

some_dict = {}  # create empty dict

some_dict["three"] = 3  # set/update value

try:
    print some_dict["four"]  # read/get value from dict, throws KeyError if given key is absent
except KeyError:
    print "key does not exist int this dictionary"

print "one" in some_dict  # check if dict contains key, O(1)
```

Another options: 

- [UserDict](https://docs.python.org/3/library/collections.html#collections.UserDict)
- [OrderedDict](https://docs.python.org/3/library/collections.html#collections.OrderedDict)
- [ChainMap](https://docs.python.org/3/library/collections.html#collections.ChainMap)



<div id="set" />

## Set  `{ }`
create set:
```python
s = {1, 2, 3}

another_set = set()
another_set.add("a")  # add value
another_set.update("c", "d")  # adding one or more elements
another_set = another_set | {"b"}  # joining two sets (union)
another_set.remove("c")  # remove element
print another_set  # set(['a', 'b', 'd'])

print "'b' found" if "b" in another_set else "'b' is not found"  # check if set contains value
```

