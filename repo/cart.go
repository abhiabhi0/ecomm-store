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
	GetUser(userID int) (models.User, bool)
	SetUser(userID int, user models.User)
	UpdateOrder(userID int, order models.Order)
	GetAllUsers() []models.User
	GetOrder(userID int) []models.Order
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

func (store *Store) GetUser(userID int) (models.User, bool) {
	if _, exists := store.Users[userID]; !exists {
		return models.User{}, false
	}
	return store.Users[userID], true
}

func (store *Store) SetUser(userID int, user models.User) {
	store.Users[userID] = user
}

func (store *Store) UpdateOrder(userID int, order models.Order) {
	store.Orders[userID] = append(store.Orders[userID], order)
}

func (store *Store) GetAllUsers() []models.User {
	users := make([]models.User, 0, len(store.Users))
	for _, user := range store.Users {
		users = append(users, user)
	}
	return users
}

func (store *Store) GetOrder(userID int) []models.Order {
	return store.Orders[userID]
}
