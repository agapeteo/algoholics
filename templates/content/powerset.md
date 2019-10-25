# Power Set

## Contents

---

- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---
 

<div id="java"/>

## Java
dynamic:

```java
package dynamic;

import java.util.*;

public class PowerSet {
    public static <T> List<List<T>> powerSet(List<T> inputSet) {
        List<List<T>> powerSet = new ArrayList<>();

        for (T inputElement: inputSet) {
            ListIterator<List<T>> psIterator = powerSet.listIterator();
            while (psIterator.hasNext()) {
                List<T> curList = new ArrayList<>(psIterator.next());
                curList.add(inputElement);
                psIterator.add(curList);
            }
            powerSet.add(Arrays.asList(inputElement));
        }
        powerSet.add(Collections.emptyList());
        return powerSet;
    }
}
```

combinatorics:
```java
package bits;

import java.util.ArrayList;
import java.util.List;

public class PowerSet {
    public static <T> List<List<T>> powerSet(List<T> inputSet) {
        List<List<T>> powerSet = new ArrayList<>();
        int max = 1 << inputSet.size(); // 2^n, where n is size of inputSet

        for (int intNumber = 0; intNumber < max; intNumber++) {
            List<T> curSet = toSubset(intNumber, inputSet);
            powerSet.add(curSet);
        }
        return powerSet;
    }

    private static <T> List<T> toSubset(int intNumber, List<T> inputSet) {
        List<T> subSet = new ArrayList<>();
        int idx = -1;
        for (int bits = intNumber; bits > 0; bits >>= 1) {
            idx++;

            if ((bits & 1) != 1) continue; // cur bit is 0

            subSet.add(inputSet.get(idx));
        }
        return subSet;
    }
}

```



<div id="go"/>

## Go
dynamic:
```go
package dynamic

func PowerSet(inSet []interface{}) [][]interface{} {
	powerSet := make([][]interface{}, 0)

	for _, elem := range inSet {
		for _, curSet := range powerSet {
			newSet := make([]interface{}, len(curSet))
			copy(newSet, curSet)
			newSet = append(newSet, elem)
			powerSet = append(powerSet, newSet)
		}
		curElemSet := make([]interface{}, 0, 1)
		curElemSet = append(curElemSet, elem)
		powerSet = append(powerSet, curElemSet)
	}
	powerSet = append(powerSet, make([]interface{}, 0))
	return powerSet
}
```

combinatorics:
```go
package bits

func PowerSet(inSet []interface{}) [][]interface{} {
	powerSet := make([][]interface{}, 0)

	max := 1 << uint(len(inSet))
	for i := 0; i < max; i++ {
		powerSet = append(powerSet, setForNUmber(i, inSet))
	}
	return powerSet
}

func setForNUmber(n int, inSet []interface{}) []interface{} {
	set := make([]interface{}, 0)

	i := 0
	num := n

	for num > 0 {
		if (num & 1) != 0 {
			set = append(set, inSet[i])
		}
		num = num >> 1
		i++
	}
	return set
}
```


<div id="python"/>

## Python
```python
def power_set(in_set):
    result_set = []
    for cur_elem in in_set:
        result_set_len = len(result_set)
        for cur_set_idx in range(result_set_len):
            new_set = result_set[cur_set_idx].copy()
            new_set.append(cur_elem)
            result_set.append(new_set)

        result_set.append([cur_elem])

    result_set.append([])

    return result_set
```




<div id="javascript"/>

## JavaScript

```javascript
function powerSet(inSet) {
    const resultSet = [];
    for (const elem of inSet) {
        const resultSetLen = resultSet.length;
        for (let i = 0; i < resultSetLen; i++) {
            const newSet = [...resultSet[i]];
            newSet.push(elem);
            resultSet.push(newSet);
        }
        resultSet.push([elem]);
    }
    resultSet.push([]);
    return resultSet;
}
```




<div id="kotlin"/>

## Kotlin

dynamic:
```kotlin
package dynamic

object PowerSet {

    fun <T> powerSet(inputSet: List<T>): List<List<T>> {
        val powerSet = mutableListOf<List<T>>()

        for (item in inputSet) {
            val powerSetIter = powerSet.listIterator()

            while (powerSetIter.hasNext()) {
                val curSet = powerSetIter.next()
                val newSet = curSet.plus(item)
                powerSetIter.add(newSet)
            }
            powerSet.add(listOf(item))
        }
        powerSet.add(emptyList())

        return powerSet
    }
}
```

combinatorics:
```kotlin
package bits

object PowerSet {

    fun <T> powerSet(inputSet: List<T>): List<List<T>> {

        fun binaryToSet(number: Int): List<T> {
            val curSet = mutableListOf<T>()
            var i = 0
            var curNumber = number
            while (curNumber > 0) {
                if ((curNumber and 1) != 0) {
                    curSet.add(inputSet[i])
                }
                curNumber = curNumber shr 1
                i++
            }
            return curSet
        }

        val powerSet = mutableListOf<List<T>>()
        val max = 1 shl inputSet.size

        for (i in 0 until max) {
            powerSet.add(binaryToSet(i))
        }
        return powerSet
    }
}
```




