# Java

### create array

### dynamic array (ArrayList)

### linked list (LinkedList)

### Queue

### Stack

```java
static void sort(int[] array){
            for (int i = 1; i < array.length; i++){
                int iValue = array[i];
                int j = i - 1;
                while (j >= 0 && array[j] > iValue){
                    array[j+1] = array[j];
                    j--;
                }
                array[j+1] = iValue;
            }
        } 
```


