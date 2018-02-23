package main

// SaleLineItem represents a sale line item
type SaleLineItem struct {
	quantity int
	product  Product
}

// GetSubTotal returns the total price of all the products on the line
func (sl *SaleLineItem) GetSubTotal() float64 {
	return float64(sl.quantity) * sl.product.GetPrice()
}
