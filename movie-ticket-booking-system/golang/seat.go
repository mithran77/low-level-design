package main

import "sync"

type Seat struct {
	ID     string
	row    int
	column int
	stype  SeatType
	price  float64
	status SeatStatus
	mu     sync.RWMutex
}

func NewSeat(id string, row int, col int, stype SeatType, price float64, status SeatStatus) *Seat {
	return &Seat{
		ID:     id,
		row:    row,
		column: col,
		stype:  stype,
		price:  price,
		status: status,
	}
}

func (s *Seat) GetID() string {
	return s.ID
}

func (s *Seat) GetPrice() float64 {
	return s.price
}

func (s *Seat) GetStatus() SeatStatus {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.status
}

func (s *Seat) SetStatus(status SeatStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.status = status
}
