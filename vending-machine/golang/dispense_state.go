package main

import "fmt"

type DispenceState struct {
	vendingMachine *VendingMachine
}

func (s *DispenceState) SelectProduct(product *Product) {
	fmt.Println("Please collect product first.")
}
func (s *DispenceState) InsertCoin(coin Coin) {
	fmt.Println("Please collect product.")
	// Return Coin
}
func (s *DispenceState) InsertNote(note Note) {
	fmt.Println("Please collect product.")
	// Return Note
}

func (s *DispenceState) DispenceProduct() {
	vm := s.vendingMachine
	fmt.Println("Please collect: ", vm.selected_product.name)
	vm.SetState(vm.return_change_state)
}

func (s *DispenceState) ReturnChange() { fmt.Println("Please collect product first.") }
