package repo

import (
	"ecomm-store/models"
	"fmt"
)

type CartStore interface {
	AddItem(item models.Item) models.Cart
	GetCart() models.Cart
	ClearCart()
}

type Store struct {
	CartData models.Cart
}

func (store *Store) AddItem(item models.Item) models.Cart {
	store.CartData.Items = append(store.CartData.Items, item)
	fmt.Printf("updated cart: %v\n", store.CartData) //TODO remove in the end print stmts
	return store.CartData
}

func (store *Store) GetCart() models.Cart {
	return store.CartData
}

func (store *Store) ClearCart() {
	store.CartData.Items = []models.Item{}
}
