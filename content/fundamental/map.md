+++
title = "Associative arrays and maps"
template = "level.html"
description = "Associative arrays and maps, hash maps"
[extra]
    keywords = "map, hashmap, hashtable, dictionary, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Associative array
# (Map, Dictionary)

## Contents

---
- [Dynamic array, Map description](#description)
- [In practice](#practice)
- [Build-in data structures for next languages](#build-in)

---


<div id="description" />

## Dynamic array, Map description
According to [Wikipedia](https://en.wikipedia.org/wiki/Associative_array):
associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears at most once in the collection

Excellent explanation by Brian Kernighan:
<iframe width="800" height="500" src="https://www.youtube.com/embed/qTZJLJ3Gm6Q" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

``TODO: more details``


<div id="practice"/>

## In practice
`Associative array (map)` is used very often in practice - usually for dictionaries or when quick lookup required

> Potential problem: `hash collision`
>


``TODO: more details``

update FROM Rust in action book:

"You might encounter someone with a computer science background preferring to use
the term hash table. Perl and Ruby strip that off and call them hashes. Lua does the
opposite and uses the term table. Many communities name the structure after a
metaphor, such as a dictionary (one term is being associated with a “definition”) or
a map (programmers, following mathematicians, are mapping from one value to
another). Other communities prefer naming based on the role that the structure plays.
PHP describes them as associative arrays. JavaScript’s objects tend to be implemented
as a key/value pair collection and so generic term object suffices. Static languages tend
to name them according to how they are implemented. C++ and Java distinguish
between a hash map and a tree map.

Rust uses the terms HashMap and BTreeMap to define two implementations of the same
abstract data type. Rust is closest to C++ and Java in this regard.
"



<div id="build-in" />

## Build-in data structures for languages:
 
- [Java](/languages/java#map)
- [Go](/languages//golang#map)
- [JavaScript](/languages/javascript#map) 
- [Python](/languages//python#map)  
- [Kotlin](/languages/kotlin#map)
- [Rust](/languages/rust#map)  


