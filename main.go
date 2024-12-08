package main

import (
	"ecomm-store/api"
	"ecomm-store/handlers"
	"ecomm-store/repo"
	"ecomm-store/service"
	"fmt"
	"net/http"
)

func main() {
	// Create an instance of the in-memory cart store
	store := repo.NewStore()

	cartService := service.CartService{Store: *store}

	// Create an instance of CartHandler with the store
	cartHandler := &handlers.CartHandler{
		Store:   store,
		CartSvc: cartService,
	}

	// Register the AddToCart endpoint
	http.HandleFunc(api.AddToCartEndpoint, cartHandler.AddToCart)
	http.HandleFunc(api.CheckoutEndpoint, cartHandler.Checkout)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
