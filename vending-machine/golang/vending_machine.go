package main

type VendingMachine struct {
	inventory           *Inventory
	idle_state          *IdleState
	ready_state         *ReadyState
	dispence_state      *DispenceState
	return_change_state *ReturnChangeState
	current_state       VendingMachineState
	selected_product    *Product
	total_payment       float64
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		inventory: NewInventory(),
	}
	vm.idle_state = &IdleState{vm}
	vm.ready_state = &ReadyState{vm}
	vm.dispence_state = &DispenceState{vm}
	vm.return_change_state = &ReturnChangeState{vm}
	vm.current_state = vm.idle_state
	return vm
}

func (vm *VendingMachine) SelectProduct(product *Product) {
	vm.current_state.SelectProduct(product)
}

func (vm *VendingMachine) InsertCoin(coin Coin) {
	vm.current_state.InsertCoin(coin)
}

func (vm *VendingMachine) InsertNote(note Note) {
	vm.current_state.InsertNote(note)
}

func (vm *VendingMachine) DispenceProduct() {
	vm.current_state.DispenceProduct()
}

func (vm *VendingMachine) ReturnChange() {
	vm.current_state.ReturnChange()
}

func (vm *VendingMachine) SetState(state VendingMachineState) {
	vm.current_state = state
}

func (vm *VendingMachine) AddCoin(coin Coin) {
	vm.total_payment += float64(coin)
}

func (vm *VendingMachine) AddNote(note Note) {
	vm.total_payment += float64(note)
}

func (vm *VendingMachine) ResetPayment() {
	vm.total_payment = 0
}

func (vm *VendingMachine) ResetSelection() {
	vm.selected_product = nil
}
