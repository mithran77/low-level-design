package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type HotelManagementSystem struct {
	guests       map[string]*Guest
	rooms        map[string]*Room
	reservations map[string]*Reservation
	mu           sync.RWMutex
}

var (
	instance *HotelManagementSystem
	once     sync.Once
)

func GetHotelManagementSystem() *HotelManagementSystem {
	once.Do(func() {
		instance = &HotelManagementSystem{
			guests:       make(map[string]*Guest),
			rooms:        make(map[string]*Room),
			reservations: make(map[string]*Reservation),
		}
	})
	return instance
}

func (s *HotelManagementSystem) AddGuest(guest *Guest) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.guests[guest.ID] = guest
}

func (s *HotelManagementSystem) GetGuest(guest_id string) *Guest {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.guests[guest_id]
}

func (s *HotelManagementSystem) AddRoom(room *Room) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rooms[room.ID] = room
}

func (s *HotelManagementSystem) GetRoom(room_id string) *Room {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.rooms[room_id]
}

func (s *HotelManagementSystem) BookRoom(guest *Guest, room *Room, checkInDate time.Time, checkOutDate time.Time) (*Reservation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := room.Book(); err != nil {
		return nil, err
	}

	rid := s.generateReservationID()
	reservation := NewReservation(rid, guest, room, checkInDate, checkOutDate)
	s.reservations[reservation.ID] = reservation

	return reservation, nil
}

func (s *HotelManagementSystem) CancelReservation(reservation_id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	reservation, exisits := s.reservations[reservation_id]
	if !exisits {
		return fmt.Errorf("Reservation not found")
	}
	err := reservation.Cancel()
	if err != nil {
		return err
	}

	delete(s.reservations, reservation_id)
	return nil
}

func (s *HotelManagementSystem) CheckIn(reservation_id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	reservation, exisits := s.reservations[reservation_id]
	if !exisits {
		return fmt.Errorf("Reservation not found")
	}

	err := reservation.Room.CheckIn()
	if err != nil {
		return err
	}

	return nil
}

func (s *HotelManagementSystem) CheckOut(reservation_id string, payment Payment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	reservation, exisits := s.reservations[reservation_id]
	if !exisits {
		return fmt.Errorf("Reservation not found")
	}

	if reservation.Status != ReservationStatusConfirmed {
		return fmt.Errorf("invalid reservation status")
	}

	days := reservation.CheckOutDate.Sub(reservation.CheckInDate).Hours() / 24
	amount := days * reservation.Room.Price

	err := payment.ProcessPayment(amount)
	if err != nil {
		return err
	}

	err = reservation.Room.CheckOut()
	if err != nil {
		return err
	}

	delete(s.reservations, reservation_id)
	return nil
}

func (s *HotelManagementSystem) generateReservationID() string {
	reservation_id, _ := uuid.NewV7()
	return reservation_id.String()
}
