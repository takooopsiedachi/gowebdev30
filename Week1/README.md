# 30-Day GoLang Web Development Plan: Zero to Hero

- [Day 1: Development Environment Setup and Visual Studio Code Configuration](#day-1-development-environment-setup-and-visual-studio-code-configuration)
- [Day 2: Learn Go Basics - Part 1](#day-2-learn-go-basics---part-1)
- [Day 3: Learn Go Basics - Part 2](#day-3-learn-go-basics---part-2)
- [Day 4: Learn Go Basics - Part 3](#day-4-learn-go-basics---part-3)
- [Day 5: Study Go Packages and Imports](#day-5-study-go-packages-and-imports)
- [Day 6: Learn About Go's Standard Library](#day-6-learn-about-gos-standard-library)
- [Day 7: Complete Exercises and Practice](#day-7-complete-exercises-and-practice)

## Day 1: Development Environment Setup and Visual Studio Code Configuration

Today's task is to install Go and set up your development environment on macOS using Homebrew. Follow the steps below to complete the setup:

### 1. Install Homebrew

If you haven't installed Homebrew yet, open your terminal and run the following command:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### 2. Install Go using Homebrew

Install Go by running the following command in your terminal:

```bash
brew install go
```

### 3. Set up GOPATH and environment variables

Open your terminal and add the following lines to your .bash_profile, .bashrc, or .zshrc file, depending on which shell you're using:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin
```

Save the file and then run source ~/.bash_profile, source ~/.bashrc, or source ~/.zshrc to reload your shell configuration.

### 4. Install Visual Studio Code using Homebrew

Install Visual Studio Code using Homebrew by running the following command in your terminal:

```bash
brew install --cask visual-studio-code
```

### 5. Install the Go extension for Visual Studio Code

1. Open Visual Studio Code.
2. Click on the Extensions view icon on the sidebar or press Cmd+Shift+X to open the Extensions panel.
3. Search for "Go" in the search bar.
4. Find the extension by "Go Team at Google" and click the Install button.

### 6. Configure Visual Studio Code settings for Go

1. Open Visual Studio Code settings by clicking on the gear icon in the lower left corner and selecting "Settings" or by pressing Cmd+,.
2. In the settings search bar, type "Go" to filter Go-related settings.
3. Customize your Go settings as needed. For example, enable "Go: Use Language Server" to use the Go language server for better code intelligence.

## Day 2: Learn Go Basics - Part 1

Today, you'll focus on learning the basics of Go, including syntax, data types, and variables.

### Study Go syntax

Understand the basic structure of a Go program, including package declarations, imports, and the main function. A simple Go program looks like this:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

- `package main`: Every Go program starts with a package declaration. The `main` package is the entry point of the program.
- `import "fmt"`: This line imports the `fmt` package, which provides functions for formatted I/O.
- `func main()`: The `main` function is the starting point of the program execution.

### Learn about data types

Study Go's basic data types, including integers, floats, booleans, and strings.

- Integers: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- Floats: `float32`, `float64`
- Booleans: `bool`
- Strings: `string`

### Understand variables

Learn how to declare and initialize variables in Go, as well as how to use short variable declarations.

1. Declare a variable with a specific type:

```go
var x int
```

2. Declare a variable with an initial value:

```go
var y int = 42
```

3. Declare a variable without specifying the type (Go will infer the type based on the initial value):

```go
var z = "Hello"
```

4. Short variable declaration (available only within a function):

```go
a := 42
b := "World"
```

Now that you have a basic understanding of Go syntax, data types, and variables, continue to practice and explore these concepts.

## Day 3: Learn Go Basics - Part 2

Today, you'll continue learning Go basics by focusing on constants, loops, and conditionals.

### Learn about constants

Understand the use of constants in Go and how to declare them. Constants are values that cannot be changed once they are initialized.

1. Declare a constant:

```go
const pi = 3.14159
```

2. Declare multiple constants at once:

```go
const (
    red   = "red"
    blue  = "blue"
    green = "green"
)
```

### Study loops

Learn about for loops and range in Go, and understand how to use them for iteration.

1. Basic for loop:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

2. For loop with a condition (similar to a while loop):

```go
n := 1
for n < 100 {
    n *= 2
}
```

3. Iterate over a slice or array using range:

```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

### Understand conditionals

Learn how to use if, else, and switch statements in Go.

1. If statement:

```go
if x > 0 {
    fmt.Println("x is positive")
}
```

2. If-else statement:

```go
if y < 0 {
    fmt.Println("y is negative")
} else {
    fmt.Println("y is non-negative")
}
```

3. If-else if-else statement:

```go
if z < 0 {
    fmt.Println("z is negative")
} else if z > 0 {
    fmt.Println("z is positive")
} else {
    fmt.Println("z is zero")
}
```

4. Switch statement:

```go
switch color {
case "red":
    fmt.Println("The color is red")
case "blue":
    fmt.Println("The color is blue")
case "green":
    fmt.Println("The color is green")
default:
    fmt.Println("Unknown color")
}
```

Now that you have learned about constants, loops, and conditionals in Go, continue to practice and explore these concepts.

## Day 4: Learn Go Basics - Part 3

Today, you'll continue learning Go basics by focusing on functions, error handling, and slices.

### Study functions

Learn how to define and call functions in Go, including how to work with multiple return values and named return values.

1. Define a function:

```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

2. Call a function:

```go
greet("Alice")
```

3. Function with return value:

```go
func square(x int) int {
    return x * x
}
```

4. Function with multiple return values:

```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}
```

5. Function with named return values:

```go
func split(sum int) (x, y int) {
    x = sum / 2
    y = sum - x
    return
}
```

### Learn about error handling

Understand Go's approach to error handling, and learn how to use the error type and the defer statement.

1. Functions that return errors:

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

```

2. Check for errors:

```go
data, err := readFile("file.txt")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Data:", data)
```

3. Defer statement:

```go
func processFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // ... process the file ...
}
```

### Explore slices

Learn about slices in Go, including how to create, access, and modify them.

1. Create a slice:

```go
numbers := []int{1, 2, 3, 4, 5}
```

2. Access elements in a slice:

```go
fmt.Println(numbers[0]) // 1
```

3. Modify elements in a slice:

```go
numbers[1] = 42
```

4. Slice length and capacity:

```go
fmt.Println(len(numbers)) // 5
fmt.Println(cap(numbers)) // 5
```

5. Create a slice with `make`:

```go
newSlice := make([]int, 3, 5) // length: 3, capacity: 5
```

6. Append elements to a slice:

```go
numbers = append(numbers, 6, 7, 8)
```

Now that you have learned about functions, error handling, and slices in Go, continue to practice and explore these concepts.

## Day 5: Study Go packages and imports

### Go module

🐹 Complete the tutorial at https://go.dev/doc/tutorial/

### Learn about packages

Understand the concept of packages in Go, and learn how to create and use custom packages.

1. Create a new directory for your package:

```bash
mkdir -p $GOPATH/src/myproject/mypackage
```

2. Create a new Go file inside the package directory:

```bash
touch $GOPATH/src/myproject/mypackage/mypackage.go
```

3. Define your package and its exported functions:

```go
package mypackage

func MyFunction() {
    // ... implementation ...
}
```

### Study imports

Learn how to import packages into your Go programs, and understand the difference between local and remote imports.

1. Import a package from the standard library:

```go
import "fmt"
```

2. Import a local package:

```go
import "myproject/mypackage"
```

3. Import a remote package:

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })
    router.Run()
}

```

## Day 6: Learn about Go's Standard Library

Today, you'll explore common packages in Go's standard library and practice using them in small programs.

### Explore common packages

Study the functionality provided by common packages in Go's standard library, such as `fmt`, `io`, `bufio`, `os`, `strings`, `strconv`, and `time`.

- `fmt`: Provides formatted I/O functions, similar to C's `printf` and `scanf`.
- `io`: Contains basic interfaces for I/O operations, such as `Reader` and `Writer`.
- `bufio`: Implements buffered I/O for an efficient reading and writing experience.
- `os`: Provides functions for working with the operating system, such as file operations and environment variables.
- `strings`: Contains functions for manipulating strings, such as searching, replacing, and splitting.
- `strconv`: Provides functions for converting strings to and from other data types, such as integers and floats.
- `time`: Contains functionality for working with time and dates, such as parsing, formatting, and calculating durations.

### Practice using standard library packages

Write small Go programs to practice using the standard library packages you learned about.

1. Read a file using `os` and `bufio`:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
```

2. Manipulate strings using `strings` and `strconv`:

```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    s := "Hello, Gopher!"

    fmt.Println(strings.ToUpper(s))
    fmt.Println(strings.ReplaceAll(s, "Gopher", "World"))

    n := 42
    sn := strconv.Itoa(n)
    fmt.Println("The number is: " + sn)
}
```

3. Work with time using `time`:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("Current time:", now)

    future := now.Add(72 * time.Hour)
    fmt.Println("72 hours later:", future)

    layout := "2006-01-02 15:04:05"
    parsedTime, err := time.Parse(layout, "2023-01-01 00:00:00")
    if err != nil {
        fmt.Println("Error parsing time:", err)
        return
    }
    fmt.Println("Parsed time:", parsedTime)
}
```

Continue to explore and practice using Go's standard library packages to become more familiar with their functionality and use cases.

## Day 7: Complete Exercises and Practice

Today, you'll focus on practicing your Go skills by completing exercises from online resources.

### Go through the Go Tour

Complete the interactive exercises in the Go Tour (https://tour.golang.org) to reinforce your understanding of Go basics. The Go Tour covers a wide range of topics, from basic syntax and data types to more advanced concepts such as concurrency and interfaces.

### Practice on exercism.io

Work on Go exercises from exercism.io (https://exercism.io/tracks/go) to gain more hands-on experience. Exercism.io offers a variety of coding exercises that cater to different skill levels, from beginner to advanced. By working through these exercises, you'll get a better understanding of Go's features and best practices, and you'll be able to apply what you've learned to real-world projects.

As you practice, don't forget to consult the Go documentation and ask questions in online forums when you encounter challenges. This will help you solidify your knowledge and become more comfortable with Go as a programming language.
