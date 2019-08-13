# Binary search

## Contents

---

- [Binary search description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="description" />

## Binary search

Complexity `O log(N)` - Quick search in sorted collection of data 

There are two way to find mid value:
- low + high / 2
- low + (high - low) / 2
Later one helps to avoid memory overflow



<div id="practice"/>

## In practice

> `TODO:` more details


<div id="java"/>

## Java
> `java.util` already contains `Collections.binarySearch()` and `Arrays.binarySearch()`  binary search implementations
> this function expect that collection is sorted. If collection contain multiple target values - there is no guarantee which index of those will be returned

```java
public static <T> int binarySearch(List<? extends Comparable<? super T>> list, T value) {
        int lowIdx = 0;
        int highIdx = list.size() - 1;

        while (lowIdx <= highIdx) {
            int midIdx = lowIdx + (highIdx - lowIdx / 2);

            Comparable<? super T> midValue = list.get(midIdx);
            if (midValue == value) {
                return midIdx;
            }
            if (midValue.compareTo(value) < 1) {
                lowIdx = midIdx + 1;
            } else {
                highIdx = midIdx - 1;
            }
        }
        return -(lowIdx + 1); // not found, also shows where it stopped with minus sign
    }
}
```


<div id="go"/>

## Go

> `TODO: `



<div id="python"/>

## Python
> `TODO: `



<div id="javascript"/>

## JavaScript
> `TODO: `



<div id="kotlin"/>

## Kotlin

> `TODO: `

