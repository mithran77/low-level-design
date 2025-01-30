package main

type VendingMachineState interface {
	SelectProduct(product *Product)
	InsertCoin(coin Coin)
	InsertNote(note Note)
	DispenceProduct()
	ReturnChange()
}
