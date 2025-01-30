package main

import "fmt"

type ReadyState struct {
	vendingMachine *VendingMachine
}

func (s *ReadyState) SelectProduct(product *Product) {
	fmt.Println("Product already selected. Please make payment.")
}

func (s *ReadyState) InsertCoin(coin Coin) {
	vm := s.vendingMachine
	vm.AddCoin(coin)
	s.checkPaymentStatus()
}

func (s *ReadyState) InsertNote(note Note) {
	vm := s.vendingMachine
	vm.AddNote(note)
	s.checkPaymentStatus()
}

func (s *ReadyState) DispenceProduct() { fmt.Println("Please make payment first.") }
func (s *ReadyState) ReturnChange() {
	vm := s.vendingMachine
	// change := vm.total_payment - vm.selected_product.price
	if vm.total_payment > 0 {
		fmt.Println("Change returned: ", vm.total_payment)
		vm.ResetPayment()
	} else {
		fmt.Println("No change to return.")
	}
	vm.ResetSelection()
	vm.SetState(vm.idle_state)
}

func (s *ReadyState) checkPaymentStatus() {
	vm := s.vendingMachine
	if vm.total_payment >= vm.selected_product.price {
		vm.SetState(vm.dispence_state)
	}
}
