## Extra Course: Golang Modules

Golang modules are the official way to manage dependencies in your Go projects. In this extra course, you'll learn how to create and use Go modules.

### Initialize a Go module

To create a new Go module, navigate to your project's root directory and run the following command:

```bash
go mod init <module-name>
```

This command creates a new `go.mod` file that lists your module's name and its dependencies.

### Add dependencies

To add a dependency to your project, simply import it in your Go code:

```go
import (
    "github.com/gorilla/mux"
)
```

When you run `go build` or `go run`, Go will automatically download the required dependencies and update your `go.mod` and `go.sum` files.

### Update dependencies

To update a specific dependency to a newer version, use the `go get` command followed by the package path and version:

```bash
go get github.com/gorilla/mux@v1.8.0
```

To update all dependencies to their latest versions, run:

```bash
go get -u
```

### Remove unused dependencies

To remove unused dependencies from your `go.mod` file, run the following command:

```bash
go mod tidy
```

This command removes any packages that are not needed by your project.

### Work with multiple versions of a dependency

In some cases, you may need to work with multiple versions of a dependency. You can use the `replace` directive in your `go.mod` file to achieve this:

```
replace (
    github.com/gorilla/mux => github.com/gorilla/mux/v2 v2.0.1
)
```

This directive tells Go to replace the imported version of `github.com/gorilla/mux` with the specified version (`v2.0.1`).

By understanding and using Golang modules, you can effectively manage your project's dependencies and ensure that your applications are reliable and maintainable.

## Creating Your Own Package and Versioning Workflow

In this example, you'll learn how to create your own Go package and manage its releases and versions using a common workflow.

### Creating a Go package

1. Create a new directory for your package:

```bash
mkdir mypackage
```

2. Change to the newly created directory:

```bash
cd mypackage
```

3. Initialize a new Go module:

```bash
go mod init github.com/yourusername/mypackage
```

4. Create a `mypackage.go` file with the following content:

```go
package mypackage

import "fmt"

func Hello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Releasing a new version of your package

1. Initialize a Git repository in your package directory:

```bash
git init
```

2. Add your files to the Git repository and commit:

```bash
git add .
git commit -m "Initial commit"
```

3. Create a new GitHub repository and push your code to it:

```bash
git remote add origin https://github.com/yourusername/mypackage.git
git push -u origin master
```

4. Create a new version tag for your package:

```bash
git tag v0.1.0
git push origin v0.1.0
```

This tag represents the first release (v0.1.0) of your package.

### Using your package in another project

1. In another Go project, import your package:

```go
import (
    "github.com/yourusername/mypackage"
)
```

2. Use your package's `Hello` function:

```go
mypackage.Hello("John")
```

When you build or run your project, Go will automatically download your package and include it as a dependency in the `go.mod` file.

### Updating your package version

1. Make changes to your package, such as adding new functions or improving existing ones.
2. Commit your changes:

```bash
git add .
git commit -m "Add new features"
```

3. Create a new version tag and push it to GitHub:

```bash
git tag v0.2.0
git push origin v0.2.0
```

Users of your package can now update to the new version by running `go get`:

```bash
go get github.com/yourusername/mypackage@v0.2.0
```

By following this workflow, you can create your own Go packages and manage their versions effectively. This will help you maintain a clean and organized codebase and easily share your work with others.
