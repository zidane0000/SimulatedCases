# RESTful API Example

This example demonstrates how to create and execute a basic RESTful API using Go.

## RESTful API Overview

A RESTful API is an architectural style for designing networked applications. It relies on a stateless, client-server, cacheable communications protocol -- the HTTP protocol. Here are some key principles and components of RESTful API design:

1. **Resources**: The key abstraction of information in REST is a resource. Resources are identified by URIs. In this example, the resource is represented by the `Resource` struct with `ID` and `Name` fields.

2. **HTTP Methods**: RESTful APIs use standard HTTP methods to perform operations on resources:
   - **GET**: Retrieve a resource.
   - **POST**: Create a new resource.
   - **PUT**: Update an existing resource.
   - **DELETE**: Remove a resource.
   - **HEAD**: Retrieve metadata about a resource without the body.

3. **Statelessness**: Each request from a client contains all the information needed to process the request. The server does not store any state about the client session.

4. **Representations**: Resources can have different representations, such as JSON or XML. In this example, JSON is used for resource representation.

5. **Status Codes**: Use standard HTTP status codes to indicate the result of an operation (e.g., `200 OK`, `201 Created`, `404 Not Found`).

## Installation

Before you start, ensure that Go is installed on your system. Then, follow these steps to set up the project:

1. Initialize a Go module (if not already initialized):
    ```sh
    go mod init restfulApi
    ```

## Usage

1. Ensure Go is installed on your local machine. You can download it from [golang.org](https://golang.org/).

2. Run the Go program:
    - Save the provided code into a file named `restfulApi.go`.
    - Execute the program using the following command:
      ```sh
      go run restfulApi.go
      ```

3. Test the API:
    - Use a tool like `curl` or Postman to interact with the API.
    - Example `curl` commands:
      - **GET** a resource:
        ```sh
        curl -i -X GET "http://localhost:8080/resource?id=1"
        ```
      - **POST** a new resource:
        ```sh
        curl -i -X POST -H "Content-Type: application/json" -d '{"id":"1","name":"Sample"}' "http://localhost:8080/resource"
        ```
      - **PUT** to update a resource:
        ```sh
        curl -i -X PUT -H "Content-Type: application/json" -d '{"id":"1","name":"Updated Sample"}' "http://localhost:8080/resource?id=1"
        ```
      - **DELETE** a resource:
        ```sh
        curl -i -X DELETE "http://localhost:8080/resource?id=1"
        ```
      - **HEAD** to check a resource:
        ```sh
        curl -i -I "http://localhost:8080/resource?id=1"
        ```

This document should help you understand the basic execution of a RESTful API using Go.