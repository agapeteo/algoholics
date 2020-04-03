+++
title = "Recursive directory traversal"
template = "level.html"
description = "Recursive directory traversal"
[extra]
    keywords = "Recursive directory traversal, recursion, learn, explain, tutorial, data structures, algorithms, graph, tree, interview, java, golang, kotlin, python, javascript, rust"
+++

# Recursive directory traversal

## Contents

---

- [Complexity](#complexity)
- [Java](#java)
- [Go](#go)
- [Python](#python)
- [Kotlin](#kotlin)
- [Rust](#rust)

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
import os


def print_dir_content(dir_abs, tabs=0, ignore_hidden=True):
    for f in os.listdir(dir_abs):
        if ignore_hidden and f.startswith("."):
            continue

        tabs_str = "".join("\t" * tabs)
        print tabs_str + f + "/"

        full_path = os.path.join(dir_abs, f)
        if os.path.isdir(full_path):
            print_dir_content(full_path, tabs + 1)
        else:
            print "\t" + tabs_str + f


print_dir_content("/full/path/to/directory/")
```
> Note: This example is given for learning recursive approach 
> 
> Python already has build-in [`os.walk`](https://docs.python.org/3/library/os.html#os.walk)


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


<div id="rust"/>

## Rust
```rust
use std::path::Path;

pub fn traverse_dirs(dir_path: &Path, ignore_hidden: bool) {
    if !dir_path.exists() || !dir_path.is_dir() {
        panic!(format!("incorrect path {}", dir_path.to_str().unwrap()));
    }
    traverse_dir(dir_path, 0, ignore_hidden);
}

fn traverse_dir(path: &Path, tabs_count: usize, ignore_hidden: bool) {
    let f_name = path.file_name().unwrap().to_str().unwrap_or("?invalid_file_name?");
    if ignore_hidden && f_name.starts_with('.') {
        return;
    }
    let tabs_str = "\t".repeat(tabs_count);

    let f_name = tabs_str + f_name;

    println!("{}", f_name);

    if path.is_dir() {
        for f in path.read_dir().unwrap() {
            traverse_dir(&f.unwrap().path(), tabs_count + 1, ignore_hidden);
        }
    }
}

#[test]
fn test() {
    traverse_dirs(Path::new("/Users/some_user/some_folder"), true);
}
```