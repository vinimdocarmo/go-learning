package main

import (
	"errors"
	"fmt"
)

// PRODUCT

type Product struct {
	id          string
	description string
	price       float64
}

func (p Product) GetPrice() float64 {
	return p.price
}

// SALE LINE ITEM

type SaleLineItem struct {
	quantity int
	product  Product
}

func (sl *SaleLineItem) GetSubTotal() float64 {
	return float64(sl.quantity) * sl.product.GetPrice()
}

// SALE

type Sale struct {
	date          string
	time          string
	saleLineItems []SaleLineItem
}

func (s *Sale) GetTotal() float64 {
	total := float64(0)

	for _, sl := range s.saleLineItems {
		total += sl.GetSubTotal()
	}

	return total
}

func (s *Sale) AddLineItem(quantity int, p Product) {
	sl := SaleLineItem{quantity: quantity, product: p}
	s.saleLineItems = append(s.saleLineItems, sl)
}

func NewSale() Sale {
	return Sale{saleLineItems: []SaleLineItem{}}
}

// REGISTER

type SaleController struct {
	sale           Sale
	saleEnded      bool
	productCatalog []Product
}

func (c *SaleController) StartSale() {
	c.saleEnded = false
	c.sale = NewSale()
	c.productCatalog = []Product{
		Product{id: "001", price: 35000, description: "Fiat Palio"},
		Product{id: "002", price: 1500, description: "Moto X"},
	}
}

func (c *SaleController) EnterItem(quantity int, id string) error {
	product, err := c.GetProductById(id)

	if err != nil {
		return err
	}

	c.sale.AddLineItem(quantity, product)

	return nil
}

func (c *SaleController) GetProductById(id string) (Product, error) {
	for _, product := range c.productCatalog {
		if product.id == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Product not found")
}

func (c *SaleController) PrintSaleInfo() {
	for i, sl := range c.sale.saleLineItems {
		fmt.Printf("%v.\n", i)
		fmt.Printf("Product:\t %v\n", sl.product.description)
		fmt.Printf("Price:\t\t %v\n\n", sl.product.GetPrice())
	}
	fmt.Printf("Total price of sale: R$%v\n", c.sale.GetTotal())
}

func (c *SaleController) EndSale() *Sale {
	return &c.sale
}

func main() {

	c := SaleController{}
	c.StartSale()
	c.EnterItem(1, "001")
	c.EnterItem(2, "002")
	c.EndSale()

	c.PrintSaleInfo()
}
