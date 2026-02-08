package main

type Request struct {
	source_floor      int
	destination_floor int
}

func NewRequest(sfloor int, dfloor int) *Request {
	return &Request{
		source_floor:      sfloor,
		destination_floor: dfloor,
	}
}
