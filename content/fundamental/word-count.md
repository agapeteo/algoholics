+++
title = "Word count"
template = "level.html"
description = "Word count using map"
[extra]
    keywords = "word count, map, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Word count

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
```java
public static Map<String, Integer> workCount(String text) {
    Map<String, Integer> result = new HashMap<>();

    String[] words = text.split(" ");

    for (String word : words) {
        Integer count = result.get(word);
        if (count == null) {
            result.put(word, 1);
        } else {
            result.put(word, count + 1);
        }
    }
    return result;
}
```

using lambda:
```java
public static Map<String, Long> lambdaWordCount(String text) {
    return Arrays.asList(text.split(" "))
            .stream()
            .collect(
                    groupingBy(
                            Function.identity(), counting()));
}
```


<div id="go"/>

## Go

```go
func WordCount(text string) map[string]int {
	words := make(map[string]int)
	for _, w := range strings.Split(text, " ") {
		_, ok := words[w]
		if ok {
			words[w] += 1
		} else {
			words[w] = 1
		}
	}
	return words
}
```


<div id="python"/>

## Python

```python
def word_count(in_text):
    words = {}
    for word in in_text.split():
        if word in words:
            words[word] += 1
        else:
            words[word] = 1
    return words
```


<div id="javascript"/>

## JavaScript

```javascript
function wordCount(text) {
  const words = new Map()
  for (const word of text.split(" ")) {
    const count = words[word]
    if (count) {
      words[word] = count + 1
    } else {
      words[word] = 1
    }
  }
  return words
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
fun wordCount(text: String): Map<String, Int> {
    return text.split(" ").groupingBy { it }.eachCount()
}
```


<div id="rust"/>

## Rust 

```rust
fn word_count(text: &str) -> HashMap<&str, i32> {
    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        *map.entry(word).or_insert(0) += 1;
    }
    map
}
```
