# Linked list

## Contents

---
- [Linked list description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)

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


