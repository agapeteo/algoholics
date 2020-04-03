+++
title = "Binary heap"
template = "level.html"
description = "Binary heap data structure"
[extra]
    keywords = "Binary heap, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Binary heap

## Contents

---

- [Binary heap description](#description)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Binary heap

- push - `O log(N)`
 
- pop (get and delete maximum/minimum) - `O log(N)`

- peek (get maximum/minimum) - `O(1)`
 



<div id="java"/>

## Java
interface:
```java
package tree;

import java.util.Collection;

public interface Heap<E extends Comparable<E>> {
    E peak();

    E pop();

    void push(E element);

    int size();

    Collection<E> elements();

}
```

implementation:
```java
package tree;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

public class MaxBinaryHeap<E extends Comparable<E>> implements Heap<E> {
    private final List<E> container;

    public MaxBinaryHeap() {
        container = new ArrayList<>();
    }

    @Override
    public E peak() {
        return container.size() > 0 ? container.get(container.size() - 1) : null;
    }

    @Override
    public E pop() {
        if (container.isEmpty()) return null;

        if (container.size() == 1) {
            return container.remove(0);
        }

        E result = container.get(0);
        E latest = container.get(container.size() - 1);
        container.remove(container.size() - 1);
        container.set(0, latest);

        siftDown(0);

        return result;
    }

    private void siftDown(int idx) {
        if (idx < 0 || idx >= container.size() / 2) return;

        int leftIdx = leftChildIdx(idx);
        int rightIdx = rightChildIdx(idx);

        int largerChildIdx;
        if (rightIdx < container.size() && container.get(rightIdx).compareTo(container.get(leftIdx)) > 0) {
            largerChildIdx = rightIdx;
        } else {
            largerChildIdx = leftIdx;
        }

        if (container.get(largerChildIdx).compareTo(container.get(idx)) > 0) {
            swap(idx, largerChildIdx);
            siftDown(largerChildIdx);
        }
    }

    @Override
    public void push(E element) {
        if (element == null) {
            throw new IllegalArgumentException("null values are not valid elements");
        }
        container.add(element);
        siftUp(container.size() - 1);
    }

    private void siftUp(int idx) {
        if (idx <= 0) return;

        int parentIdx = parentIdx(idx);
        if (container.get(idx).compareTo(container.get(parentIdx)) > 0) {
            swap(idx, parentIdx);
            siftUp(parentIdx);
        }
    }

    private void swap(int i, int j) {
        E tmp = container.get(i);
        container.set(i, container.get(j));
        container.set(j, tmp);
    }

    @Override
    public int size() {
        return container.size();
    }

    @Override
    public Collection<E> elements() {
        List<E> result = new ArrayList<>(container.size());
        result.addAll(container);
        return result;
    }

    private int parentIdx(int idx) {
        return (idx - 1) / 2;
    }

    private int leftChildIdx(int idx) {
        return idx * 2 + 1;
    }

    private int rightChildIdx(int idx) {
        return idx * 2 + 2;
    }
}
```


<div id="go"/>

## Go

```go
package tree

import "errors"

var NoSuchElementError = errors.New("no such element")

type MaxHeap struct {
	elm []int
}

func (H *MaxHeap) Peak() (int, error) {
	if len(H.elm) == 0 {
		return -1, NoSuchElementError
	}
	return H.elm[0], nil
}

func (H *MaxHeap) Pop() (int, error) {
	if len(H.elm) == 0 {
		return -1, NoSuchElementError
	}

	var result = H.elm[0]

	if len(H.elm) == 1 {
		H.elm = make([]int, 0)
		return result, nil
	}

	H.elm[0] = H.elm[len(H.elm)-1] // put last to first
	H.elm = H.elm[:len(H.elm)-1]   // trim last element in slice

	siftDown(0, H.elm)

	return result, nil
}

func siftDown(i int, e []int) {
	if i < 0 || i >= len(e)/2 {
		return // no need to sift down leaf elements
	}

	var l = leftChildIdx(i)
	var r = rightChildIdx(i)

	var largerChild = l
	if r < len(e) && e[r] > e[l] {
		largerChild = r
	}

	if e[largerChild] > e[i] {
		swap(i, largerChild, e)
		siftDown(largerChild, e)
	}
}

func (H *MaxHeap) Push(v int) {
	if len(H.elm) == 0 {
		H.elm = append(H.elm, v)
		return
	}

	H.elm = append(H.elm, v)

	siftUp(len(H.elm)-1, H.elm)
}

func siftUp(i int, e []int) {
	if i <= 0 {
		return
	}
	var p = parentIdx(i)

	if e[i] > e[p] {
		swap(i, p, e)
		siftUp(p, e)
	}
}

func swap(i int, j int, e []int) {
	e[i], e[j] = e[j], e[i]
}

func (H *MaxHeap) Elements() []int {
	var result = make([]int, len(H.elm))
	copy(result, H.elm)
	return result
}

func (H *MaxHeap) Size() int {
	return len(H.elm)
}

func leftChildIdx(i int) int {
	return i*2 + 1
}

func rightChildIdx(i int) int {
	return i*2 + 2
}

func parentIdx(i int) int {
	return (i - 1) / 2
}
```



<div id="python"/>

## Python

```python
class Heap:
    def __init__(self):
        self.container = []

    def peek(self):
        if len(self.container) == 0:
            return None
        return self.container[0]

    def pop(self):
        if len(self.container) == 0:
            return None

        if len(self.container) == 1:
            return self.container.pop(0)

        result = self.container[0]
        last = self.container.pop()
        self.container[0] = last
        self._sift_down(0)
        return result

    def push(self, element):
        self.container.append(element)
        self._sift_up(len(self.container) - 1)

    def __len__(self):
        return len(self.container)

    def _sift_up(self, idx):
        if idx == 0:
            return
        parent_idx = self.parent_idx(idx)
        if self.container[idx] > self.container[parent_idx]:
            self._swap(idx, parent_idx)
            self._sift_up(parent_idx)

    def _sift_down(self, idx):
        if idx >= len(self.container) // 2:
            return

        right_child_idx = self.right_child_idx(idx)
        left_child_idx = self.left_child_idx(idx)
        greater_child_idx = left_child_idx

        if right_child_idx < len(self.container) and self.container[right_child_idx] > self.container[left_child_idx]:
            greater_child_idx = right_child_idx

        if self.container[greater_child_idx] > self.container[idx]:
            self._swap(idx, greater_child_idx)
            self._sift_down(greater_child_idx)

    def _swap(self, i, j):
        self.container[i], self.container[j] = self.container[j], self.container[i]

    @staticmethod
    def parent_idx(idx):
        return (idx - 1) // 2

    @staticmethod
    def left_child_idx(idx):
        return idx * 2 + 1

    @staticmethod
    def right_child_idx(idx):
        return idx * 2 + 2
```



<div id="javascript"/>

## JavaScript
```javascript
class MaxBinaryHeap {
    constructor() {
        this.container = [];
    }

    size() {
        return this.container.length;
    }

    peek() {
        if (this.container.length == 0) {
            return undefined;
        }
        return this.container[0];
    }

    push(element) {
        this.container.push(element);
        this.siftUp(this.container.length - 1);
    }

    pop() {
        if (this.container.length == 0) {
            return undefined;
        }
        if (this.container.length == 1) {
            return this.container.pop();
        }
        const result = this.container[0];
        const last = this.container.pop();
        this.container[0] = last;
        this.siftDown(0);

        return result;
    }

    siftUp(idx) {
        if (idx == 0) return;

        const parentIdx = MaxBinaryHeap.parentIdx(idx);
        if (this.container[idx] > this.container[parentIdx]) {
            this.swap(idx, parentIdx);
            this.siftUp(parentIdx);
        }
    }

    siftDown(idx) {
        if (idx >= Math.floor(this.container.length / 2)) {
            return
        }
        const leftChildIdx = MaxBinaryHeap.leftChildIdx(idx);
        const rightChildIdx = MaxBinaryHeap.rightChildIdx(idx);
        let greaterChildIdx = leftChildIdx;
        if (rightChildIdx < this.container.length && this.container[rightChildIdx] > this.container[leftChildIdx]) {
            greaterChildIdx = rightChildIdx;
        }
        if (this.container[greaterChildIdx] > this.container[idx]) {
            this.swap(greaterChildIdx, idx);
            this.siftDown(greaterChildIdx);
        }
    }

    swap(i, j) {
        [this.container[i], this.container[j]] = [this.container[j], this.container[i]]
    }

    static parentIdx(idx) {
        return Math.floor((idx - 1) / 2)
    }

    static rightChildIdx(idx) {
        return idx * 2 + 2
    }

    static leftChildIdx(idx) {
        return idx * 2 + 1
    }
}
```



<div id="kotlin"/>

## Kotlin

interface:
```kotlin
package tree

interface Heap<E : Comparable<E>> {

    fun peek(): E
    fun pop(): E
    fun push(value: E)

    fun elements(): List<E>
    fun size(): Int
}
```

implementation:
```kotlin
 package tree
 
 class BinaryHeap<E : Comparable<E>> : Heap<E> {
 
     private val elements = mutableListOf<E>()
 
     override fun size(): Int {
         return elements.size
     }
 
     override fun peek(): E {
         assert (elements.isNotEmpty() ) { "heap is empty" }
         return elements[0]
     }
 
     override fun pop(): E {
         assert (elements.isNotEmpty() ) { "heap is empty" }
 
         if (elements.size == 1){
             return elements.removeAt(0)
         }
 
         val result = elements[0]
         val latest = elements.removeAt(elements.size - 1)
         elements[0] = latest
 
         siftDown(0)
 
         return result
     }
 
     private fun siftDown(idx: Int) {
         if (idx < 0 || idx >= elements.size/2) return
 
         val leftIdx = leftChildIdx(idx)
         val rightIdx = rightChildIdx(idx)
 
         val largerChildIdx = if (rightIdx < elements.size && elements[rightIdx] > elements[leftIdx]) {
             rightIdx
         } else {
             leftIdx
         }
 
         if (elements[largerChildIdx] > elements[idx]){
             swap(idx, largerChildIdx)
             siftDown(largerChildIdx)
         }
     }
 
     override fun push(value: E) {
         elements.add(value)
 
         siftUp(elements.size - 1)
     }
 
     private fun siftUp(idx: Int) {
         if (idx <= 0) return
 
         val parentIdx = parentIdx(idx)
 
         if (elements[idx] > elements[parentIdx]){
             swap(idx, parentIdx)
             siftUp(parentIdx)
         }
     }
 
     override fun elements(): List<E> {
         val result = mutableListOf<E>()
 
         result.addAll(elements)
 
         return result
     }
 
     private fun swap(idxI: Int, idxJ: Int){
         val tmp: E = elements[idxI]
         elements[idxI] = elements[idxJ]
         elements[idxJ] = tmp
     }
 
     private fun leftChildIdx(idx: Int): Int = idx * 2 + 1
     private fun rightChildIdx(idx: Int): Int = idx * 2 + 2
     private fun parentIdx(idx: Int): Int = (idx - 1) / 2
 }
```

<div id="rust"/>

## Rust
```rust
use std::collections::VecDeque;

pub struct BinaryHeap<'a, T: Ord> {
    vec: VecDeque<&'a T>
}

impl<'a, T: Ord> BinaryHeap<'a, T> {
    pub fn new() -> Self {
        BinaryHeap { vec: VecDeque::new() }
    }

    pub fn peek_min(&self) -> Option<&T> {
        self.vec.get(0).copied()
    }

    pub fn push(&mut self, element: &'a T) {
        self.vec.push_back(element);
        self.sift_up(self.vec.len() - 1);
    }

    fn sift_up(&mut self, idx: usize) {
        if idx == 0 {
            return;
        }
        let parent_idx = parent_idx(idx);
        if self.vec[parent_idx] > self.vec[idx] {
            self.vec.swap(parent_idx, idx);
            self.sift_up(parent_idx);
        }
    }

    pub fn pop_min(&mut self) -> Option<&T> {
        let result = self.vec.pop_front();

        if self.vec.len() > 1 {
            let last = self.vec.pop_back().unwrap();
            self.vec.push_front(last);
            self.sift_down(0);
        }
        result
    }

    fn sift_down(&mut self, idx: usize) {
        if idx >= self.vec.len() / 2 {
            return;
        }

        let left_child_idx = left_child_idx(idx);
        let right_child_idx = right_child_idx(idx);

        let smaller_child_idx =
            if right_child_idx < self.vec.len() && self.vec[right_child_idx] < self.vec[left_child_idx] {
                right_child_idx
            } else {
                left_child_idx
            };

        if self.vec[smaller_child_idx] < self.vec[idx] {
            self.vec.swap(idx, smaller_child_idx);
            self.sift_down(smaller_child_idx);
        }
    }

    pub fn len(&self) -> usize {
        self.vec.len()
    }

    pub fn is_empty(&self) -> bool {
        self.vec.is_empty()
    }
}

fn parent_idx(idx: usize) -> usize {
    (idx - 1) / 2
}

fn left_child_idx(idx: usize) -> usize {
    idx * 2 + 1
}

fn right_child_idx(idx: usize) -> usize {
    idx * 2 + 2
}

#[test]
fn test() {
    let mut heap: BinaryHeap<i32> = BinaryHeap::new();

    assert_eq!(heap.is_empty(), true);
    assert_eq!(heap.peek_min(), None);

    heap.push(&-1);
    assert_eq!(heap.len(), 1);

    assert_eq!(heap.peek_min(), Some(&-1));
    assert_eq!(heap.pop_min(), Some(&-1));
    assert_eq!(heap.is_empty(), true);

    heap.push(&-99);
    heap.push(&12);
    heap.push(&5);
    heap.push(&1);
    heap.push(&4);
    heap.push(&0);
    heap.push(&-9);

    assert_eq!(heap.pop_min(), Some(&-99));
    assert_eq!(heap.pop_min(), Some(&-9));

    heap.push(&10);

    let mut actual = vec![];

    while let Some(e) = heap.pop_min() {
        actual.push(e.clone());
    }

    assert_eq!(actual, [0, 1, 4, 5, 10, 12]);
}
```