package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// insertData inserts a single document into the specified MongoDB collection.
// Parameters:
// - collection: The MongoDB collection where the document will be inserted.
// - data: The document to be inserted into the collection. For example, bson.M{"name": "Alice", "age": 30}.
func insertData(collection *mongo.Collection, data interface{}) {
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("Error inserting data:", err)
	}
}

// deleteData deletes all documents that match the specified filter from the specified MongoDB collection.
// Parameters:
// - collection: The MongoDB collection where the documents will be deleted.
// - filter: The filter that determines which documents will be deleted. For example, bson.M{"name": "Alice"}.
func deleteData(collection *mongo.Collection, filter interface{}) {
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error deleting data:", err)
	}
}

// updateData updates all documents that match the specified filter in the specified MongoDB collection.
// Parameters:
// - collection: The MongoDB collection where the documents will be updated.
// - filter: The filter that determines which documents will be updated. For example, bson.M{"name": "Alice"}.
// - update: The update that will be applied to the documents. For example, bson.M{"$set": bson.M{"age": 31}}.
func updateData(collection *mongo.Collection, filter interface{}, update interface{}) {
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Error updating data:", err)
	}
}

// queryData queries all documents that match the specified filter from the specified MongoDB collection.
// Parameters:
// - collection: The MongoDB collection where the documents will be queried.
// - filter: The filter that determines which documents will be queried. For example, bson.M{"name": "Alice"}.
func queryData(collection *mongo.Collection, filter interface{}) {
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error querying data:", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		// To decode into a struct, use cursor.Decode()
	}
}

func main() {
	// Initialize parameters
	url := "mongodb://localhost:27017"
	databaseName := "your_database"
	collectionName := "your_collection"

	// Set client options
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()

	// Check the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database(databaseName).Collection(collectionName)

	// Example data
	document := map[string]interface{}{"name": "John", "age": 30}
	criteria := map[string]interface{}{"name": "John"}
	update := map[string]interface{}{"$set": map[string]interface{}{"age": 31}}

	// Insert data
	insertData(collection, document)

	// Query data
	queryData(collection, criteria)

	// Update data
	updateData(collection, criteria, update)

	// Delete data
	deleteData(collection, criteria)
}
