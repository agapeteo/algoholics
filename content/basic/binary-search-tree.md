+++
title = "Binary search tree"
template = "level.html"
description = "Binary search tree"
[extra]
    keywords = "binary search tree, tree, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, basic level"
+++

# Binary search tree

## Contents

---

- [Binary search description](#description)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Binary search tree

Most operations `O log(N)` 



<div id="java"/>

## Java


```java
package tree;

import java.util.ArrayList;
import java.util.List;

public class UnbalancedBinarySearchTree<E extends Comparable<E>> {
    static class Node<E> {
        E value;
        Node<E> left;
        Node<E> right;

        Node(E value) {
            this.value = value;
        }
    }

    private Node<E> root;

    public void add(E value) {
        if (root == null) {
            root = new Node<>(value);
            return;
        }
        addToRootNode(root, value);
    }

    private void addToRootNode(Node<E> rootNode, E value) {
        if (value.compareTo(rootNode.value) <= 0) {
            if (rootNode.left == null) {
                rootNode.left = new Node<>(value);
            } else {
                addToRootNode(rootNode.left, value);
            }
        } else {
            if (rootNode.right == null) {
                rootNode.right = new Node<>(value);
            } else {
                addToRootNode(rootNode.right, value);
            }
        }
    }

    public boolean contains(E value) {
        return rootNodeContains(root, value);
    }

    private boolean rootNodeContains(Node<E> rootNode, E value) {
        if (rootNode == null) return false;

        if (value == rootNode.value) return true;

        if (value.compareTo(rootNode.value) < 0) {
            return rootNodeContains(rootNode.left, value);
        } else {
            return rootNodeContains(rootNode.right, value);
        }
    }

    public List<E> elementsSorted() {
        List<E> result = new ArrayList<>();
        appendSortedForRootNode(root, result);
        return result;
    }

    private void appendSortedForRootNode(Node<E> rootNode, List<E> elements) {
        if (rootNode == null) return;

        appendSortedForRootNode(rootNode.left, elements);
        elements.add(rootNode.value);
        appendSortedForRootNode(rootNode.right, elements);
    }

    public int depth() {
        return depthFromNode(root, 0);
    }

    private int depthFromNode(Node<E> node, int curDepth) {
        if (node == null || (node.left == null && node.right == null)) {
            return curDepth;
        }

        return Math.max(
                depthFromNode(node.left, curDepth + 1),
                depthFromNode(node.right, curDepth + 1)
        );
    }
}

```


<div id="go"/>

## Go

```go
package tree

type node struct {
	value int
	left *node
	right *node
}

func newNode(value int) *node{
	return &node{value, nil, nil}
}

type UnbalancedBinarySearchTree struct {
	root *node
}

func (t *UnbalancedBinarySearchTree) Add(newValue int) {
	if t.root == nil{
		t.root = newNode(newValue)
		return
	}

	addToNode(t.root, newValue)
}

func addToNode(n *node, newValue int) {
	if newValue <= n.value {
		if n.left == nil{
			n.left = newNode(newValue)
		} else {
			addToNode(n.left, newValue)
		}
	} else {
		if n.right == nil {
			n.right = newNode(newValue)
		} else {
			addToNode(n.right, newValue)
		}
	}
}

func (t *UnbalancedBinarySearchTree) OrderedElements() []int{
	if t.root == nil {
		return make([]int, 0)
	}

	var result = addToSlice(t.root, make([]int, 0))

	return result
}

func addToSlice(n *node, values []int) []int {
	if n == nil {
		return values
	}

	values = addToSlice(n.left, values)
	values = append(values, n.value)
	values = addToSlice(n.right, values)

	return values
}

func (t UnbalancedBinarySearchTree) Contains(value int) bool{
	if t.root == nil {
		return false
	}

	return containsInNode(t.root, value)
}

func containsInNode(n *node, value int) bool {
	if n.value == value{
		return true
	}

	if value < n.value{
		if n.left == nil{
			return false
		} else {
			return containsInNode(n.left, value)
		}
	} else {
		if n.right == nil{
			return false
		} else {
			return containsInNode(n.right, value)
		}
	}
}

func (t *UnbalancedBinarySearchTree) Depth() int{
	if t.root == nil{
		return 0
	}

	return nodeDepth(t.root, 0)
}

func nodeDepth(n *node, curDepth int) int {
	var leftDepth = curDepth
	var rightDepth = curDepth

	if n.left != nil {
		leftDepth = nodeDepth(n.left, curDepth + 1)
	}

	if n.right != nil {
		rightDepth = nodeDepth(n.right, curDepth + 1)
	}

	if leftDepth > rightDepth{
		return leftDepth
	} else {
		return rightDepth
	}
}

```



<div id="python"/>

## Python

```python
class Node:
    def __init__(self, input_value):
        self.value = input_value
        self.left = None
        self.right = None


class Tree:
    def __init__(self):
        self.root = None
        pass

    def add(self, input_value):
        if not self.root:
            self.root = Node(input_value)
            return
        self.add_to_node(self.root, input_value)

    def add_to_node(self, root_node, input_value):
        if input_value <= root_node.value:
            if not root_node.left:
                root_node.left = Node(input_value)
            else:
                self.add_to_node(root_node.left, input_value)
        else:
            if not root_node.right:
                root_node.right = Node(input_value)
            else:
                self.add_to_node(root_node.right, input_value)

    def __contains__(self, input_value):
        return self.contains_from_node(self.root, input_value)

    def contains_from_node(self, root_node, input_value):
        if not root_node:
            return False
        if root_node.value == input_value:
            return True
        if input_value < root_node.value:
            return self.contains_from_node(root_node.left, input_value)
        else:
            return self.contains_from_node(root_node.right, input_value)

    def elements_sorted(self):
        result_list = []
        self.elements_sorted_from_node(self.root, result_list)
        return result_list

    def elements_sorted_from_node(self, root_node, input_list):
        if not root_node:
            return
        self.elements_sorted_from_node(root_node.left, input_list)
        input_list.append(root_node.value)
        self.elements_sorted_from_node(root_node.right, input_list)

```



<div id="javascript"/>

## JavaScript
```javascript
function Node(inputValue) {
  this.value = inputValue
  this.left
  this.right
}

class Tree {
  constructor(){
    this.root;
  }

  add(value) {
    if (!this.root) {
      this.root = new Node(value)
      return
      
    }
    this.addToNode(this.root, value)
  }

  addToNode(node, value){
    if (value <= node.value) {
      if (!node.left) {
        node.left = new Node(value)
      } else {
        this.addToNode(node.left, value)
      }
    } else {
      if (!node.right) {
        node.right = new Node(value )
      } else {
        this.addToNode(node.right, value)
      }
    }   
  }

  contains(value) {
    return this.containsInNode(this.root, value)
  }

  containsInNode(node, value) {
    if (!node) return false
   
    if (node.value == value) return true

    if (value < node.value) {
      return this.containsInNode(node.left, value)
    } else {
      return this.containsInNode(node.right, value)
    }   
  }

  elements(){    
    result = []
    this.elementsForNode(this.root, result)
    return result
  }

  elementsForNode(node, arr) {
    if (node == null) {
      return
    }
    this.elementsForNode(node.left, arr)
    arr.push(node.value)
    this.elementsForNode(node.right, arr)
  }

  depth() {
    return this.depthForNode(this.root, 0)
  }
  
  depthForNode(node, curDepth) {
    if (!node || (!node.left && !node.right)) return curDepth
    return Math.max(
      this.depthForNode(node.left, curDepth + 1),
      this.depthForNode(node.right, curDepth + 1),
    )
  }
  
}
```



<div id="kotlin"/>

## Kotlin
interface:
```kotlin
package tree

interface BinarySearchTree<E : Comparable<E>> {
    fun add(value: E)
    fun elements(): List<E>
    fun contains(value: E): Boolean
    fun orderedElements(): List<E>
    fun depth(): Int
}
```

implementation:
```kotlin
 package tree
 
 import kotlin.math.max
 
 class UnbalancedBinarySearchTree<E : Comparable<E>> : BinarySearchTree<E> {
 
     private var root: Node<E>? = null
 
     private data class Node<E>(val value: E, var left: Node<E>? = null, var right: Node<E>? = null)
 
     override fun add(value: E) {
         if (root == null) {
             root = Node(value)
             return
         }
 
         addNode(root!!, value)
     }
 
     private fun addNode(rootNode: Node<E>, value: E) {
         if (value <= rootNode.value) {
             if (rootNode.left == null) {
                 rootNode.left = Node(value)
             } else {
                 addNode(rootNode.left!!, value)
             }
         } else {
             if (rootNode.right == null) {
                 rootNode.right = Node(value)
             } else {
                 addNode(rootNode.right!!, value)
             }
         }
     }
 
     override fun elements(): List<E> {
         if (root == null) return listOf()
 
         val result = mutableListOf<E>()
 
         addRootElements(root, result)
 
         return result
     }
 
     private fun addRootElements(root: Node<E>?, list: MutableList<E>) {
         if (root == null) return
 
         list.add(root.value)
         addRootElements(root.left, list)
         addRootElements(root.right, list)
     }
 
     override fun contains(value: E): Boolean {
         if (root == null) return false
 
         return containsFromRoot(root, value)
     }
 
     private fun containsFromRoot(root: Node<E>?, value: E): Boolean {
         if (root == null) return false
 
         if (root.value == value) return true
 
         return if (value < root.value) {
             containsFromRoot(root.left, value)
         } else {
             containsFromRoot(root.right, value)
         }
     }
 
     override fun orderedElements(): List<E> {
         if (root == null) return listOf()
 
         val result = mutableListOf<E>()
         addElement(root, result)
         return result
     }
 
     private fun addElement(node: Node<E>?, list: MutableList<E>) {
         if (node == null) return
 
         addElement(node.left, list)
         list.add(node.value)
         addElement(node.right, list)
     }
 
     override fun depth(): Int {
         if (root == null) return 0
 
         return depthNode(root!!, 0)
     }
 
     private fun depthNode(node: Node<E>, curDepth: Int): Int {
         return max(
                 if (node.left == null) {
                     curDepth
                 } else {
                     depthNode(node.left!!, curDepth + 1)
                 }, if (node.right == null) {
             curDepth
         } else {
             depthNode(node.right!!, curDepth + 1)
         })
     }     
 }

```


<div id="rust"/>

## Rust

```rust
pub struct BinaryTree<T: PartialOrd + Clone> {
    root: Option<Box<Node<T>>>
}

impl<T: PartialOrd + Clone> BinaryTree<T> {
    pub fn new() -> Self {
        BinaryTree { root: None }
    }

    pub fn add(&mut self, value: T) {
        let new_node = Node::new_boxed(value);
        match self.root {
            None => {
                self.root.replace(new_node);
            }
            Some(ref mut root) => {
                root.add(new_node);
            }
        }
    }

    pub fn contains(&self, item: T) -> bool {
        match self.root {
            None => { false }
            Some(ref node) => {
                node.contains(item)
            }
        }
    }

    pub fn depth(&self) -> usize {
        let mut size = 0_usize;

        if let Some(root_node) = &self.root {
            size = root_node.depth(size);
        }
        size
    }

    pub fn collect_values(&self) -> Vec<T> {
        let mut vec: Vec<T> = Vec::new();

        if let Some(root_node) = &self.root {
            root_node.collect_values(&mut vec);
        }
        vec
    }
}

impl<T: PartialOrd + Clone> Default for BinaryTree<T> {
    fn default() -> Self {
        BinaryTree::new()
    }
}

struct Node<T: PartialOrd + Clone> {
    value: T,
    left: Option<Box<Node<T>>>,
    right: Option<Box<Node<T>>>,
}

impl<T: PartialOrd + Clone> Node<T> {
    fn new_boxed(value: T) -> Box<Node<T>> {
        Box::new(Node { value, left: None, right: None })
    }

    fn add(&mut self, other_node: Box<Node<T>>) {
        if other_node.value > self.value {
            match self.right {
                None => {
                    self.right.replace(other_node);
                }
                Some(ref mut right_node) => {
                    right_node.add(other_node);
                }
            }
        } else { // add to left branch
            match self.left {
                None => {
                    self.left.replace(other_node);
                }
                Some(ref mut left_node) => {
                    left_node.add(other_node);
                }
            }
        }
    }

    fn contains(&self, item: T) -> bool {
        if self.value == item {
            return true;
        }
        if item > self.value {
            if let Some(right_node) = &self.right {
                return right_node.contains(item);
            }
        } else { // search in left branch
            if let Some(left_node) = &self.left {
                return left_node.contains(item);
            }
        }
        false
    }

    fn depth(&self, cur_size: usize) -> usize {
        let mut left_depth = 0_usize;
        if let Some(left_node) = &self.left {
            left_depth = left_node.depth(cur_size + 1);
        }

        let mut right_depth = 0_usize;
        if let Some(right_node) = &self.right {
            right_depth = right_node.depth(cur_size + 1);
        }
        if left_depth == 0 && right_depth == 0 {
            cur_size
        } else {
            std::cmp::max(left_depth, right_depth)
        }
    }

    fn collect_values(&self, vec: &mut Vec<T>) {
        if let Some(left_node) = &self.left {
            left_node.collect_values(vec);
        }
        let value = &self.value;
        vec.push(value.to_owned());

        if let Some(right_node) = &self.right {
            right_node.collect_values(vec);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn collect_values() {
        let mut tree = BinaryTree::new();
        tree.add(6);
        tree.add(3);
        tree.add(8);
        tree.add(1);
        tree.add(2);
        tree.add(7);
        tree.add(10);
        tree.add(4);
        tree.add(5);
        tree.add(9);

        let values = tree.collect_values();
        assert_eq!(values, (1..=10).collect::<Vec<i32>>());
    }

    #[test]
    fn contains() {
        let mut tree = BinaryTree::new();
        tree.add(6);
        tree.add(3);
        tree.add(8);

        assert_eq!(tree.contains(6), true);
        assert_eq!(tree.contains(3), true);
        assert_eq!(tree.contains(8), true);

        assert_eq!(tree.contains(5), false);
        assert_eq!(tree.contains(1), false);
        assert_eq!(tree.contains(7), false);

        tree.add(5);
        tree.add(1);
        tree.add(7);

        assert_eq!(tree.contains(5), true);
        assert_eq!(tree.contains(1), true);
        assert_eq!(tree.contains(7), true);
    }

    #[test]
    fn depth() {
        let mut tree = BinaryTree::new();
        assert_eq!(tree.depth(), 0);

        tree.add(6);
        assert_eq!(tree.depth(), 0);

        tree.add(3);
        assert_eq!(tree.depth(), 1);

        tree.add(8);
        assert_eq!(tree.depth(), 1);

        tree.add(2);
        assert_eq!(tree.depth(), 2);

        tree.add(10);
        assert_eq!(tree.depth(), 2);

        tree.add(11);
        assert_eq!(tree.depth(), 3);
    }

    #[test]
    fn depth_right_unbalanced() {
        let mut tree = BinaryTree::new();
        const N: usize = 100;

        for i in 0..=N {
            tree.add(i);
        }
        assert_eq!(tree.depth(), N);
    }

    #[test]
    fn depth_left_unbalanced() {
        let mut tree = BinaryTree::new();
        const N: usize = 100;

        for i in (0..=N).rev() {
            tree.add(i);
        }
        assert_eq!(tree.depth(), N);
    }
}
```