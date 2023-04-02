# 30-Day GoLang Web Development Plan: Zero to Hero

- [Day 15: Study the net/http package](#day-15-study-the-nethttp-package)
- [Day 16: Learn about Gorilla Mux](#day-16-learn-about-gorilla-mux)
- [Day 17: Study Go web frameworks - Part 1](#day-17-study-go-web-frameworks---part-1)
- [Day 18: Study Go web frameworks - Part 2](#day-18-study-go-web-frameworks---part-2)
- [Day 19: Learn about middleware in Go web applications](#day-19-learn-about-middleware-in-go-web-applications)
- [Day 20: Work with authentication and authorization - Part 1](#day-20-work-with-authentication-and-authorization---part-1)
- [Day 21: Work with authentication and authorization - Part 2](#day-21-work-with-authentication-and-authorization---part-2)

## Day 15: Study the net/http package

### Understand the basics of the `net/http` package

The `net/http` package is part of the Go standard library and provides essential functionalities to create HTTP servers and clients.

### Learn how to create HTTP servers using `http.ListenAndServe`

To create an HTTP server, use the `http.ListenAndServe` function, which takes two arguments: the address to listen on and a handler to handle incoming requests. A simple example:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
```

### Explore how to create HTTP clients using `http.Get`, `http.Post`, and `http.NewRequest`
To create an HTTP client and make requests, you can use `http.Get`, `http.Post`, or `http.NewRequest`. A simple example using `http.Get`:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response body:", string(body))
}
```

### Study the `http.ResponseWriter` and `http.Request` types

The `http.ResponseWriter` and `http.Request` types are used to write the HTTP response and access information about the incoming request, respectively.

`http.ResponseWriter`: Provides methods to write the response body, set headers, and set the status code. Commonly used methods include `Write`, `WriteHeader`, and `Header`.
`http.Request`: Contains information about the incoming request, such as the HTTP method, URL, headers, and body. Important fields include `Method`, `URL`, `Header`, and `Body`.

## Day 16: Learn about Gorilla Mux

### Study the Gorilla Mux package (https://github.com/gorilla/mux)

Gorilla Mux is a powerful URL router and dispatcher for Go. It provides advanced routing capabilities and is compatible with the `net/http` package.

### Learn how to create a router and register routes with Gorilla Mux

To create a router and register routes with Gorilla Mux, first import the package, create a new router, and then use the `HandleFunc` method to register your routes. For example:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", r)
}
```

### Understand how to use path variables and query parameters

Gorilla Mux allows you to define path variables and access query parameters easily. For example:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		name := r.URL.Query().Get("name")

		fmt.Fprintf(w, "User ID: %s, Name: %s", id, name)
	})

	http.ListenAndServe(":8080", r)
}
```

### Explore middleware support in Gorilla Mux

Gorilla Mux supports middleware, which allows you to add functionality that is executed before or after your route handlers. You can create custom middleware by implementing the `http.Handler` interface:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", r)
}
```

## Day 17: Study Go web frameworks - Part 1

### Explore popular Go web frameworks like Echo and Gin

Two popular web frameworks for Go are Echo (https://echo.labstack.com/) and Gin (https://github.com/gin-gonic/gin). Both are designed to be fast and easy to use, with a focus on simplicity and performance.

### Choose a web framework to learn more in-depth (e.g., Gin)

For this example, we will choose Gin as our web framework to learn more in-depth. Gin is widely used and well-documented, which makes it a good choice for beginners.

### Study the basic structure of a web application using your chosen framework

To get started with Gin, you'll need to install it and import it into your project. Then, create a new Gin router and define your routes and handlers. Here's a basic example:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}
```

### Learn how to create routes, handle requests, and send responses

With Gin, you can create routes using methods like `GET`, `POST`, `PUT`, and `DELETE`. These methods are attached to the router and accept a path and a handler function as arguments. In the handler function, you'll have access to the `*gin.Context` object, which allows you to read request data and send responses.

For example, to create a route that accepts POST requests, you can use the following code:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, "Hello, %s!", name)
	})

	r.Run(":8080")
}
```

## Day 18: Study Go web frameworks - Part 2

### Continue studying your chosen web framework

Keep learning about your chosen web framework (e.g., Gin) by exploring its documentation and examples.

### Understand how to use middleware with the framework

Middleware are functions that can be executed before or after a request handler. They are useful for tasks such as logging, authentication, or data processing. In Gin, you can use middleware by calling the `Use` method on the router:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware example
	r.Use(func(c *gin.Context) {
		// Do something before the handler
		c.Next()
		// Do something after the handler
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}
```

### Learn about error handling and custom error responses

Error handling in Gin can be done using the `AbortWithError` method on the `*gin.Context` object. You can also set custom error responses by using the `c.JSON` method:

```go
package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/error", func(c *gin.Context) {
		err := errors.New("An error occurred")
		c.AbortWithError(http.StatusInternalServerError, err)
	})

	r.GET("/custom-error", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "A custom error message",
		})
	})

	r.Run(":8080")
}
```

### Explore the framework's features for working with data, such as JSON and form data

Gin makes it easy to work with data in your web application. To parse JSON data from a request, use the `BindJSON` method:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "User created", "user": user})
	})

	r.Run(":8080")
}
```

To read form data, you can use methods like `PostForm` or `DefaultPostForm`:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")

		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8080")
}
```

## Day 19: Learn about middleware in Go web applications

### Understand the concept of middleware

Middleware are functions that can be executed before or after a request handler. They are useful for tasks such as logging, authentication, or data processing. Middleware can be added to the request processing pipeline and can modify the request or response objects.

### Learn how to create custom middleware functions

To create a custom middleware function, you need to define a function that takes a `http.Handler` as an argument and returns an `http.Handler`:

```go
func customMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do something before the handler
		next.ServeHTTP(w, r)
		// Do something after the handler
	})
}
```

### Explore how to use middleware for logging, authentication, and request validation

Here's an example of using a middleware for logging:

```go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("%s %s took %s", r.Method, r.URL.Path, duration)
	})
}
```

For authentication:

```go
func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check authentication here (e.g., check for a valid token)
		isAuthenticated := true

		if isAuthenticated {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
```

For request validation:

```go
func validationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request here (e.g., check for required headers)
		isValid := true

		if isValid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
	})
}
```

### Study how to use middleware with Gorilla Mux or your chosen web framework

In Gorilla Mux, you can use middleware by calling the `Use` method on the router:

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Add middleware to the router
	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", r)
}
```

In other web frameworks like Gin, middleware can be added using the `Use` method on the router instance:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Add middleware to the router
	r.Use(loggingMiddleware)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}

func loggingMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	log.Printf("%s %s took %s", c.Request.Method, c.Request.URL.Path, duration)
}
```

## Day 20: Work with authentication and authorization - Part 1

### Learn the basics of authentication and authorization

Authentication is the process of verifying the identity of a user, while authorization is the process of determining what actions or resources the authenticated user is allowed to access.

### Understand different authentication techniques, such as Basic Auth, JWT, and OAuth

- Basic Auth: A simple authentication scheme that sends a username and password with each request, usually encoded in the `Authorization` header.
- JWT (JSON Web Tokens): A token-based authentication method that sends a signed token with each request, typically in the `Authorization` header.
- OAuth: An authorization framework that allows users to grant third-party applications access to their resources without sharing their credentials directly.

### Study how to implement authentication in your Go web application using your chosen framework

For example, using Gin, you can create a simple middleware for JWT authentication:

```go
func jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Validate the token (e.g., check its signature and expiration time)
		// ...

		c.Next()
	}
}
```

### Create a simple login system and protect certain routes

Here's an example of a simple login system and protecting routes using Gin:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Login route
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "test" && password == "password" {
			// Generate a token (e.g., a JWT token) and send it in the response
			// ...
			c.JSON(http.StatusOK, gin.H{"message": "Logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		}
	})

	// Protected route
	authorized := r.Group("/")
	authorized.Use(jwtAuthMiddleware())
	{
		authorized.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, protected route!"})
		})
	}

	r.Run(":8080")
}
```

## Day 21: Work with authentication and authorization - Part 2

### Learn how to implement role-based or permission-based authorization in your Go web application

Role-based or permission-based authorization is the process of granting access to specific resources or actions based on the role or permissions assigned to a user. You can implement this in your Go web application by creating a middleware function that checks the user's role or permissions before granting access.

### Explore how to use middleware for authorization

Here's an example of a role-based authorization middleware using Gin:

```go
func roleBasedAuthorizationMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user's role from the JWT token or another source
		// For example, you can add the user's role as a claim in the JWT token
		userRole := c.GetString("userRole")

		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
```

### Study best practices for securing your web application

- Securely store passwords: Use a strong password hashing algorithm like bcrypt or Argon2 to store user passwords.
- Use HTTPS: Ensure that all connections to your web application are encrypted using HTTPS. You can use tools like Let's Encrypt to obtain free SSL certificates.
- Validate user input: Always validate and sanitize user input to protect against injection attacks, such as SQL injection or cross-site scripting (XSS).

### Review and practice the concepts learned during Week 3

Go through the concepts you learned during Week 3 and complete exercises related to authentication, authorization, middleware, and web frameworks. Identify areas for further study and improvement.
