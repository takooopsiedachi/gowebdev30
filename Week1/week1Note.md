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

### 7. Test your Go environment configuration

To ensure that your Go environment is set up correctly, follow these steps:

1. Create a new directory for your Go workspace:

```bash
mkdir -p $GOPATH/src/hello
```

2. Change to the newly created directory:

```bash
cd $GOPATH/src/hello
```

3. Create a new file named main.go with the following content:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

4. Build and run the program:

```bash
go build
./hello
```

If your Go environment is configured correctly, you should see the following output:

```
Hello, Go!
```

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
