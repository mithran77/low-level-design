package main

type Inventory struct {
	products map[*Product]int
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[*Product]int)}
}

func (inv *Inventory) AddProduct(p *Product, quantity int) {
	inv.products[p] = quantity
}

func (inv *Inventory) IsAvailable(p *Product) bool {
	quantity, exists := inv.products[p]
	if exists && quantity > 0 {
		return true
	}
	return false
}
