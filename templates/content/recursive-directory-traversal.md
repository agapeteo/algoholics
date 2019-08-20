# Recursive directory traversal

## Contents

---

- [Complexity](#complexity)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [JavaScript](#javascript)
- [Kotlin](#kotlin)

---


<div id="complexity"/>


## Complexity
`O(N)`

 

<div id="java"/>

## Java
```java
public static void main(String[] args) {
    printDirContent("/path/to/folder/", true);
}

static void printDirContent(String dir, boolean ignoreHidden) {
    System.out.println("directory structure of " + dir);
    printDirContent(new File(dir), 0, ignoreHidden);
}

static void printDirContent(File dir, int tabsCount, boolean ignoreHidden) {

    if (!dir.exists()) throw new IllegalArgumentException(dir + " does not exist");
    if (!dir.isDirectory()) throw new IllegalArgumentException(dir + " is not directory");

    StringBuilder tabs = new StringBuilder();
    for (int i = 0; i < tabsCount; i++) {
        tabs.append("\t");
    }

    StringBuilder dirOutput = new StringBuilder();
    dirOutput
            .append(tabs)
            .append(dir.getName())
            .append("/");
    System.out.println(dirOutput.toString());

    for (File file : dir.listFiles()) {
        if (ignoreHidden && file.getName().startsWith(".")) continue;

        if (file.isDirectory()) {
            printDirContent(file, tabsCount + 1, ignoreHidden);
        } else {
            StringBuilder fileOutput = new StringBuilder();
            fileOutput
                    .append("\t")
                    .append(tabs)
                    .append(file.getName());
            System.out.println(fileOutput.toString());
        }
    }
}
```



<div id="go"/>

## Go
```go
func printDirContent(dir string, tabs int, ignoreHidden bool) {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("dir does not exist: %s", dir)
		}
		log.Fatal(err)
	}

	if !info.IsDir() {
		log.Fatalf("%s is not directory", dir)
	}

	tabsBuilder := strings.Builder{}
	for i := 0; i < tabs; i++ {
		tabsBuilder.WriteString("\t")
	}
	dirParts := strings.Split(dir, "/")
	fmt.Printf("%s%s/\n", tabsBuilder.String(), dirParts[len(dirParts) - 1])

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			if ignoreHidden && strings.HasPrefix(f.Name(), ".") {
				continue
			}
			printDirContent(filepath.Join(dir, f.Name()), tabs+1, ignoreHidden)
		} else {
			fmt.Printf("\t%s%s\n", tabsBuilder.String(), f.Name())
		}
	}
}
```
> Note: This example is given for learning recursive approach 
> 
> Go already has build-in [`filepath.Walk`](https://golang.org/pkg/path/filepath/#example_Walk)


<div id="python"/>

## Python
```python
#  TODO:
```


<div id="javascript"/>

## JavaScript
```javascript
// TODO:
```


<div id="kotlin"/>

## Kotlin
```kotlin
fun printDirContent(dir: File, tabsCount: Int = 0, ignoreHidden: Boolean = true) {
    require(dir.exists()) { "$dir must exist" }
    require(dir.isDirectory) { "$dir must be directory" }

    val tabs = StringBuilder()
    repeat(tabsCount) { tabs.append("\t") }

    val dirOutput = StringBuilder()
    dirOutput.append(tabs).append(dir.name).append("/")
    
    println(dirOutput.toString())

    for (file in dir.listFiles()) {
        if (ignoreHidden && file.name.startsWith(".")) continue

        if (file.isDirectory) {
            printDirContent(file, tabsCount + 1, ignoreHidden)
        } else {
            val fileOutput = StringBuilder()
            fileOutput.append("\t").append(tabs).append(file.name)
            
            println(fileOutput.toString())
        }
    }
}
```