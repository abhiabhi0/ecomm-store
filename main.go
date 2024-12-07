package main

import (
	"ecomm-store/handlers"
	"ecomm-store/repo"
	"fmt"
	"net/http"
)

func main() {
	// Create an instance of the in-memory cart store
	cartStore := &repo.Store{}

	// Create an instance of CartHandler with the store
	cartHandler := &handlers.CartHandler{
		Store: cartStore,
	}

	// Register the AddToCart endpoint
	http.HandleFunc("/v1/cart/add", cartHandler.AddToCart)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
