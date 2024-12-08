package repo

import (
	"ecomm-store/models"
	"fmt"
)

// CartStore interface defines the operations for the cart
type CartStore interface {
	AddItemToCart(userID int, item models.Item) models.Cart
	GetCart(userID int) models.Cart
	ClearCart(userID int)
}

// Store holds the cart and order data
type Store struct {
	Carts  map[int]models.Cart    // User's cart by userID
	Users  map[int]models.User    // Users' information
	Orders map[int][]models.Order // List of orders for each user, keyed by userID
}

func NewStore() *Store {
	return &Store{
		Carts:  make(map[int]models.Cart),
		Users:  make(map[int]models.User),
		Orders: make(map[int][]models.Order),
	}
}

// AddItemToCart adds an item to the user's cart
func (store *Store) AddItemToCart(userID int, item models.Item) models.Cart {
	if _, exists := store.Users[userID]; !exists {
		store.Users[userID] = models.User{ID: userID, TotalOrders: 0, DiscountUsed: false}
	}

	if _, exists := store.Carts[userID]; !exists {
		store.Carts[userID] = models.Cart{}
	}
	cart := store.Carts[userID]
	cart.Items = append(cart.Items, item)
	store.Carts[userID] = cart
	fmt.Printf("Updated cart for user %d: %v\n", userID, store.Carts[userID]) //TODO remove in the end print statements
	return store.Carts[userID]
}

// GetCart retrieves the user's cart
func (store *Store) GetCart(userID int) models.Cart {
	return store.Carts[userID]
}

// ClearCart clears the user's cart
func (store *Store) ClearCart(userID int) {
	store.Carts[userID] = models.Cart{}
}
