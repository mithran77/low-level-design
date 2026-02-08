package main

import (
	"fmt"
	"sync"
)

type Room struct {
	ID     string
	Type   RoomType
	Price  float64
	status RoomStatus
	mu     sync.RWMutex
}

func NewRoom(id string, rt RoomType, price float64) *Room {
	return &Room{
		ID:     id,
		Type:   rt,
		Price:  price,
		status: RoomStatusAvailable,
	}
}

func (r *Room) Book() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.status != RoomStatusAvailable {
		return fmt.Errorf("Room not available")
	}
	r.status = RoomStatusBooked
	return nil
}

func (r *Room) CheckIn() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.status != RoomStatusBooked {
		return fmt.Errorf("Room not Booked")

	}
	r.status = RoomStatusOccupied
	return nil
}

func (r *Room) CheckOut() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.status != RoomStatusOccupied {
		fmt.Println("Room not Occupied")
	}
	r.status = RoomStatusAvailable
	return nil
}

func (r *Room) GetStatus() RoomStatus {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.status
}
