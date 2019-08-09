# Java

## Contents
---
- [Array](#array)
- [ArrayList - dynamic array](#dynamic-array)
- [LinkedList](#linked-list)
- [Map](#map)
- [Set](#set)
- [Queue](#queue)
- [Stack](#stack)
- [Not a humble opinion about language](#pros-cons)

---

<div id="array"/>

## Array
create array:
```java
int[] a = new int[5];
int[] a = new int[]{1, 2, 3, 4, 5}; 
```

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


<div id="stack" />

## Stack
interface `Dequeue` (**_extends `Queue`_**)

Implementation examples:

- `ArrayDeque`  
- `LinkedList` 
- `ArrayDeque`  
- `LinkedBlockingDeque`

Example of simple implementation of LinkedList in Java [here](/linkedlist.html#java)
<br/>


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

