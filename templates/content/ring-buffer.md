# Queue

## Contents
---

- [Ring buffer description](#description)
- [In practice](#practice)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---

<div id="description"/>

## Description
see [wiki page on Ring Buffer/Circular Buffer](https://en.wikipedia.org/wiki/Circular_buffer)


<div id="practice"/>

## In Practice 
Used often as fast alternative to queue based on arrays.

Also used when you want to see last N records, like `top` command



<div id="java"/>

## Java
```java
package arrays;

import list.Queue;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.NoSuchElementException;

public class EvictingRingBuffer<E> implements Queue<E> {
    private final int maxSize;
    private final ArrayList<E> container;

    private int writeIdx = 0;
    private int readIdx = 0;
    private int count = 0;

    public EvictingRingBuffer(int maxSize) {
        this.maxSize = maxSize;

        container = new ArrayList(Arrays.asList(new Object[maxSize]));
    }

    @Override
    public void enqueue(E element) {
        container.set(writeIdx, element);
        writeIdx = increment(writeIdx);
        if (count == maxSize) {
            readIdx = writeIdx;
        } else {
            count++;
        }
    }

    @Override
    public E dequeue() throws NoSuchElementException {
        if (count == 0) throw new NoSuchElementException();

        E result = container.get(readIdx);

        readIdx = increment(readIdx);
        count--;

        return result;
    }

    public List<E> elements() {
        List<E> result = new ArrayList<>(maxSize);

        int curIdx = readIdx;
        for (int i = 0; i < count; curIdx = increment(curIdx), i++) {
            result.add(container.get(curIdx));
        }
        return result;
    }

    private int increment(int idx) {
        return (idx + 1) % container.size();
    }

}

```



<div id="go"/>

## Go

```go
package arrays

type RingBuffer struct {
	container []interface{}
	readIdx   int
	writeIdx  int
	count     int
	max       int
}

func NewRingBuffer(maxSize int) *RingBuffer {
	buf := new(RingBuffer)
	buf.max = maxSize
	buf.container = make([]interface{}, buf.max)
	return buf
}

func (buf *RingBuffer) Dequeue() (elem interface{}, ok bool) {
	ok = true
	if buf.count == 0 {
		return nil, false
	}
	elem = buf.container[buf.readIdx]

	buf.readIdx = buf.nextIdx(buf.readIdx)
	buf.count--
	return
}

func (buf *RingBuffer) Enqueue(elem interface{}) {
	buf.container[buf.writeIdx] = elem
	buf.writeIdx = buf.nextIdx(buf.writeIdx)
	if buf.count == buf.max {
		buf.readIdx = buf.writeIdx
	} else {
		buf.count++
	}
}

func (buf *RingBuffer) Elements() (result []interface{}) {
	result = make([]interface{}, 0, buf.max)

	curReadIdx := buf.readIdx
	for i := 0; i < buf.count; i++ {
		result = append(result, buf.container[curReadIdx])
		curReadIdx = buf.nextIdx(curReadIdx)
	}
	return
}

func (buf *RingBuffer) nextIdx(curIdx int) int {
	return (curIdx + 1) % buf.max
}

```


<div id="python"/>

## Python


```python
class RingBuffer:
    def __init__(self, input_capacity):
        self.max_size = input_capacity
        self.container = [None] * input_capacity
        self.count = 0
        self.read_idx = 0
        self.write_idx = 0

    def enqueue(self, in_element):
        self.container[self.write_idx] = in_element
        self.write_idx = self._next_idx(self.write_idx)
        if self.count == self.max_size:
            self.read_idx = self.write_idx
        else:
            self.count = self.count + 1

    def dequeue(self):
        result = self.container[self.read_idx]

        self.read_idx = self._next_idx(self.read_idx)
        self.count = self.count - 1

        return result

    def elements(self):
        result = []
        cur_idx = self.read_idx
        for i in range(0, self.count):
            result.append(self.container[cur_idx])
            cur_idx = self._next_idx(cur_idx)
        return result

    def _next_idx(self, cur_idx):
        return (cur_idx + 1) % self.max_size

```



<div id="javascript"/>

## JavaScript

```javascript
class RingBuffer {
  constructor(capacity) {   
    this.container = Array(capacity).fill({})
    this.readIdx = 0
    this.writeIdx = 0
    this.count = 0
    this.maxSize = capacity
  }

  enqueue(element) {
    this.container[this.writeIdx] = element
    this.writeIdx = this.nextIdx(this.writeIdx)
    if (this.count === this.maxSize) {
      this.readIdx = this.writeIdx
    } else {
      this.count++
    }
  }

  dequeue() {
    const result = this.container[this.readIdx]

    this.readIdx = this.nextIdx(this.readIdx)
    this.count--

    return result
  }

  elements() {
    const result = []
    let curIdx = this.readIdx 
    for (i = 0; i < this.count; i++) {
      result.push(this.container[curIdx])
      curIdx = this.nextIdx(curIdx)
    }
    return result
  }

  nextIdx(idx) {
    return (idx + 1) % this.maxSize    
  }
}
```

<div id="kotlin"/>

## Kotlin

```kotlin
package arrays

class EvictingRingBuffer<T>(val maxSize: Int) {
    private val container: Array<T> = arrayOfNulls<Any?>(maxSize) as Array<T>
    private var elementsCount = 0
    private var writeIdx = 0
    private var readIdx = 0

    fun enqueue(element: T) {
        container[writeIdx] = element
        writeIdx = incrementIdx(writeIdx)
        if (elementsCount == maxSize) {
            readIdx = writeIdx
        } else {
            elementsCount++
        }
    }

    fun dequeue(): T {
        require(elementsCount > 0) { "buffer is empty" }

        val element = container[readIdx]
        elementsCount--
        readIdx = incrementIdx(readIdx)

        return element
    }

    fun elements(): Array<T> {
        val result: Array<T> = arrayOfNulls<Any?>(elementsCount) as Array<T>
        var idx = readIdx
        for (i in 0 until elementsCount) {
            result[i] = container[idx]
            idx = incrementIdx(idx)
        }
        return result
    }

    private fun incrementIdx(idx: Int): Int {
        return (idx + 1) % maxSize
    }
}
```