package service_test

import (
	"ecomm-store/models"
	"ecomm-store/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Unit test for GenerateDiscountCode
func TestGenerateDiscountCode(t *testing.T) {
	mockStore := new(MockCartStore)

	// Mock data
	userID := 1
	user := models.User{
		ID:           userID,
		TotalOrders:  2,
		DiscountUsed: false,
		Coupons:      []string{"DISCOUNT10"},
	}

	// Setting expectations for the mock methods
	mockStore.On("GetUser", userID).Return(user, true)
	mockStore.On("SetUser", userID, mock.Anything).Return()

	// Create AdminService instance
	adminService := &service.AdminService{Store: mockStore}

	// Call GenerateDiscountCode
	err := adminService.GenerateDiscountCode(userID)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, user.Coupons, 1)
	assert.Equal(t, "DISCOUNT10", user.Coupons[0])

	// Check that the mock methods were called as expected
	mockStore.AssertExpectations(t)
}

func TestGenerateDiscountCode_UserNotFound(t *testing.T) {
	mockStore := new(MockCartStore)

	userID := 1

	// Setting expectations for the mock methods
	mockStore.On("GetUser", userID).Return(models.User{}, false)

	// Create AdminService instance
	adminService := &service.AdminService{Store: mockStore}

	// Call GenerateDiscountCode
	err := adminService.GenerateDiscountCode(userID)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// Check that the mock methods were called as expected
	mockStore.AssertExpectations(t)
}

// Unit test for GetAdminReport
func TestGetAdminReport(t *testing.T) {
	mockStore := new(MockCartStore)

	// Mock data for users and orders
	user1 := models.User{
		ID:           1,
		TotalOrders:  2,
		DiscountUsed: false,
		Coupons:      []string{"DISCOUNT10"},
	}
	user2 := models.User{
		ID:           2,
		TotalOrders:  1,
		DiscountUsed: false,
		Coupons:      []string{},
	}

	order1 := models.Order{
		UserID: 1,
		Items:  []models.Item{{ID: 1, Name: "Item1", Price: 50.0}},
		Total:  50.0,
		Coupon: "DISCOUNT10",
	}

	order2 := models.Order{
		UserID: 2,
		Items:  []models.Item{{ID: 2, Name: "Item2", Price: 75.0}},
		Total:  75.0,
		Coupon: "",
	}

	// Setting expectations for the mock methods
	mockStore.On("GetAllUsers").Return([]models.User{user1, user2})
	mockStore.On("GetOrder", 1).Return([]models.Order{order1})
	mockStore.On("GetOrder", 2).Return([]models.Order{order2})

	// Create AdminService instance
	adminService := &service.AdminService{Store: mockStore}

	// Call GetAdminReport
	report, err := adminService.GetAdminReport()

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, report, 2)

	// Check user 1 report
	assert.Equal(t, 1, report[1].TotalItems)
	assert.Equal(t, 50.0, report[1].TotalAmount)
	assert.Contains(t, report[1].DiscountCodes, "DISCOUNT10")
	assert.Equal(t, 5.0, report[1].TotalDiscount)

	// Check user 2 report
	assert.Equal(t, 1, report[2].TotalItems)
	assert.Equal(t, 75.0, report[2].TotalAmount)
	assert.Empty(t, report[2].DiscountCodes)
	assert.Equal(t, 0.0, report[2].TotalDiscount)

	// Check that the mock methods were called as expected
	mockStore.AssertExpectations(t)
}
