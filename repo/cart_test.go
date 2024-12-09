package repo_test

import (
	"ecomm-store/models"
	"ecomm-store/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	store := repo.NewStore()

	// Test AddItemToCart
	t.Run("AddItemToCart", func(t *testing.T) {
		item := models.Item{ID: 1, Name: "Laptop", Price: 1000.0}
		cart := store.AddItemToCart(1, item)
		assert.Equal(t, 1, len(cart.Items))
		assert.Equal(t, item, cart.Items[0])
	})

	// Test GetCart
	t.Run("GetCart", func(t *testing.T) {
		cart := store.GetCart(1)
		assert.Equal(t, 1, len(cart.Items))
		assert.Equal(t, "Laptop", cart.Items[0].Name)
	})

	// Test ClearCart
	t.Run("ClearCart", func(t *testing.T) {
		store.ClearCart(1)
		cart := store.GetCart(1)
		assert.Equal(t, 0, len(cart.Items))
	})

	// Test GetUser and SetUser
	t.Run("GetUser and SetUser", func(t *testing.T) {
		user := models.User{ID: 1, Name: "John Doe", TotalOrders: 5, DiscountUsed: true}
		store.SetUser(1, user)
		retrievedUser, exists := store.GetUser(1)
		assert.True(t, exists)
		assert.Equal(t, user, retrievedUser)
	})

	// Test UpdateOrder
	t.Run("UpdateOrder", func(t *testing.T) {
		order := models.Order{
			UserID: 1,
			Items:  []models.Item{{ID: 1, Name: "Laptop", Price: 1000.0}},
			Total:  1000.0,
			Coupon: "DISCOUNT10",
		}
		store.UpdateOrder(1, order)
		orders := store.GetOrder(1)
		assert.Equal(t, 1, len(orders))
		assert.Equal(t, order, orders[0])
	})

	// Test GetAllUsers
	t.Run("GetAllUsers", func(t *testing.T) {
		users := store.GetAllUsers()
		assert.Equal(t, 1, len(users))
		assert.Equal(t, "John Doe", users[0].Name)
	})

	// Test GetOrder
	t.Run("GetOrder", func(t *testing.T) {
		orders := store.GetOrder(1)
		assert.Equal(t, 1, len(orders))
		assert.Equal(t, "DISCOUNT10", orders[0].Coupon)
	})
}
