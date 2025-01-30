package main

import "fmt"

type ReturnChangeState struct {
	vendingMachine *VendingMachine
}

func (s *ReturnChangeState) SelectProduct(product *Product) {
	fmt.Println("Please collect change.")
}

func (s *ReturnChangeState) InsertCoin(coin Coin) {
	fmt.Println("Please collect change.")
	// Return Coin
}
func (s *ReturnChangeState) InsertNote(note Note) {
	fmt.Println("Please collect change.")
	// Return Note
}
func (s *ReturnChangeState) DispenceProduct() { fmt.Println("Please collect change first.") }

func (s *ReturnChangeState) ReturnChange() {
	vm := s.vendingMachine
	change := vm.total_payment - vm.selected_product.price
	if change > 0 {
		fmt.Println("Change returned: ", (change))
	} else {
		fmt.Println("No change to return.")
	}
	vm.ResetSelection()
	vm.ResetPayment()
	vm.SetState(vm.idle_state)
}
