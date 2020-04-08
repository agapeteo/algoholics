+++
title = "Trie"
template = "level.html"
description = "Trie data structure"
[extra]
    keywords = "Trie, prefix, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Trie

## Contents

---

- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---
 

<div id="java"/>

## Java

interface:
```java
package trie;

public interface Trie {

    void add(String word);

    boolean contains(String word);
}

```
implementation:
```java
package trie;

import java.util.HashMap;
import java.util.Map;

public class BasicTrie implements Trie {
    private static class Node {
        Map<Character, Node> nodes = new HashMap();
        boolean isLast;
    }

    private Node root = new Node();

    @Override
    public void add(String word) {
        if (word.isEmpty()) return;

        add(root, word.toCharArray(), 0);
    }

    private void add(Node node, char[] chars, int idx) {
        if (idx >= chars.length) {
            return;
        }
        char curChar = chars[idx];
        Node nextNode = node.nodes.get(curChar);
        if (nextNode == null) {
            nextNode = new Node();
            node.nodes.put(curChar, nextNode);
        }
        if (idx == chars.length - 1) {
            nextNode.isLast = true;
        }
        add(nextNode, chars, idx + 1);
    }

    @Override
    public boolean contains(String word) {
        if (word.isEmpty()) return false;

        return contains(root, word.toCharArray(), 0);
    }

    private boolean contains(Node node, char[] chars, int idx) {
        if (idx >= chars.length) return false;

        char curChar = chars[idx];
        Node nextNode = node.nodes.get(curChar);

        if (nextNode == null) {
            return false;
        } else {
            if (idx == chars.length - 1 && nextNode.isLast) {
                return true;
            }
        }
        return contains(nextNode, chars, idx + 1);
    }
}
```

<div id="go"/>

## Go

```go
package trie

type node struct {
	nodes map[rune]*node
	last  bool
}

func newNode() *node {
	return &node{nodes: make(map[rune]*node), last: false}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

func (t *Trie) Add(str string) {
	if len(str) == 0 {
		return
	}
	runes := []rune(str)
	addToNode(t.root, runes, 0)
}

func addToNode(n *node, runes []rune, idx int) {
	if idx >= len(runes) {
		return
	}
	r := runes[idx]

	var nextNode *node
	var ok bool
	if nextNode, ok = n.nodes[r]; !ok {
		nextNode = newNode()
		n.nodes[r] = nextNode
	}
	if idx == len(runes)-1 {
		nextNode.last = true
	}
	addToNode(nextNode, runes, idx+1)
}

func (t *Trie) Contains(str string) bool {
	if len(str) == 0 {
		return false
	}
	runes := []rune(str)
	return containsFromNode(t.root, runes, 0)
}

func containsFromNode(n *node, runes []rune, idx int) bool {
	if idx >= len(runes) {
		return false
	}
	r := runes[idx]

	var nextNode *node
	var ok bool
	if nextNode, ok = n.nodes[r]; !ok {
		return false
	}
	if idx == len(runes)-1 && nextNode.last {
		return true
	}
	return containsFromNode(nextNode, runes, idx+1)
}
```


<div id="python"/>

## Python
```python
class Node:
    def __init__(self):
        self.nodes = {}
        self.is_last = False


class Trie:
    def __init__(self):
        self.root = Node()

    def add(self, text):
        cur_node = self.root
        idx = 0
        for char in text:
            try:
                next_node = cur_node.nodes[char]
            except KeyError:
                next_node = Node()
                cur_node.nodes[char] = next_node
            if idx == len(text) - 1:
                next_node.is_last = True
            cur_node = next_node
            idx += 1

    def __add__(self, other):
        self.add(other)

    def __contains__(self, text):
        cur_node = self.root
        idx = 0
        for char in text:
            try:
                next_node = cur_node.nodes[char]
            except KeyError:
                return False
            if idx == len(text) - 1 and next_node.is_last:
                return True
            cur_node = next_node
            idx += 1
        return False

```



<div id="javascript"/>

## JavaScript

```javascript
class Node {
    constructor() {
        this.nodes = new Map();
        this.isLast = false;
    }
}

class Trie {
    constructor() {
        this.root = new Node();
    }

    add(text) {
        let curNode = this.root;
        for (let i = 0; i < text.length; i++) {
            const char = text[i];
            let nextNode = curNode.nodes.get(char);
            if (!nextNode) {
                nextNode = new Node();
                curNode.nodes.set(char, nextNode);
            }
            if (i == text.length - 1) {
                nextNode.isLast = true;
            }
            curNode = nextNode;
        }
    }

    contains(text) {
        let curNode = this.root;
        for (let i = 0; i < text.length; i++) {
            const char = text[i];
            let nextNode = curNode.nodes.get(char);
            if (!nextNode) {
                return false;
            }
            if (i == text.length - 1 && nextNode.isLast) {
                return true;
            }
            curNode = nextNode;
        }
        return false;
    }
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package trie

class SimpleTrie : Trie {

    private data class Node(val childrenMap: HashMap<Char, Node> = HashMap(), var isLast: Boolean = false)

    private val root = Node()

    override fun add(str: String) {

        fun add(charIdx: Int, node: Node) {
            if (charIdx >= str.length) return

            val curChar = str[charIdx]
            var curNode = node.childrenMap[curChar]

            if (curNode == null) {
                curNode = Node()
                node.childrenMap[curChar] = curNode
            }
            if (charIdx == str.length - 1) {
                curNode.isLast = true
            }
            add(charIdx + 1, curNode)
        }
        add(0, root)
    }

    override fun contains(str: String): Boolean {
        fun contains(charIdx: Int, node: Node): Boolean {
            if (charIdx >= str.length) return false

            val curChar = str[charIdx]
            val curNode = node.childrenMap[curChar] ?: return false

            if (charIdx == str.length - 1 && curNode.isLast) return true

            return contains(charIdx + 1, curNode)
        }
        return contains(0, root)
    }
}
```


<div id="rust"/>

## Rust

```rust
use std::collections::HashMap;
use std::cell::RefCell;
use std::rc::Rc;
use std::collections::hash_map::Entry;

struct Node {
    nodes: HashMap<char, Rc<RefCell<Node>>>,
    is_final: bool,
}

impl Default for Node {
    fn default() -> Self {
        Node { nodes: HashMap::new(), is_final: false }
    }
}

impl Node {
    fn new() -> Self {
        Node::default()
    }
}

pub struct Trie {
    root: Rc<RefCell<Node>>
}

impl Default for Trie {
    fn default() -> Self {
        Trie { root: Rc::new(RefCell::new(Node::new())) }
    }
}

fn to_chars(str: &str) -> Vec<char> {
    let mut chars: Vec<char> = Vec::new();
    for char in str.chars() {
        chars.push(char);
    }
    chars
}

impl Trie {
    pub fn new() -> Self {
        Trie::default()
    }

    #[allow(dead_code)]
    pub fn add(&mut self, str: &str) {
        Trie::add_to_node(&to_chars(str), self.root.clone(), 0)
    }

    fn add_to_node(chars: &[char], node: Rc<RefCell<Node>>, idx: usize) {
        if idx >= chars.len() {
            return;
        }
        let char = chars[idx];

        let mut ref_mut_node = node.borrow_mut();
        let entry = ref_mut_node.nodes.entry(char);

        let next_node = match entry {
            Entry::Occupied(occupied_entry) => {
                occupied_entry.get().clone()
            }
            Entry::Vacant(vacant_entry) => {
                let new_node = Rc::new(RefCell::new(Node::new()));
                vacant_entry.insert(new_node.clone());
                new_node
            }
        };

        if idx == chars.len() - 1 {
            next_node.borrow_mut().is_final = true;
        }

        Trie::add_to_node(chars, next_node, idx + 1)
    }

    pub fn contains(&self, str: &str) -> bool {
        Trie::contains_from_node(&to_chars(str), self.root.clone(), 0)
    }

    fn contains_from_node(chars: &[char], node: Rc<RefCell<Node>>, idx: usize) -> bool {
        if idx >= chars.len() {
            return false;
        }

        let char = chars[idx];

        let ref_node = node.borrow();
        let next_node = ref_node.nodes.get(&char);

        if next_node.is_none() {
            return false;
        }
        let next_node = next_node.unwrap().clone();
        if idx == chars.len() - 1 && next_node.borrow().is_final {
            return true;
        }

        Trie::contains_from_node(chars, next_node, idx + 1)
    }
}

#[test]
fn test() {
    let mut trie = Trie::new();

    trie.add("cat");
    trie.add("can");
    trie.add("cast");
    trie.add("boost");

    assert_eq!(trie.contains("cat"), true);
    assert_eq!(trie.contains("can"), true);
    assert_eq!(trie.contains("cas"), false);
    assert_eq!(trie.contains("casting"), false);
    assert_eq!(trie.contains("cost"), false);
    assert_eq!(trie.contains("boost"), true);
    assert_eq!(trie.contains("but"), false);
    assert_eq!(trie.contains("bot"), false);
}
```

