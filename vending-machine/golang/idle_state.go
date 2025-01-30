package main

import "fmt"

type IdleState struct {
	vendingMachine *VendingMachine
}

func (s *IdleState) SelectProduct(product *Product) {
	vm := s.vendingMachine
	if vm.inventory.IsAvailable(product) {
		vm.selected_product = product
		vm.SetState(vm.ready_state)
		fmt.Println("Product selected: ", product.name)
	} else {
		fmt.Println("Product not available: ", product.name)
	}
}

func (s *IdleState) InsertCoin(coin Coin) {
	fmt.Println("Please Select a product first.")
	// Return coin
}
func (s *IdleState) InsertNote(note Note) {
	fmt.Println("Please Select a product first.")
	// Return Note
}
func (s *IdleState) DispenceProduct() { fmt.Println("Please Select a product and make payment.") }
func (s *IdleState) ReturnChange()    { fmt.Println("No change to return.") }
