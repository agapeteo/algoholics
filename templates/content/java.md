# Java

## Contents
---
- [Array](#array)
- [List interface](#list)
- [ArrayList - dynamic array](#dynamic-array)
- [LinkedList](#linked-list)
- [Map](#map)
- [Set](#set)
- [Queue](#queue)
- [Not a humble opinion about language](#pros-cons)

---

<div id="array"/>

## Array
create array:
```java
int[] a = new int[5];
int[] a = new int[]{1, 2, 3, 4, 5}; 
```


<div id="list"/>
## List interface
notable methods:

- `get(int index)` - _returns the element at the specified position in this list_

- `add(E e)` - _appends the specified element to the end of this list (optional operation)_

- `add(int index, E element)` - _inserts the specified element at the specified position in this list (optional operation)_

- `set(int index, E element)` - _replaces the element at the specified position in this list with the specified element (optional operation)_

- `remove(int index)` - _removes the element at the specified position in this list (optional operation)_

- `remove(Object o)` - _removes the first occurrence of the specified element from this list, if it is present (optional operation)_

- `size()` - _returns the number of elements in this list_

See more details about `List` interface and list of methods [here](https://docs.oracle.com/javase/8/docs/api/java/util/List.html)


<div id="dynamic-array"/>

## ArrayList - dynamic array
create array list:
```java
List<Integer> someList = new ArrayList();
```

<div id="linked-list"/>

## Linked list
create linked list:
```java
List<Integer> someList = new LinkedList();
```


<div id="map" />
 
## Map
> Starting `Java 8` worst case of `HashMap` changed from `O(n^2)` to `O(log(N))`. 
When Java runtime detects high hash collision it changed `linked list` to `tree` for collection of values. 
 
create map:
```java
Map<String, Integer> dict = new HashMap();
```
Other implementations:

- `TreeMap` (_implements `NavigableMap`_)
- `ConcurrentHashMap`
- `ConcurrentSkipListMap`
- `SynchronizedMap`
- `Hashtable` (_old, not used now_)

See more details about `Map` interface and list of methods [here](https://docs.oracle.com/javase/8/docs/api/java/util/Map.html)


<div id="set" />

## Set
```java
Set<Integer> s = new HashSet();
```
other implementations of `Set`:

- `TreeSet` implements `SortedSet` and `NavigableSet` interfaces
- `ConcurrentSkipListSet`
- `CopyOnWriteArraySet`
- `SynchronizedSet`
- `LinkedHashSet`

See more details about `Set` interface and list of methods [here](https://docs.oracle.com/javase/8/docs/api/java/util/Set.html)


<div id="queue" />

## Queue
interface `Queue` with multiple implementations, like: 

- `LinkedList` 
- `LinkedBlockingQueue`
- `ConcurrentLinkedQueue`
- `ConcurrentLinkedDeque`
- `DelayQueue`
- `LinkedTransferQueue`
- `PriorityQueue`
- `SynchronousQueue`


<div id="pros-cons" />

## Not a humble opinion about language
**pros**

- easy to find job and developers
- huge choice of libraries
- relatively fast
- platform independence
- can be used for multiple tasks 


**cons**

- verbose
- too big for simple tasks

