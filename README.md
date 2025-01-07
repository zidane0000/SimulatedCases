# Simulated Cases

This repository contains various simulated cases for learning and testing purposes.

## Cases

### 1. Go Channel

This example demonstrates how to use channels for communication between goroutines in Go.

- **File:** `Gochannel.go`
- **Description:** This showcases the basic usage of channels in Go, including sending and receiving data between goroutines. The code includes functions to handle data streaming, buffering, and flushing, as well as a function to simulate sending random data through a channel. The `main` function initializes the channel and starts the goroutines for data sending and streaming.


### 2. PostgreDB

This example demonstrates how to query and prioritize data from a PostgreSQL database using Go.

- **File:** `PostgreDB.go`
- **Description:** This code connects to a PostgreSQL database, queries data from two tables (`IMAGES` and `POSTS`), and prioritizes the data based on the creation time using a priority queue. The `main` function simulates the database connection, queries the data, combines and prioritizes it, and then prints the prioritized items.

### 3. MongoDB

This example demonstrates how to perform basic operations (insert, query, update, delete) on MongoDB using Go.

- **File:** `mongoDB.go`
- **Description:** This code connects to a MongoDB database, performs basic CRUD operations, and prints the results. The `main` function initializes the MongoDB client, connects to the database, and executes the insert, query, update, and delete operations on a sample document. The example uses a local MongoDB instance running on `mongodb://localhost:27017`.
