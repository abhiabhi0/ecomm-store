package handlers

import (
	"ecomm-store/models"
	"ecomm-store/service"
	"encoding/json"
	"net/http"
	"strconv"
)

// CartHandler handles cart-related requests
type CartHandler struct {
	CartSvc *service.CartService
}

// AddToCart handles adding an item to the user's cart
func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add item to cart for the user
	updatedCart := h.CartSvc.Store.AddItemToCart(userID, item)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCart)
}

// Checkout handles the checkout process for the user
func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service to perform checkout
	checkoutResponse, err := h.CartSvc.Checkout(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkoutResponse)
}
