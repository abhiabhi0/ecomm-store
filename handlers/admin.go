package handlers

import (
	"ecomm-store/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	AdminService *service.AdminService
}

// GenerateCoupon generates a discount code for a user
func (h *AdminHandler) GenerateCoupon(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL query parameters
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Generate coupon using AdminService
	err = h.AdminService.GenerateDiscountCode(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate coupon: %v", err), http.StatusBadRequest)
		return
	}

	// Respond with success message
	json.NewEncoder(w).Encode(map[string]string{"message": "Coupon added successfully"})
}

// GetAdminReport generates a report for admin
func (h *AdminHandler) GetAdminReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.AdminService.GetAdminReport()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate report: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the report
	json.NewEncoder(w).Encode(report)
}
