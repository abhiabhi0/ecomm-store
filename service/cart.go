package service

import (
	"ecomm-store/models"
	"ecomm-store/repo"
	"fmt"
)

type CartService struct {
	Store        repo.Store
	AdminService *AdminService
}

func (s *CartService) Checkout(userID int) (*models.CheckoutResponse, error) {
	// Get the user's data
	user, exists := s.Store.Users[userID]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	// Get the user's cart
	cart := s.Store.GetCart(userID)
	if len(cart.Items) == 0 {
		return nil, fmt.Errorf("no items in cart")
	}

	// Increase total orders
	user.TotalOrders++

	// Calculate total amount
	totalAmount := 0.0
	for _, item := range cart.Items {
		totalAmount += item.Price
	}

	couponCode := ""
	if user.TotalOrders%3 == 0 && user.Coupons != nil && len(user.Coupons) > 0 {
		couponCode = user.Coupons[0]
	}

	// Apply discount if coupon is available
	amountAfterDiscount := totalAmount
	if couponCode != "" && !user.DiscountUsed {
		amountAfterDiscount = totalAmount * 0.90 // Apply 10% discount
		s.Store.Users[userID] = models.User{
			ID:           user.ID,
			TotalOrders:  user.TotalOrders,
			DiscountUsed: true, // Mark discount as used
		}
	}

	// Update user order count in store
	s.Store.Users[userID] = user

	// Clear the cart after checkout
	s.Store.ClearCart(userID)

	checkoutResponse := &models.CheckoutResponse{
		UserID:              userID,
		TotalAmount:         totalAmount,
		CouponCode:          couponCode,
		OrderNumber:         user.TotalOrders,
		AmountAfterDiscount: amountAfterDiscount,
	}

	return checkoutResponse, nil
}
