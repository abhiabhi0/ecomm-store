package models

// Item represents an item in the cart
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Cart represents the cart for a user
type Cart struct {
	Items []Item `json:"items"`
}

// User represents a customer/user of the ecommerce store
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	TotalOrders  int    `json:"total_orders"`
	DiscountUsed bool   `json:"discount_used"`
}

// Order represents an order placed by a user
type Order struct {
	UserID int     `json:"user_id"`
	Items  []Item  `json:"items"`
	Total  float64 `json:"total"`
	Coupon string  `json:"coupon,omitempty"`
}

type CheckoutResponse struct {
	UserID              int     `json:"user_id"`
	TotalAmount         float64 `json:"total_amount"`
	CouponCode          string  `json:"coupon_code,omitempty"`
	OrderNumber         int     `json:"order_number"`
	AmountAfterDiscount float64 `json:"amount_after_discount"`
}
