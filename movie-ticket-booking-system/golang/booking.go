package main

import "sync"

type Booking struct {
	id          string
	user        *User
	show        *Show
	Seats       []*Seat
	total_price float64
	status      BookingStatus
	mu          sync.RWMutex
}

func NewBooking(id string, user *User, show *Show, seats []*Seat, price float64, status BookingStatus) *Booking {
	return &Booking{
		id:          id,
		user:        user,
		show:        show,
		Seats:       seats,
		total_price: price,
		status:      status,
	}
}

func (b *Booking) GetID() string {
	return b.id
}

func (b *Booking) GetStatus() BookingStatus {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.status
}

func (b *Booking) SetStatus(status BookingStatus) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.status = status
}
