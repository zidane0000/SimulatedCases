package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var resources = make(map[string]Resource)

func main() {
	http.HandleFunc("/resource", resourceHandler)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// resourceHandler routes the HTTP request to the appropriate function based on the method.
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request to handle.
func resourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getResource(w, r)
	case http.MethodPost:
		createResource(w, r)
	case http.MethodPut:
		updateResource(w, r)
	case http.MethodDelete:
		deleteResource(w, r)
	case http.MethodHead:
		headResource(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getResource retrieves a resource by ID and writes it to the response.
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request to handle.
func getResource(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if resource, ok := resources[id]; ok {
		json.NewEncoder(w).Encode(resource)
	} else {
		http.Error(w, "Resource not found", http.StatusNotFound)
	}
}

// createResource decodes a JSON payload from the request body and adds it to the resources map.
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request to handle.
func createResource(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	resources[resource.ID] = resource
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

// updateResource updates an existing resource identified by ID with new data from the request body.
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request to handle.
func updateResource(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var resource Resource
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if _, ok := resources[id]; ok {
		resources[id] = resource
		json.NewEncoder(w).Encode(resource)
	} else {
		http.Error(w, "Resource not found", http.StatusNotFound)
	}
}

// deleteResource removes a resource identified by ID from the resources map.
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request to handle.
func deleteResource(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if _, ok := resources[id]; ok {
		delete(resources, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Resource not found", http.StatusNotFound)
	}
}

// headResource handles the HTTP HEAD request by returning headers without a body.
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request to handle.
func headResource(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if resource, ok := resources[id]; ok {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Resource-Name", resource.Name)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Resource not found", http.StatusNotFound)
	}
}
