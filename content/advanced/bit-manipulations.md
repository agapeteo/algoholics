+++
title = "Bit manipulations"
template = "level.html"
description = "Bit manipulations"
[extra]
    keywords = "bit manipulations, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, rust, javascript, advanced"
+++

# Bit manipulations

## Contents

---

- [Description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)
- [Rust](#rust)

---


<div id="description" />

## Description
``TODO`` 



<div id="practice"/>

## In practice
Usually should not be used for regular business logic 
Sometime used for:

- fast multiplying or dividing by 2
- hash calculations
- as as compact and fast options check
- etc  



<div id="java"/>

## Java
```java
package bits;

public class BitManipulations {

    public static int twoInBinary() {
        return 0b10;
    }

    public static String toBinaryStr(int number) {
        return Integer.toBinaryString(number);
    }

    public static boolean checkBitSet(int idx, int number) {
        int mask = 1 << idx;
        return ((number & mask) != 0);
    }

    public static int setBit(int idx, int number) {
        int mask = 1 << idx;
        return number | mask;
    }

    public static int clearBit(int idx, int number) {
        int mask = ~(1 << idx);
        return number & mask;
    }

    public static int toggleBit(int idx, int number) {
        int mask = 1 << idx;
        return number ^ mask;
    }

    public static boolean isEven(int number) {
        return (number & 1) == 0;
    }
}
```


<div id="go"/>

## Go

```go
package bits

import (
	"fmt"
	"strconv"
)

func TwoInBinary() int64 {
	base := 2
	bitSize := 8

	n, _ := strconv.ParseInt("10", base, bitSize)
	return n
}

func IntToBinaryStr(number int64) string {
	return fmt.Sprintf("%b", number)
}

func IntToBinaryStrconv(number int64) string {
	return strconv.FormatInt(number, 2)
}

func CheckBitSet(idx uint, n int) bool {
	mask := 1 << idx
	return (n & mask) != 0
}

func SetBit(idx uint, n int) int {
	mask := 1 << idx
	return n | mask
}

func ClearBit(idx uint, n int) int {
	mask := ^(1 << idx)
	return n & mask
}

func toggleBit(idx uint, n int) int {
	mask := 1 << idx
	return n ^ mask
}

func IsEven(n int) bool {
	return (1 & n) == 0
}

```


<div id="python"/>

## Python

```python
def two_in_binary():
    return 0b10


def two_in_binary_2():
    return int("10", 2)


def number_to_binary_str(in_number):
    return format(in_number, "b")


def number_to_binary_str_2(in_number):
    return "{0:b}".format(in_number)


def check_bit(idx, in_number):
    mask = 1 << idx
    return mask & in_number != 0


def set_bit(idx, in_number):
    mask = 1 << idx
    return mask | in_number


def clear_bit(idx, in_number):
    mask = ~(1 << idx)
    return in_number & mask


def toggle_bit(idx, in_number):
    mask = 1 << idx
    return in_number ^ mask


def is_even(in_number):
    return in_number & 1 == 0
```



<div id="javascript"/>

## JavaScript

```javascript
function twoInBinary() {
    return 0b10
}

function twoInBinary2() {
    return parseInt("10", 2)
}

function numberToBinaryStr(number) {
    return number.toString(2)
}

function checkBit(idx, number) {
    const mask = 1 << idx;
    return ((number & mask) != 0)
}

function setBit(idx, number) {
    const mask = 1 << idx;
    return number | mask
}

function clearBit(idx, number) {
    const mask = ~(1 << idx);
    return number & mask
}

function toggleBit(idx, number) {
    const mask = 1 << idx;
    return number ^ mask
}

function isEven(number) {
    return (1 & number) == 0
}
```



<div id="kotlin"/>

## Kotlin

```kotlin
package bits

object BitManipulations {
    fun twoInBinary(): Int {
        return 0b10
    }

    fun numberToBinaryStr(number: Int): String {
        return number.toString(2)
    }

    fun checkBit(idx: Int, number: Int): Boolean {
        return number and (1 shl idx) != 0
    }

    fun setBit(idx: Int, number: Int): Int {
        return number or (1 shl idx)
    }

    fun clearBit(idx: Int, number: Int): Int {
        val mask = (1 shl idx).inv()
        return number and mask
    }

    fun toggleBit(idx: Int, number: Int): Int {
        return number xor (1 shl idx)
    }

    fun isEven(number: Int): Boolean {
        return (1 and number) == 0
    }
}
```

<div id="rust"/>

## Rust

```rust
fn two_in_binary_format() -> usize {
    0b10
}

fn to_binary_str(n: u8) -> String {
    format!("{:08b}", n)
}

fn check_bit_set(n: usize, idx: usize) -> bool {
    ((1 << idx) & n) != 0
}

fn set_bit(n: usize, idx: usize) -> usize {
    (1 << idx) | n
}

fn clear_bit(n: usize, idx: usize) -> usize {
    let mask = !(1 << idx);
    n & mask
}

fn toggle_bit(n: usize, idx: usize) -> usize {
    (1 << idx) ^ n
}

fn is_even(n: usize) -> bool {
    (n & 1) != 1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_in_binary_format() {
        assert_eq!(two_in_binary_format(), 2);
    }

    #[test]
    fn test_to_binary_str() {
        assert_eq!(to_binary_str(2), "00000010")
    }

    #[test]
    fn test_check_bit_set() {
        assert_eq!(check_bit_set(1, 0), true);
        assert_eq!(check_bit_set(1, 1), false);

        assert_eq!(check_bit_set(2, 0), false);
        assert_eq!(check_bit_set(2, 1), true);
        assert_eq!(check_bit_set(2, 2), false);

        assert_eq!(check_bit_set(3, 0), true);
        assert_eq!(check_bit_set(3, 1), true);
        assert_eq!(check_bit_set(3, 2), false);

        assert_eq!(check_bit_set(4, 0), false);
        assert_eq!(check_bit_set(4, 1), false);
        assert_eq!(check_bit_set(4, 2), true);
    }

    #[test]
    fn test_set_bit() {
        assert_eq!(set_bit(1, 0), 1);
        assert_eq!(set_bit(1, 1), 3);

        assert_eq!(set_bit(2, 0), 3);
        assert_eq!(set_bit(2, 1), 2);
        assert_eq!(set_bit(2, 2), 6);

        assert_eq!(set_bit(3, 0), 3);
        assert_eq!(set_bit(3, 1), 3);
        assert_eq!(set_bit(3, 2), 7);
    }

    #[test]
    fn test_toggle_bit() {
        assert_eq!(toggle_bit(1, 0), 0);
        assert_eq!(toggle_bit(1, 1), 3);

        assert_eq!(toggle_bit(2, 0), 3);
        assert_eq!(toggle_bit(2, 1), 0);
        assert_eq!(toggle_bit(2, 2), 6);

        assert_eq!(toggle_bit(3, 0), 2);
        assert_eq!(toggle_bit(3, 1), 1);
        assert_eq!(toggle_bit(3, 2), 7);
    }


    #[test]
    fn test_is_even() {
        assert_eq!(is_even(0), true);
        assert_eq!(is_even(1), false);
        assert_eq!(is_even(2), true);
        assert_eq!(is_even(3), false);
        assert_eq!(is_even(4), true);
        assert_eq!(is_even(11), false);
    }

    #[test]
    fn test_clear_bit() {
        assert_eq!(clear_bit(1, 0), 0);
        assert_eq!(clear_bit(2, 1), 0);
        assert_eq!(clear_bit(7, 1), 5);
    }
}
```

