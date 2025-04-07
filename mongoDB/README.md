# MongoDB Example

This example demonstrates how to perform basic operations (insert, query, update, delete) on MongoDB using Go.

## Installation

Before you start, make sure you have Go installed on your system. Then, follow these steps to install the required Go packages:

1. Initialize a Go module (if not already initialized):

    ```sh
    go mod init your_module_name
    ```

2. Install the MongoDB Go driver:

    ```sh
    go get go.mongodb.org/mongo-driver/mongo
    ```

## Usage

1. Ensure MongoDB is installed and running on your local machine. You can check if MongoDB is running by executing the following command in your terminal:

    ```sh
    mongo --eval "db.runCommand({ connectionStatus: 1 })"
    ```

   If MongoDB is running, this command will return the connection status. If not, start MongoDB using:

    ```sh
    mongod
    ```

2. Run the Go program:

    ```sh
    go run mongoDB.go
    ```

This will connect to MongoDB, insert a document, query it, update it, and finally delete it.

## Example Code

The example code provided in `mongoDB.go` performs the following operations:

- **Insert Data**: Inserts a single document into the specified MongoDB collection.
- **Query Data**: Queries all documents that match the specified filter from the specified MongoDB collection.
- **Update Data**: Updates all documents that match the specified filter in the specified MongoDB collection.
- **Delete Data**: Deletes all documents that match the specified filter from the specified MongoDB collection.

Make sure to replace `"your_database"` and `"your_collection"` with the actual database and collection names you are working with.

## Notes

- The example uses a local MongoDB instance running on `mongodb://localhost:27017`. Update the URI if your MongoDB instance is hosted elsewhere.
- The example data includes a document with `{"name": "John", "age": 30}`. Adjust the data as needed for your use case.
