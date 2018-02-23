package main

func main() {

	c := NewSaleController()
	c.StartSale()
	c.EnterItem(1, "001")
	c.EnterItem(2, "002")
	c.EndSale()

	c.PrintSaleInfo()
}
