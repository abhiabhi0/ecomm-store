package main

import (
	"ecomm-store/api"
	"ecomm-store/handlers"
	"ecomm-store/middleware"
	"ecomm-store/repo"
	"ecomm-store/service"
	"fmt"
	"net/http"
)

func main() {
	// Create an instance of the in-memory cart store
	store := repo.NewStore()

	cartService := &service.CartService{Store: store}

	adminService := &service.AdminService{Store: store}

	// Create an instance of CartHandler with the store
	cartHandler := &handlers.CartHandler{
		CartSvc: cartService,
	}

	adminHandler := &handlers.AdminHandler{AdminService: adminService}

	// Register the AddToCart endpoint
	http.HandleFunc(api.AddToCartEndpoint, cartHandler.AddToCart)
	http.HandleFunc(api.CheckoutEndpoint, cartHandler.Checkout)

	http.Handle(api.GenerateDiscountCodeEndpoint, middleware.AdminAuthMiddleware(http.HandlerFunc(adminHandler.GenerateCoupon)))
	http.Handle(api.AdminReportEndpoint, middleware.AdminAuthMiddleware(http.HandlerFunc(adminHandler.GetAdminReport)))

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
