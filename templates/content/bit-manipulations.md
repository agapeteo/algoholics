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



