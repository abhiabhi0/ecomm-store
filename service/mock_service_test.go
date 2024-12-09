package service_test

import (
	"ecomm-store/models"

	"github.com/stretchr/testify/mock"
)

// Mock for the CartStore
type MockCartStore struct {
	mock.Mock
}

func (m *MockCartStore) AddItemToCart(userID int, item models.Item) models.Cart {
	args := m.Called(userID, item)
	return args.Get(0).(models.Cart)
}

func (m *MockCartStore) GetCart(userID int) models.Cart {
	args := m.Called(userID)
	return args.Get(0).(models.Cart)
}

func (m *MockCartStore) ClearCart(userID int) {
	m.Called(userID)
}

func (m *MockCartStore) GetUser(userID int) (models.User, bool) {
	args := m.Called(userID)
	return args.Get(0).(models.User), args.Bool(1)
}

func (m *MockCartStore) SetUser(userID int, user models.User) {
	m.Called(userID, user)
}

func (m *MockCartStore) UpdateOrder(userID int, order models.Order) {
	m.Called(userID, order)
}

func (m *MockCartStore) GetAllUsers() []models.User {
	args := m.Called()
	return args.Get(0).([]models.User)
}

func (m *MockCartStore) GetOrder(userID int) []models.Order {
	args := m.Called(userID)
	return args.Get(0).([]models.Order)
}
