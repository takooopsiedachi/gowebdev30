# 30-Day GoLang Web Development Plan: Zero to Hero

- [Day 8: Learn HTTP basics](#day-8-learn-http-basics)
- [Day 9: Study RESTful APIs](#day-9-study-restful-apis)
- [Day 10: Learn about JSON](#day-10-learn-about-json)
- [Day 11: Introduction to databases in Go](#day-11-introduction-to-databases-in-go)
- [Day 12: Working with SQL databases in Go](#day-12-working-with-sql-databases-in-go)
- [Day 13: Working with NoSQL databases in Go](#day-13-working-with-nosql-databases-in-go)
- [Day 14: Practice and review](#day-14-practice-and-review)

# Day 8: Learn HTTP basics

## HTTP Protocol

HTTP (Hypertext Transfer Protocol) is the foundation of data communication on the World Wide Web. It is a request-response protocol that allows clients (e.g., web browsers) to request resources (e.g., HTML documents, images, or videos) from servers.

## Request Methods

HTTP defines several request methods, each with a specific purpose:

- `GET`: Retrieve a resource from the server.
- `POST`: Submit data to the server, usually to create a new resource.
- `PUT`: Update an existing resource on the server.
- `DELETE`: Delete a resource from the server.

## HTTP Status Codes

HTTP status codes are three-digit numbers that indicate the outcome of an HTTP request:

- `1xx` (Informational): The request was received, and the server is continuing to process it.
- `2xx` (Successful): The request was successfully received, understood, and accepted.
- `3xx` (Redirection): Further action needs to be taken to complete the request.
- `4xx` (Client Error): The request contains bad syntax or cannot be fulfilled by the server.
- `5xx` (Server Error): The server failed to fulfill a valid request.

Some common status codes include:

- `200 OK`: The request has succeeded.
- `201 Created`: The request has succeeded, and a new resource was created as a result.
- `400 Bad Request`: The server cannot process the request due to client error.
- `404 Not Found`: The requested resource could not be found on the server.
- `500 Internal Server Error`: The server encountered an error while processing the request.

## HTTP Headers

HTTP headers are key-value pairs sent with an HTTP request or response. They provide metadata about the request or response, such as content type, content length, or caching directives.

Common HTTP headers include:

- `Content-Type`: Specifies the media type of the resource, such as text/html, application/json, or image/jpeg.
- `Content-Length`: Indicates the size of the response body in bytes.
- `Cache-Control`: Specifies caching directives for the resource, such as no-cache, private, or max-age.
- `Authorization`: Provides authentication credentials for the request, usually in the format Authorization: Bearer <token>.
- `User-Agent`: Contains information about the client making the request, such as the browser and operating system.

By understanding the basics of HTTP, you will be better equipped to build and debug web applications, interact with APIs, and troubleshoot issues that may arise during development.

## Go's `net/http` Package

The `net/http` package in Go's standard library provides functionality for building HTTP clients and servers. It includes functions for making requests, handling responses, and creating HTTP servers.

### Example: Simple HTTP Server

Here's an example of a simple HTTP server using Go's `net/http` package:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":8080", nil)
}
```

This example creates an HTTP server that listens on port 8080 and responds with "Hello, world!" for every request.

### Example: Simple HTTP Client

Here's an example of a simple HTTP client using Go's net/http package:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
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

	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
}
```

This example sends a GET request to "http://example.com" and prints the status code and response body to the console.

By studying the HTTP protocol, request methods, status codes, headers, and the `net/http` package in Go, you'll have a solid foundation for building and interacting with web applications in Go.

# Day 9: Study RESTful APIs

## Principles of RESTful API Design

REST (Representational State Transfer) is an architectural style for designing networked applications. RESTful APIs use HTTP methods to manipulate resources, allowing clients to interact with a server by performing actions such as creating, reading, updating, and deleting resources.

Here are some principles of RESTful API design:

1. Stateless: Each request from a client to a server must contain all the information needed to understand and process the request. The server should not store information about the client between requests.

2. Client-Server: The client and server are separate entities that communicate over a network. The client is responsible for the user interface, while the server manages resources and data.

3. Cacheable: Responses from the server can be cached by the client to improve performance.

4. Layered System: The architecture can be composed of multiple layers, with each layer having a specific responsibility. This separation of concerns can improve scalability and maintainability.

5. Uniform Interface: The API should have a consistent interface, making it easier for clients to interact with the server.

### 🍰 Layered System

- A RESTful API should be designed using a layered system approach.
- Each layer has a specific responsibility and can be maintained independently of others.
- This separation of concerns allows for better maintainability, flexibility, and scalability.
- Common layers in a RESTful API architecture:
  - Presentation layer: Responsible for handling user interfaces and interactions.
  - Application layer: Handles business logic, processing requests, and managing resources.
  - Data access layer: Manages data storage and retrieval.

### 🔗 Uniform Interface

- A key principle in RESTful API design is having a consistent and uniform interface.
- The uniform interface simplifies the architecture by allowing different components to interact using a standardized set of methods and conventions.
- Benefits of a uniform interface include:
  - Improved simplicity and consistency for developers.
  - Enhanced visibility of interactions between components.
  - Greater decoupling between the client and server.
- Key constraints of a uniform interface:
  - Resource identification: Each resource should be uniquely identifiable using a consistent naming convention (e.g., URIs).
  - Resource manipulation: The API should support standard methods (e.g., GET, POST, PUT, DELETE) for interacting with resources.
  - Self-descriptive messages: Messages should include enough information for the recipient to understand how to process the message (e.g., by including metadata).
  - Hypermedia as the engine of application state (HATEOAS): Responses should include links to related resources, allowing clients to navigate the API dynamically.


## Architectural Style of RESTful APIs

RESTful APIs are built around the concept of resources, which are identified by unique URLs. Clients interact with these resources using standard HTTP methods (GET, POST, PUT, DELETE) and communicate with the server using standard media types (such as JSON or XML).

## Design Constraints of RESTful APIs

Some design constraints of RESTful APIs include:

1. Resource identification: Resources should be uniquely identifiable using URLs.

2. Resource manipulation: Resources should be manipulated using standard HTTP methods.

3. Self-descriptive messages: Requests and responses should include enough information for the recipient to understand their purpose.

4. Hypermedia as the engine of application state (HATEOAS): Clients should be able to navigate the API by following hyperlinks provided in responses.

## Implementation Best Practices

When implementing a RESTful API, consider the following best practices:

1. Use appropriate HTTP methods: Use GET for retrieving resources, POST for creating new resources, PUT for updating existing resources, and DELETE for removing resources.

2. Use proper HTTP status codes: Return appropriate status codes to indicate the outcome of a request (e.g., 200 OK, 201 Created, 400 Bad Request, 404 Not Found).

3. Use descriptive URLs: Design URLs that are easy to understand and describe the resource being accessed.

4. Support versioning: Include a version number in the URL or header to allow for changes in the API without breaking existing clients.

5. Use standard media types: Use standard media types (such as JSON or XML) for requests and responses.

6. Document your API: Provide clear and concise documentation to help clients understand how to use your API.

7. Implement security measures: Protect your API using authentication, authorization, and encryption mechanisms as needed.

## Gin Framework 🍸

Gin is a high-performance, minimalist web framework for the Go programming language. It's designed for building RESTful APIs and web applications with ease and speed. Gin provides a simple, expressive API with built-in support for routing, middleware, and data binding.

Github: https://github.com/gin-gonic/gin
Website: https://gin-gonic.com/

# Day 10: Learn about JSON

## JSON Data Format

JSON (JavaScript Object Notation) is a lightweight data interchange format that is easy for humans to read and write, and easy for machines to parse and generate. JSON is a text format that is completely language-independent but uses conventions familiar to programmers of the C family of languages, including C, C++, C#, Java, JavaScript, Perl, Python, and many others.

JSON data is represented as key-value pairs, where keys are strings and values can be strings, numbers, booleans, objects, or arrays. Objects are delimited by curly braces (`{}`) and key-value pairs within objects are separated by commas. Arrays are delimited by square brackets (`[]`) and values within arrays are also separated by commas.

## Parsing and Generating JSON Data

Parsing JSON data involves converting a JSON-formatted string into a data structure that can be manipulated by your programming language, while generating JSON data involves converting a data structure into a JSON-formatted string.

## Go's encoding/json Package

In Go, the `encoding/json` package provides functionality for encoding and decoding JSON data. Here are some common functions and methods in the package:

- `json.Marshal`: Converts a Go data structure (such as a struct or a map) into a JSON-formatted byte slice.
- `json.Unmarshal`: Parses a JSON-formatted byte slice into a Go data structure (such as a struct or a map).
- `json.NewEncoder`: Creates a new JSON encoder that writes to a specified `io.Writer`.
- `json.NewDecoder`: Creates a new JSON decoder that reads from a specified `io.Reader`.

## Practice Working with JSON Data in Go

Here's a simple example of working with JSON data in Go:

1. Define a struct to represent your data:

```go
type Person struct {
    Name string json:"name"
    Age int json:"age"
}
```

2. Create an instance of the struct and marshal it into JSON:

```go
p := Person{Name: "Alice", Age: 30}
jsonData, err := json.Marshal(p)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))
```

3. Parse JSON data into a struct:

```go
jsonData := []byte({"name": "Bob", "age": 25})
var p Person
err := json.Unmarshal(jsonData, &p)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", p)
```

By learning about JSON and practicing with Go's `encoding/json` package, you'll be able to work with JSON data effectively in your Go applications.

# Day 11: Introduction to Databases in Go

## Basics of Databases

A database is an organized collection of data, stored and accessed electronically. Databases are used to store and manage large amounts of data efficiently, and they provide various functions to retrieve, insert, update, and delete data.

## 🗄️ Database Key Concepts

### 1. Database Management System (DBMS)
- A software application that enables users to create, manage, and manipulate databases.
- Examples: MySQL, PostgreSQL, Oracle Database, Microsoft SQL Server, SQLite, MongoDB.

### 2. Data Model
- A method of organizing and representing data within a database.
- Common data models: Relational, Hierarchical, Network, Object-oriented, Document, Graph.

### 3. Schema
- The structure of a database, including tables, columns, data types, constraints, and relationships.
- Acts as a blueprint for how the database is organized.

### 4. Table
- A collection of related data organized into rows and columns.
- In a relational database, tables are used to store data about entities.

### 5. Column
- A vertical structure within a table that holds a specific type of data.
- Each column has a name, data type, and optional constraints (e.g., NOT NULL, UNIQUE).

### 6. Row (Record)
- A horizontal structure within a table that represents a single instance of the data.
- Rows are made up of one or more columns and store the actual data in a table.

### 7. Primary Key
- A unique identifier for each row in a table.
- Ensures that each row can be distinguished from all others.

### 8. Foreign Key
- A column (or set of columns) in one table that refers to the primary key of another table.
- Used to establish relationships between tables and enforce referential integrity.

### 9. Index
- A database object that improves the speed of data retrieval.
- Acts as a lookup table or reference for the actual data in a table.

### 10. SQL (Structured Query Language)
- A standardized programming language for managing relational databases.
- Used for creating, querying, updating, and deleting data, as well as managing database structures.

## SQL and NoSQL Databases

There are two main types of databases: SQL (Structured Query Language) and NoSQL (Not Only SQL).

1. SQL databases are relational databases that use a schema to define the structure of the data. They store data in tables, with each table containing rows and columns. Examples of SQL databases include PostgreSQL, MySQL, and SQLite.

2. NoSQL databases are non-relational databases that can store data in various formats, such as key-value, document, column-family, and graph. They do not have a fixed schema, and they are designed to scale horizontally. Examples of NoSQL databases include MongoDB, Redis, and Cassandra.

## 🗄️ SQL vs NoSQL Databases: Pros, Cons, and Comparison

### SQL Databases

#### Pros
- **Structured Data**: SQL databases are designed for structured data, making them a good fit for applications with well-defined data models.
- **ACID Compliance**: SQL databases are ACID-compliant, ensuring data consistency, integrity, and reliability.
- **Strong Schema**: The schema enforces data integrity and relationships, leading to fewer data inconsistencies.
- **Mature Technology**: SQL databases have a long history and are well-established, offering robust tools, libraries, and community support.

#### Cons
- **Limited Scalability**: SQL databases can struggle with horizontal scalability, making it challenging to handle large volumes of data or high write loads.
- **Rigid Schema**: The fixed schema can be inflexible and time-consuming to modify when data models evolve.
- **Complex Queries**: SQL queries can become complex and difficult to optimize for performance, particularly with large datasets or complex data relationships.

### NoSQL Databases

#### Pros
- **Flexible Data Model**: NoSQL databases support semi-structured or unstructured data, making them a good fit for applications with dynamic or evolving data models.
- **Scalability**: NoSQL databases are designed for horizontal scaling, handling large volumes of data or high write loads more efficiently.
- **Schema-less**: The absence of a fixed schema allows for more flexibility and agility in handling diverse or changing data structures.
- **Variety of Types**: NoSQL databases come in various types (e.g., document, key-value, column-family, graph), offering diverse options to suit specific use cases.

#### Cons
- **Eventual Consistency**: Some NoSQL databases are not ACID-compliant, opting for eventual consistency to achieve better performance and scalability. This can lead to temporary data inconsistencies.
- **Less Mature**: NoSQL databases are relatively newer and may lack the maturity, tools, libraries, and community support available for SQL databases.
- **Not Ideal for Complex Relationships**: Some types of NoSQL databases are not well-suited for handling complex data relationships or multi-table join operations.

### Comparison

- **Data Model**: SQL databases use a relational model, while NoSQL databases use various data models, such as document, key-value, column-family, or graph.
- **Schema**: SQL databases have a fixed schema, whereas NoSQL databases can be schema-less or have a flexible schema.
- **Scalability**: SQL databases are generally better suited for vertical scaling, while NoSQL databases are designed for horizontal scaling.
- **Consistency**: SQL databases are typically ACID-compliant, ensuring strong consistency, whereas some NoSQL databases opt for eventual consistency.
- **Use Cases**: SQL databases are well-suited for applications with well-defined data models and complex relationships, while NoSQL databases are better suited for applications with dynamic, evolving, or diverse data models and high scalability requirements.


## Go's database/sql Package

The `database/sql` package in Go provides a general interface for working with SQL databases. It defines several types, functions, and methods that allow you to connect to a database, execute queries, and retrieve results.

Here are some common functions and methods in the `database/sql` package:

- `sql.Open`: Opens a new database connection using a specified driver and connection string.
- `sql.DB.Query`: Executes a query that returns rows.
- `sql.DB.Exec`: Executes a query that does not return rows (such as an INSERT, UPDATE, or DELETE statement).
- `sql.Rows.Scan`: Copies the values of the current row into specified variables.
- `sql.DB.Prepare`: Prepares an SQL statement for later execution.

## Database Drivers in Go

In Go, database drivers are used to communicate with different databases. The drivers implement the interfaces defined by the `database/sql` package. To use a specific database, you need to import the appropriate driver package and register it with the `database/sql` package.

For example, to work with PostgreSQL, you can use the `github.com/lib/pq` driver:

```go
import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=mysecretpassword dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ...
}
```

# Day 12: Working with SQL Databases in Go

## Set Up a SQL Database

Choose a SQL database such as PostgreSQL or MySQL to set up on your local machine. Follow the installation instructions for your chosen database, and create a new database and table to use in your Go application.

## Connect to a Database Using Go's database/sql Package

To connect to a SQL database using Go, you will need to use the `database/sql` package and import the appropriate database driver.

For example, to connect to a PostgreSQL database, you would need to import the `github.com/lib/pq` driver:

```go
import (
	"database/sql"
	_ "github.com/lib/pq"
)
```

Then, create a connection string with your database information and use the sql.Open function to establish a connection:

```go
connStr := "user=yourusername password=yourpassword dbname=yourdbname sslmode=disable"
db, err := sql.Open("postgres", connStr)
if err != nil {
	log.Fatal(err)
}
defer db.Close()
```

## Perform CRUD Operations Using Go
Once you have connected to your database, you can perform CRUD (Create, Read, Update, Delete) operations using the database/sql package.

### Create
To insert data into a table, use the sql.DB.Exec method:

```go
_, err := db.Exec("INSERT INTO your_table (column1, column2) VALUES ($1, $2)", value1, value2)
if err != nil {
	log.Fatal(err)
}
```

### Read
To query data from a table, use the sql.DB.Query method and sql.Rows.Scan to iterate over the result set:

```go
rows, err := db.Query("SELECT column1, column2 FROM your_table")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
	var column1 string
	var column2 int

	err := rows.Scan(&column1, &column2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Column1: %s, Column2: %d\n", column1, column2)
}
```

### Update
To update data in a table, use the sql.DB.Exec method:

```go
_, err := db.Exec("UPDATE your_table SET column1 = $1 WHERE column2 = $2", newValue1, conditionValue2)
if err != nil {
	log.Fatal(err)
}
```

### Delete
To delete data from a table, use the sql.DB.Exec method:

```go
_, err := db.Exec("DELETE FROM your_table WHERE column1 = $1", conditionValue1)
if err != nil {
	log.Fatal(err)
}
```

By understanding how to set up a SQL database, connect to it using Go's database/sql package, and perform CRUD operations, you will be well-equipped to work with SQL databases in your Go applications.

# Day 13: Working with NoSQL Databases in Go

## Set Up a NoSQL Database

Choose a NoSQL database such as MongoDB or Redis to set up on your local machine. Follow the installation instructions for your chosen database. For this example, we will use Redis.

## Connect to a NoSQL Database Using a Go Driver

To connect to a NoSQL database using Go, you will need to use a Go driver for your chosen database.

For example, to connect to a Redis database, you would need to import the `github.com/go-redis/redis/v8` driver:

```go
import (
	"context"
	"github.com/go-redis/redis/v8"
)
```

Then, create a new Redis client and use the redis.NewClient function to establish a connection:

```go
ctx := context.Background()
client := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

_, err := client.Ping(ctx).Result()
if err != nil {
	log.Fatal(err)
}
defer client.Close()
```

## Perform CRUD Operations Using Go and a NoSQL Database
Once you have connected to your database, you can perform CRUD (Create, Read, Update, Delete) operations using the Go driver for your NoSQL database.

### Create
To insert data into Redis, use the redis.Client.Set method:

```go
err = client.Set(ctx, "key", "value", 0).Err()
if err != nil {
	log.Fatal(err)
}
```

### Read
To retrieve data from Redis, use the redis.Client.Get method:

```go
value, err := client.Get(ctx, "key").Result()
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Value: %s\n", value)
```

### Update
To update data in Redis, you can use the same redis.Client.Set method as in the Create step:

```go
err = client.Set(ctx, "key", "new-value", 0).Err()
if err != nil {
	log.Fatal(err)
}
```

### Delete
To delete data from Redis, use the redis.Client.Del method:

```go
_, err = client.Del(ctx, "key").Result()
if err != nil {
	log.Fatal(err)
}
```

By understanding how to set up a NoSQL database, connect to it using a Go driver, and perform CRUD operations, you will be well-equipped to work with NoSQL databases in your Go applications.

# Day 14: Practice and Review

Today, you will practice and review the concepts learned during Week 2. Use this time to reinforce your understanding and identify any areas where you may need further study or improvement.

## Review the Concepts Learned During Week 2

- HTTP basics: Request methods, status codes, and headers
- RESTful APIs: Principles, design, and implementation
- JSON: Data format, parsing, and generation using Go's `encoding/json` package
- Databases in Go: SQL and NoSQL databases, database drivers, and CRUD operations

## Complete Exercises Related to HTTP, RESTful APIs, JSON, and Databases in Go

Find exercises, tutorials, or projects that involve the concepts you've learned this week. Some resources to consider include:

- [Go By Example](https://gobyexample.com/): Offers hands-on examples for various Go concepts
- [Gophercises](https://gophercises.com/): Coding exercises to practice your Go skills
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/): A book that covers various web development topics in Go, including HTTP, RESTful APIs, and working with databases

## Identify Areas for Further Study and Improvement

As you practice and review, make note of any concepts or skills you find challenging or want to learn more about. Use this list to guide your study in the coming weeks and to track your progress.

1. `Advanced Go language features`: Concurrency (goroutines and channels), interfaces, and error handling techniques.
2. `Generics`: Understand how to work with generics, which were introduced in Go 1.18.
3. `Middleware`: Learn about implementing and using middleware in Go web applications to handle common tasks like logging, authentication, and request validation.
4. `Testing`: Understand how to write and run tests in Go, including unit tests, integration tests, and benchmark tests.
5. `Websockets`: Learn how to use Websockets for real-time communication in your Go web applications.
6. `Authentication and authorization`: Study different authentication and authorization techniques, such as JWT, OAuth, and OpenID Connect.
7. `Performance optimization`: Learn how to optimize your Go applications for performance, including profiling, memory management, and best practices.
8. `Microservices`: Understand the microservices architecture and how to design and build microservices using Go.
9. `Deployment and continuous integration`: Learn about deploying Go applications to various platforms and setting up continuous integration and deployment pipelines.
10. `Working with various data storage solutions`: Explore different databases (e.g., SQL, NoSQL, and in-memory) and storage services (e.g., cloud storage, object storage) and learn how to work with them in Go.
11. `Third-party libraries and tools`: Familiarize yourself with popular third-party libraries and tools in the Go ecosystem to enhance your web development capabilities.
12. `Building and consuming APIs`: Learn how to build and consume RESTful and GraphQL APIs using Go.

By the end of Day 14, you should have a solid understanding of the concepts covered in Week 2 and be prepared to move on to Week 3, where you'll learn about Go web frameworks, tools, and advanced topics.
