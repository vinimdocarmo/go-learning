package main

// Sale represents a sale of products
type Sale struct {
	date          string
	time          string
	saleLineItems []SaleLineItem
}

// GetTotal returns the sum of all the products of all the linds in sale
func (s *Sale) GetTotal() float64 {
	total := float64(0)

	for _, sl := range s.saleLineItems {
		total += sl.GetSubTotal()
	}

	return total
}

// AddLineItem adds a new line item into a slice of lines
func (s *Sale) AddLineItem(quantity int, p Product) {
	sl := SaleLineItem{quantity: quantity, product: p}
	s.saleLineItems = append(s.saleLineItems, sl)
}

// NewSale returns a new instance of Sale
func NewSale() Sale {
	return Sale{saleLineItems: []SaleLineItem{}}
}
