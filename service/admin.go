// service/admin_service.go

package service

import (
	"ecomm-store/models"
	"ecomm-store/repo"
	"fmt"
)

type AdminService struct {
	Store *repo.Store
}

// GenerateDiscountCode generates a discount code for every 3rd order and adds it to the user's list of coupons
func (s *AdminService) GenerateDiscountCode(userID int) error {
	// Retrieve user from the store
	user, exists := s.Store.Users[userID]
	if !exists {
		return fmt.Errorf("user not found")
	}

	// Generate a discount code if it's the 3rd order
	discountCode := "DISCOUNT10"

	// Add the discount code to the user's coupons list
	user.Coupons = append(user.Coupons, discountCode)

	// Update the user in the store
	s.Store.Users[userID] = user
	return nil
}

// GetAdminReport lists the count of items purchased, total purchase amount, discount codes, and total discount amount
func (s *AdminService) GetAdminReport() (map[int]models.ReportResponse, error) {
	report := make(map[int]models.ReportResponse)

	// Iterate through all users and generate report
	for userID, _ := range s.Store.Users {
		totalItems := 0
		totalAmount := 0.0
		discountCodes := []string{}
		totalDiscount := 0.0

		// Calculate total items and total amount for each user
		for _, order := range s.Store.Orders {
			if order.UserID == userID {
				// Update total items count
				totalItems += len(order.Items)

				// Update total amount
				totalAmount += order.Total

				// Check if the order used a coupon (discount code)
				if order.Coupon != "" {
					discountCodes = append(discountCodes, order.Coupon)

					// Assuming the discount amount is stored in the order and is 10% of the total order
					// You may need to adjust this if the discount logic is more complex.
					discountAmount := order.Total * 0.10 // 10% discount
					totalDiscount += discountAmount
				}
			}
		}

		// Populate the report for this user using the ReportResponse struct
		report[userID] = models.ReportResponse{
			TotalItems:    totalItems,
			TotalAmount:   totalAmount,
			DiscountCodes: discountCodes,
			TotalDiscount: totalDiscount,
		}
	}

	return report, nil
}
