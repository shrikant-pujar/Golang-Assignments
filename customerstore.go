package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string `json:"UID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

var customerstore []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Function Called: homePage()")
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getCustomerstore()")
	json.NewEncoder(w).Encode(customerstore)
}
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item

	_ = json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON

	customerstore = append(customerstore, item) // Add item to customerstore
	json.NewEncoder(w).Encode(item)             // Show item in response JSON for verification
}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_deleteItemAtUid(params["uid"])
	json.NewEncoder(w).Encode(customerstore)
}
func _deleteItemAtUid(uid string) {
	for index, item := range customerstore {
		if item.UID == uid {
			// Delete item from Slice
			customerstore = append(customerstore[:index], customerstore[index+1:]...)
			break
		}
	}
}
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params := mux.Vars(r)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])             // Delete item
	customerstore = append(customerstore, item) // Create it again with data from request

	json.NewEncoder(w).Encode(customerstore)
}
func handleRequests() {
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/home", homePage).Methods("GET")
	router.HandleFunc("/customerstore", getCustomer).Methods("GET")
	router.HandleFunc("/customerstore/{uid}", updateItem).Methods("GET")
	router.HandleFunc("/customerstore/{uid}", updateItem).Methods("PUT")
	router.HandleFunc("/customerstore/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/customerstore", createItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	// Data store
	customerstore = append(customerstore, Item{
		UID:   "1",
		Name:  "Rajkumar",
		Email: "raj@gmail.com",
	})
	customerstore = append(customerstore, Item{
		UID:   "2",
		Name:  "Shrikant",
		Email: "shri@gmail.com",
	})
	handleRequests()
}
