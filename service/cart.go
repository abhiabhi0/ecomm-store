package service

import (
	"ecomm-store/models"
	"ecomm-store/repo"
	"fmt"
)

type CartService struct {
	Store repo.CartStore
}

func (s *CartService) Checkout(userID int) (*models.CheckoutResponse, error) {
	// Get the user's data
	user, exists := s.Store.GetUser(userID)
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	// Get the user's cart
	cart := s.Store.GetCart(userID)
	if len(cart.Items) == 0 {
		return nil, fmt.Errorf("no items in cart")
	}

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
		user.DiscountUsed = true
	}

	// Update user order count in store
	s.Store.SetUser(userID, user)

	// Create the order struct
	order := models.Order{
		UserID: userID,
		Items:  cart.Items,
		Total:  totalAmount,
		Coupon: couponCode,
	}

	// Add the order to the user's order history
	s.Store.UpdateOrder(userID, order)

	// Clear the cart after checkout
	s.Store.ClearCart(userID)

	checkoutResponse := &models.CheckoutResponse{
		UserID:              userID,
		TotalAmount:         totalAmount,
		OrderNumber:         user.TotalOrders,
		AmountAfterDiscount: amountAfterDiscount,
		CouponCode:          couponCode,
	}

	return checkoutResponse, nil
}
