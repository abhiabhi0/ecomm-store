package service_test

import (
	"ecomm-store/models"
	"ecomm-store/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Unit test for Checkout
func TestCheckout(t *testing.T) {
	mockStore := new(MockCartStore)

	// Mock data for the test
	userID := 1
	item := models.Item{ID: 1, Name: "Product 1", Price: 100.0}
	user := models.User{
		ID:           userID,
		TotalOrders:  2,
		DiscountUsed: false,
		Coupons:      []string{"DISCOUNT10"},
	}

	cart := models.Cart{
		Items: []models.Item{item},
	}

	// Setting expectations for the mock methods
	mockStore.On("GetUser", userID).Return(user, true)
	mockStore.On("GetCart", userID).Return(cart)
	mockStore.On("SetUser", userID, mock.Anything).Return()
	mockStore.On("UpdateOrder", userID, mock.Anything).Return()
	mockStore.On("ClearCart", userID).Return()

	// Create the CartService instance
	cartService := &service.CartService{Store: mockStore}

	// Call Checkout
	checkoutResponse, err := cartService.Checkout(userID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, checkoutResponse)
	assert.Equal(t, 3, checkoutResponse.OrderNumber)
	assert.Equal(t, 100.0, checkoutResponse.TotalAmount)
	assert.Equal(t, "DISCOUNT10", checkoutResponse.CouponCode)
	assert.Equal(t, 90.0, checkoutResponse.AmountAfterDiscount)

	// Check that the mock methods were called as expected
	mockStore.AssertExpectations(t)
}

func TestCheckout_NoUser(t *testing.T) {
	mockStore := new(MockCartStore)

	userID := 1

	// Setting expectations for the mock methods
	mockStore.On("GetUser", userID).Return(models.User{}, false)

	// Create the CartService instance
	cartService := &service.CartService{Store: mockStore}

	// Call Checkout
	checkoutResponse, err := cartService.Checkout(userID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, checkoutResponse)
	assert.Equal(t, "user not found", err.Error())

	// Check that the mock methods were called as expected
	mockStore.AssertExpectations(t)
}

func TestCheckout_EmptyCart(t *testing.T) {
	mockStore := new(MockCartStore)

	userID := 1
	user := models.User{
		ID:           userID,
		TotalOrders:  1,
		DiscountUsed: false,
		Coupons:      []string{"DISCOUNT10"},
	}

	// Setting expectations for the mock methods
	mockStore.On("GetUser", userID).Return(user, true)
	mockStore.On("GetCart", userID).Return(models.Cart{})

	// Create the CartService instance
	cartService := &service.CartService{Store: mockStore}

	// Call Checkout
	checkoutResponse, err := cartService.Checkout(userID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, checkoutResponse)
	assert.Equal(t, "no items in cart", err.Error())

	// Check that the mock methods were called as expected
	mockStore.AssertExpectations(t)
}
