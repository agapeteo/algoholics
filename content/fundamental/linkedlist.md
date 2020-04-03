+++
title = "Linked list"
template = "level.html"
description = "Linked list"
[extra]
    keywords = "Linked list, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Linked list

## Contents

---
- [Linked list description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description"    />

## Linked list description
Linked list is collection of not fixed number of items liked to each other by reference/pointer

``TODO: more details`` 


<div id="practice"/>

## In practice
For list of elements when most operation changing collection size **and** no searched by index required

Comparing to dynamic array linked list is used less often due to poor cache locality

``TODO: array list vs linked list``


<div id="java"/>

## Java

`Java` provides build-in implementation of linked list called `LinkedList` (surprise?)

Here is simplified version of LinkedList in `Java` which also implements `Stack` and `Queue`:
<script src="https://gist.github.com/agapeteo/40e8764a5da58bb54b9a8501feedd725.js"></script>


<div id="go"/>

## Go
If you need to use linked list or stack/queue based linked list in production prefer using [Go's container/list] (https://golang.org/pkg/container/list/)

<br/>
Another version of simple `Linked list` which also implements `Stack` and `Queue`:
<script src="https://gist.github.com/agapeteo/04a1542bf6c691c0464f8e002024484f.js"></script>



<div id="python" />

## Python
See Python's [collections.deque](https://docs.python.org/3/library/collections.html#collections.deque) which is implemented based on `Linked list` and has methods for `Stack` and `Queue`




<div id="javascript" />

## JavaScript
[Check explanation here](https://humanwhocodes.com/blog/2019/02/computer-science-in-javascript-doubly-linked-lists/)

[and code example here](https://github.com/humanwhocodes/computer-science-in-javascript/blob/master/src/data-structures/doubly-linked-list/doubly-linked-list.js)




<div id="kotlin" />

## Kotlin
Example of simplified version of `LinkedList` in Kotlin (also implements `Stack` and `Queue`)
<script src="https://gist.github.com/agapeteo/15d80043e9b1454a7d57e102ea849ef4.js"></script>


<div id="rust" />

## Rust
See example of build-in [LinkedList](/languages/rust#linked-list)

You can find Rust's optimized implementation of Linked List [here](https://github.com/rust-lang/rust/blob/master/src/liballoc/collections/linked_list.rs)

Simple implementation of Cons List as alternative version of linked List (not that useful in practice)
```rust
use std::fmt::{Display};

enum ConsList<T> {
    Empty,
    Elem(T, Box<ConsList<T>>),
}

impl<T: PartialEq + Display> ConsList<T> {
    fn new() -> Self {
        ConsList::Empty
    }

    fn print_elements(&self) {
        let mut cur_list = self;
        while let ConsList::Elem(elem, boxed_list) = cur_list {
            println!("elem: {}", elem);
            cur_list = boxed_list;
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test() {
        let list_1: ConsList<i32> = ConsList::Elem(0,
                                                   Box::new(ConsList::Elem(1,
                                                                           Box::new(ConsList::Elem(2,
                                                                                                   Box::new(ConsList::Empty))))));
        list_1.print_elements();
    }
}
```


See also about journey/tutorial on how to implement linked list in [Learn Rust With Entirely Too Many Linked Lists](https://rust-unofficial.github.io/too-many-lists/) 