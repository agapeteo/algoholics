+++
title = "Rust"
template = "level.html"
description = "data structures and algorithms in Rust programming languages"
[extra]
    keywords = "Rust, explain, tutorial, data structures, algorithms, interview"
+++

# Rust

## Contents
---
- [Array](#array)
- [Vector - dynamic array](#dynamic-array)
- [Slices](#slices)
- [Map](#map)
- [Set](#set)
- [VecDeque](#vecdeque)
- [LinkedList](#linked-list)
- [Queue](#queue)
- [Stack](#stack)
- [Not a humble opinion about language](#pros-cons)

---

<div id="array"/>

## Array
create array:
```rust
#[test]
fn array() {
    let array_of_ints: [usize; 3] = [0; 3]; // create immutable array of size 3 with default values 0

    assert_eq!(array_of_ints.len(), 3);
    assert_eq!(array_of_ints[..], [0, 0, 0]); // compare with slice


    let mut array_of_str = ["-", "-"]; // create mutable array with predefined values
    array_of_str[1] = "+";
    assert_eq!(array_of_str.len(), 2);
    assert_eq!(array_of_str[..], ["-", "+"]);

    for s in &array_of_str {
        println!("{}", s)
    }

    // iterate using iterator
    let mut arr_iter = array_of_str.iter();
    while let Some(s) = arr_iter.next() {
        println!("{}", s);
    }
}
```

see more details about arrays [here](https://doc.rust-lang.org/std/primitive.array.html) 


<div id="dynamic-array"/>

## Vec - dynamic array
create vector:
```rust
#[test]
fn vector() {
    let empty_vector: Vec<String> = vec![]; // create empty vector using vec! macro
    let non_empty_macro_vector = vec!["a", "b", "c"]; // create vector with predefined values, type is inferred

    let mut some_vec: Vec<i8> = Vec::new(); // note, empty vector doesn't allocate any memory in heap, only fat pointer in stack allocated
    some_vec.push(0);
    some_vec.push(1);
    some_vec.push(1);

    assert_eq!(some_vec.len(), 3);
    some_vec[2] = 2;
    some_vec.push(3);
    assert_eq!(some_vec, [0, 1, 2, 3]);
    assert_eq!(some_vec.len(), 4);

    while let Some(latest) = some_vec.pop() {
        println!("{}", latest)
    }
    assert_eq!(some_vec.is_empty(), true);
}
```
see more details about Vec [here](https://doc.rust-lang.org/std/vec/struct.Vec.html)


<div id="slices"/>

## Slices (view of arrays/vector/string etc)
Slices are very convenient way of search/compare/mutate etc arrays, vector and other types like string without additional heap or stack allocation (only fat pointer)

```rust
#[test]
fn slices() {
    let mut vec = vec!["a", "b", "c", "d"];

    let slice_all = &vec[..];
    assert_eq!(slice_all, ["a", "b", "c", "d"]);

    let slice_first_two = &vec[..2];
    assert_eq!(slice_first_two, ["a", "b"]);

    let slice_without_first = &vec[1..];
    assert_eq!(slice_without_first, ["b", "c", "d"]);

    let (first_half, second_half) = vec.split_at_mut(2);
    first_half[0] = "_";
    second_half[0] = "?";

    for s in &vec {
        println!("{}", s);
    }
    assert_eq!(vec, ["_", "b", "?", "d"]);
}
```
See more details about slices [here](https://doc.rust-lang.org/std/primitive.slice.html)


<div id="map" />
 
## Map
> Rust's HashMap uses quadratic probing and SIMD lookup and provides resistance against HashDoS attacks 
> Default hashing algorithm is SipHash 1-3, but you can specify your own using 'with_hasher' method

create HashMap:
```rust
use std::collections::HashMap;

#[test]
fn hashmap() {
    let mut map = HashMap::new();
    map.insert(1, "one");
    map.insert(2, "two");
    map.insert(3, "three");

    assert_eq!(map.len(), 3);
    assert_eq!(map.get(&1), Some(&"one"));
    assert_eq!(map.get(&3), Some(&"three"));
    assert_eq!(map.get(&4), None);

    for (k, v) in map {
        println!("key: {}, value: {}", k, v);
    }
}
```
see more details about HashMap [here](https://doc.rust-lang.org/std/collections/struct.HashMap.html) 


<div id="set" />

## Set
```rust
use std::collections::HashSet;

#[test]
fn hashset() {
    let mut set = HashSet::new();
    set.insert(1);
    set.insert(1);
    set.insert(2);
    set.insert(2);

    assert_eq!(set.len(), 2);
    assert_eq!(set.contains(&1), true);
    assert_eq!(set.contains(&2), true);
    set.remove(&2);
    assert_eq!(set.contains(&2), false);

    for n in set {
        println!("{}", n)
    }
}
```
see more details about Set [here](https://doc.rust-lang.org/std/collections/struct.HashSet.html)


<div id="vecdeque" />

## VecDeque  

> VecDeque is double ended queue using growable ring buffer
>
> VecDeque is preferred to LinkedList of queue in most cases
```rust
#[test]
fn vecdeque() {
    // new with capacity
    let _v: VecDeque<&str> = VecDeque::with_capacity(1_000);

    // from Vec
    let _v = VecDeque::from(vec![1, 2, 3]);

    // empty new
    let mut v: VecDeque<i32> = VecDeque::new();

    // add to end
    v.push_back(1);
    v.push_back(2);
    assert_eq!(v, [1, 2]);

    // pop last
    let last = v.pop_back();
    assert_eq!(Some(2), last);
    assert_eq!(v, [1]);


    // add to front beginning
    v.push_front(0);
    assert_eq!(v, [0, 1]);

    let first = v.pop_front();
    assert_eq!(Some(0), first);
    assert_eq!(v, [1]);

    // get by index and mutate
    if let Some(elem) = v.get_mut(0) {
        *elem *= 5;
    }

    // get by index
    assert_eq!(Some(&5), v.get(0));

    // iterate
    assert_eq!(Some(&5), v.iter().next());
}
```

see other methods like `drain`, `swap`, `splitoff`, `append`, `rotate_left` and other details about VecDeque [here](https://doc.rust-lang.org/std/collections/struct.VecDeque.html)


<div id="linked-list"/>

## Linked list
create linked list:
```rust
use std::collections::LinkedList;

#[test]
fn linked_list() {
    let mut linked_list: LinkedList<u8> = LinkedList::new();
    assert_eq!(linked_list.is_empty(), true);

    linked_list.push_back(1);
    linked_list.push_back(2);
    linked_list.push_front(0);

    assert_eq!(linked_list.pop_front(), Some(0));
    assert_eq!(linked_list.pop_back(), Some(2));
    assert_eq!(linked_list.pop_back(), Some(1));
    assert_eq!(linked_list.pop_back(), None);
    assert_eq!(linked_list.is_empty(), true);
}
```
see more details about linked list [here](https://doc.rust-lang.org/std/collections/struct.LinkedList.html) 


<div id="queue" />

## Queue
> This example uses VecDeque for queue implementation.
> Queue can also be implemented using `std::collections::LinkedList`

```rust
use std::collections::VecDeque;

#[test]
fn queue() {
    let mut queue: VecDeque<u8> = VecDeque::new();
    queue.push_back(1);
    queue.push_back(2);
    queue.push_front(0);

    assert_eq!(queue.len(), 3);

    assert_eq!(queue.get(1), Some(&1));
    assert_eq!(queue.pop_front(), Some(0));
    assert_eq!(queue.pop_back(), Some(2));
    assert_eq!(queue.pop_back(), Some(1));
    assert_eq!(queue.pop_back(), None);
}
```


<div id="stack" />

## Stack
> This example uses VecDeque for queue implementation.
> Stack can also be implemented using `std::collections::LinkedList`

```rust
use std::collections::VecDeque;

#[test]
fn stack() {
    let mut stack = VecDeque::new();
    stack.push_back(1);
    stack.push_back(2);
    stack.push_back(3);

    assert_eq!(stack.len(), 3);

    assert_eq!(stack.pop_back(), Some(3));
    assert_eq!(stack.pop_back(), Some(2));
    assert_eq!(stack.pop_back(), Some(1));
    assert_eq!(stack.pop_back(), None);
}
```


<div id="pros-cons" />

## Not a humble opinion about language
**pros**

- safe by default
- as fast as C or C++
- you have full control
- no GC 
- predictable performance 
- safe multithreading
- functional features
- types are immutable by default 
- nice community


**cons**

- steep learning curve
- ecosystem is good enough but not mature as in other popular languages

