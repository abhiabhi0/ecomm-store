package service

import (
	"ecomm-store/models"
	"ecomm-store/repo"
	"fmt"
)

// AdminService provides administrative functionalities.
type AdminService struct {
	Store repo.CartStore
}

// GenerateDiscountCode generates a discount code for every 3rd order and adds it to the user's list of coupons.
func (s *AdminService) GenerateDiscountCode(userID int) error {
	// Retrieve user from the store
	user, exists := s.Store.GetUser(userID)
	if !exists {
		return fmt.Errorf("user not found")
	}

	// Generate a discount code
	discountCode := "DISCOUNT10"

	// Add the discount code to the user's coupons list
	user.Coupons = append(user.Coupons, discountCode)

	// Update the user in the store
	s.Store.SetUser(userID, user)
	return nil
}

// GetAdminReport generates a report containing user order statistics and discounts used.
func (s *AdminService) GetAdminReport() (map[int]models.ReportResponse, error) {
	report := make(map[int]models.ReportResponse)

	allUsers := s.Store.GetAllUsers()
	fmt.Printf("All users: %v\n", allUsers) //TODO remove in the end print statements)")
	// Iterate through all users and generate a report
	for _, user := range allUsers {
		userID := user.ID
		fmt.Printf("Generating report for user %d\n", userID) //TODO remove in the end print statements)
		totalItems := 0
		totalAmount := 0.0
		discountCodes := []string{}
		totalDiscount := 0.0

		// Calculate total items and total amount for each user
		for _, order := range s.Store.GetOrder(userID) {
			fmt.Printf("Processing order for user %d: %v\n", userID, order) //TODO remove in the end print statements)
			// Update total items count
			totalItems += len(order.Items)

			// Update total amount
			totalAmount += order.Total

			// Check if the order used a coupon (discount code)
			if order.Coupon != "" {
				discountCodes = append(discountCodes, order.Coupon)

				// Assuming the discount amount is stored in the order and is 10% of the total order
				discountAmount := order.Total * 0.10 // 10% discount
				totalDiscount += discountAmount
			}
		}

		// Populate the report for this user using the ReportResponse struct
		report[userID] = models.ReportResponse{
			TotalItems:    totalItems,
			TotalAmount:   totalAmount,
			DiscountCodes: discountCodes,
			TotalDiscount: totalDiscount,
		}

		fmt.Printf("Report for user %d: %v\n", userID, report[userID]) //TODO remove in the end print statements)
	}

	return report, nil
}
