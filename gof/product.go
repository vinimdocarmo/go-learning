package main

// Product represents a product for sale
type Product struct {
	id          string
	description string
	price       float64
}

// GetPrice returns the price of the product
func (p Product) GetPrice() float64 {
	return p.price
}
