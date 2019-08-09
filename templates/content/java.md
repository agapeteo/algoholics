# Java

### create array
```java
int[] a = new int[5];
int[] a = new int[]{1, 2, 3, 4, 5}; 
```
### dynamic array - ArrayList
```java
List<Integer> someList = new ArrayList();
```

### linked list - LinkedList
```java
List<Integer> someList = new LinkedList();
```

### Queue
interface `Queue` with multiple implementations, like: 

- `LinkedList` 
- `LinkedBlockingQueue`
- `ConcurrentLinkedQueue`
- `ConcurrentLinkedDeque`
- `DelayQueue`
- `LinkedTransferQueue`
- `PriorityQueue`
- `SynchronousQueue`

### Stack
interface `Dequeue` (**_extends `Queue`_**)

Implementation examples:

- `ArrayDeque`  
- `LinkedList` 
- `ArrayDeque`  
- `LinkedBlockingDeque`

<div id="map" />
 
## Associative arrays - maps
```java
Map<String, Integer> dict = new HashMap();
```
Other implementations:

- `TreeMap` (_implements `NavigableMap`_)
- `ConcurrentHashMap`
- `ConcurrentSkipListMap`
- `SynchronizedMap`
- `Hashtable` (_old, not used now_)

### Example of simple implementation of LinkedList in Java [here](/linkedlist.html#java)


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

