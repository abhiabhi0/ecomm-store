package handlers

import (
	"ecomm-store/models"
	"ecomm-store/repo"
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	Store repo.CartStore
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedCart := h.Store.AddItem(item)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCart)
}
