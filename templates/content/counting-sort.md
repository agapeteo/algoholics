# Counting sort

## Contents

---

- [Description](#description)
- [Complexity](#complexity)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---

<div id="description"/>

## Description
Stable sorting algorithm of integer numbers which doesn't require comparison



<div id="complexity"/>

## Complexity
linear 
`O(N+k)` where `k` is max number

space complexity `O(N)`



<div id="practice"/>
Used for sorting integers where max number is not too big



<div id="java"/>

## Java
list based:
```java
 package sort;
 
 import java.util.ArrayList;
 import java.util.List;
 
 public class ListBasedCountingSort {
 
     public static void sort(int[] array) {
         List<Integer>[] countingArr = new List[max(array) + 1];
 
         for (int i = 0; i < array.length; i++) {
             List<Integer> curList = countingArr[array[i]];
             if (curList == null) {
                 curList = new ArrayList<>();
                 curList.add(array[i]);
                 countingArr[array[i]] = curList;
             } else {
                 curList.add(array[i]);
             }
         }
 
         int idx = 0;
         for (int i = 0; i < countingArr.length; i++) {
             List<Integer> curList = countingArr[i];
 
             if (curList == null) continue;
 
             for (int j = 0; j < curList.size(); j++) {
                 array[idx] = curList.get(j);
                 idx++;
             }
         }
     }
 
     private static int max(int[] array) {
         int max = Integer.MIN_VALUE;
 
         for (int i = 0; i < array.length; i++) {
             if (array[i] > max) {
                 max = array[i];
             }
         }
         return max;
     }
 }

```


without lists: 
```java
package sort;

public class CountingSort {

    public static int[] sort(int[] array) {
        int[] countingArr = new int[max(array) + 1];

        for (int i = 0; i < array.length; i++) {
            int idxByValue = array[i];
            countingArr[idxByValue]++;
        }
        int sum = 0;
        for (int i = 0; i < countingArr.length; i++) {
            sum += countingArr[i];
            countingArr[i] = sum;
        }

        int[] resultArr = new int[array.length];

        for (int i = 0; i < array.length; i++) {
            int curValue = array[i];
            int targetIdx = countingArr[curValue] - 1;
            resultArr[targetIdx] = curValue;
            countingArr[curValue]--;
        }
        return resultArr;
    }

    private static int max(int[] array) {
        int max = Integer.MIN_VALUE;

        for (int i = 0; i < array.length; i++) {
            if (array[i] > max) {
                max = array[i];
            }
        }
        return max;
    }
}

```



<div id="go"/>

## Go
```go
package sort

func max(s []int) int {
	max := s[0]
	for _, v := range s {
		if max < v {
			max = v
		}
	}
	return max
}

func CountingSort(s []int) {
	size := max(s) + 1
	countSlice := make([][]int, size)

	for _, v := range s {
		curSlice := &countSlice[v]
		*curSlice = append(*curSlice, v)
	}

	idx := 0
	for _, curSlice := range countSlice {
		for _, v := range curSlice {
			s[idx] = v
			idx++
		}
	}
}
```



<div id="python"/>

## Python
```python
def max_int(in_list):
    result = None
    for v in in_list:
        if result < v:
            result = v
    return result


def sort(in_list):
    counting_list_size = max(in_list) + 1
    counting_list = [None] * counting_list_size

    for idx_by_value in in_list:
        cur_list = counting_list[idx_by_value]
        if cur_list:
            cur_list.append(idx_by_value)
        else:
            cur_list = [idx_by_value]
            counting_list[idx_by_value] = cur_list

    cur_idx = 0
    for cur_list in counting_list:
        if not cur_list:
            continue

        for cur_value in cur_list:
            in_list[cur_idx] = cur_value
            cur_idx += 1
```


<div id="javascript"/>

## JavaScript
```javascript
function max(arr) {
    let max = null;
    for (let v of arr) {
        if (max < v) {
            max = v;
        }
    }
    return max
}

function sort(arr) {
    let countingSize = max(arr) + 1;
    const countingArr = Array(countingSize).fill(null);

    for (let v of arr) {
        let curArr = countingArr[v];
        if (!curArr) {
            curArr = [v];
            countingArr[v] = curArr
        } else {
            curArr.push(v)
        }
    }

    let idx = 0;
    for (let curArr of countingArr) {
        if (!curArr) continue;
        for (let v of curArr) {
            arr[idx] = v;
            idx++;
        }
    }
}
```


<div id="kotlin"/>

## Kotlin
list based:

```kotlin
package sort

object ListBasedCountingSort {

    fun sort(array: IntArray) {
        val countingList = arrayOfNulls<MutableList<Int>>(max(array) + 1)
        for (intValue in array) {
            var listByIdx = countingList[intValue]
            if (listByIdx == null) {
                listByIdx = mutableListOf(intValue)
                countingList[intValue] = listByIdx
            } else {
                listByIdx.add(intValue)
            }
        }
        var i = 0
        for (eachList in countingList) {
            if (eachList == null) continue

            for (eachValue in eachList) {
                array[i] = eachValue
                i++
            }
        }
    }

    private fun max(array: IntArray): Int {
        var max = Int.MIN_VALUE
        for (intValue in array) {
            if (intValue > max) {
                max = intValue
            }
        }
        return max
    }
}
```

without using lists:
```kotlin
package sort

object CountingSort {
    fun sort(array: IntArray): IntArray {
        val max = max(array)
        val countArr = IntArray(max + 1)
        for (i in array) {
            countArr[i]++
        }
        var runningSum = 0
        for (i in countArr.indices) {
            runningSum += countArr[i]
            countArr[i] = runningSum
        }

        val result = IntArray(array.size)
        for (i in array.size - 1 downTo 0) {
            val curValue = array[i]
            val curCount = countArr[curValue]
            result[curCount - 1] = curValue
            countArr[curValue]--
        }
        return result
    }

    private fun max(array: IntArray): Int {
        var max = Int.MIN_VALUE
        for (intValue in array) {
            if (intValue > max) {
                max = intValue
            }
        }
        return max
    }
}
```