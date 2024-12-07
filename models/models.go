package models

type Cart struct {
	Items []Item
}

type Item struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}
