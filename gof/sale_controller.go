package main

import (
	"errors"
	"fmt"
)

// SaleController is the interface of communication between the user and the sale system
type SaleController struct {
	sale           Sale
	saleEnded      bool
	productCatalog []Product
}

// NewSaleController creates a new controller
func NewSaleController() SaleController {
	return SaleController{sale: Sale{}, saleEnded: false, productCatalog: []Product{
		Product{id: "001", price: 35000, description: "Fiat Palio"},
		Product{id: "002", price: 1500, description: "Moto X"},
	}}
}

// StartSale starts a new sale initializing some elements
func (c *SaleController) StartSale() {
	c.sale = NewSale()
}

// EnterItem adds a new line item into the sale object
func (c *SaleController) EnterItem(quantity int, id string) error {
	product, err := c.GetProductByID(id)

	if err != nil {
		return err
	}

	c.sale.AddLineItem(quantity, product)

	return nil
}

// GetProductByID in this case is functining as a DAO
func (c *SaleController) GetProductByID(id string) (Product, error) {
	for _, product := range c.productCatalog {
		if product.id == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Product not found")
}

// PrintSaleInfo prints into stardard output informations about current sale
func (c *SaleController) PrintSaleInfo() {
	for i, sl := range c.sale.saleLineItems {
		fmt.Printf("%v.\n", i)
		fmt.Printf("Product:\t %v\n", sl.product.description)
		fmt.Printf("Price:\t\t %v\n\n", sl.product.GetPrice())
	}
	fmt.Printf("Total price of sale: R$%v\n", c.sale.GetTotal())
}

// EndSale ends the process of a sale
func (c *SaleController) EndSale() *Sale {
	return &c.sale
}
