package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type MovieTicketBookingSystem struct {
	users    map[string]*User
	movies   map[string]*Movie
	theatres map[string]*Theatre
	shows    map[string]*Show
	bookings map[string]*Booking
	// booking_count int64
	mu sync.RWMutex
}

var (
	instance *MovieTicketBookingSystem
	once     sync.Once
)

func GetMovieTicketBookingSystem() *MovieTicketBookingSystem {
	once.Do(func() {
		instance = &MovieTicketBookingSystem{
			users:    make(map[string]*User),
			movies:   make(map[string]*Movie),
			theatres: make(map[string]*Theatre),
			shows:    make(map[string]*Show),
			bookings: make(map[string]*Booking),
		}
	})
	return instance
}

func (m *MovieTicketBookingSystem) AddUser(user *User) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users[user.GetID()] = user
}

func (m *MovieTicketBookingSystem) AddTheatre(theatre *Theatre) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.theatres[theatre.GetID()] = theatre
}

func (m *MovieTicketBookingSystem) AddMovie(movie *Movie) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.movies[movie.GetID()] = movie
}

func (m *MovieTicketBookingSystem) AddShow(show *Show) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.shows[show.GetID()] = show
}

func (m *MovieTicketBookingSystem) BookTickets(user *User, show *Show, selected_seats []*Seat) (*Booking, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Are seats available?
	available, err := m.areSeatsAvailable(show, selected_seats)
	if !available {
		return nil, err
	}

	// Mark seats as booked
	m.markSeatsAsBooked(show, selected_seats)

	// Calculate total_price
	price := m.calculateTotalPrice(selected_seats)

	// Create pending booking
	bid := m.generateV4ID()
	booking := NewBooking(bid, user, show, selected_seats, price, BookingStatusPending)
	m.bookings[bid] = booking

	return booking, nil
}

func (m *MovieTicketBookingSystem) areSeatsAvailable(show *Show, selected_seats []*Seat) (bool, error) {
	for _, seat := range selected_seats {
		show_seat, exists := show.Seats[seat.ID]
		if !exists || show_seat.GetStatus() != SeatStatusAvailable {
			return false, fmt.Errorf("seat %s is not available", show_seat.ID)
		}
	}
	return true, nil
}

func (m *MovieTicketBookingSystem) markSeatsAsBooked(show *Show, selected_seats []*Seat) {
	for _, seat := range selected_seats {
		show.Seats[seat.ID].SetStatus(SeatStatusBooked)
	}
}

func (m *MovieTicketBookingSystem) calculateTotalPrice(selected_seats []*Seat) float64 {
	var price float64
	for _, seat := range selected_seats {
		price += seat.GetPrice()
	}
	return price
}

func (m *MovieTicketBookingSystem) generateV4ID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

// func (m *MovieTicketBookingSystem) generateV7ID() string {
// 	id, _ := uuid.NewV7()
// 	return id.String()
// }

func (m *MovieTicketBookingSystem) ConfirmBooking(booking_id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	booking, exisits := m.bookings[booking_id]
	if !exisits {
		return fmt.Errorf("booking not found")
	}
	if booking.GetStatus() != BookingStatusPending {
		return fmt.Errorf("booking is not in pending state")
	}
	booking.SetStatus(BookingStatusConfirmed)
	return nil
}

func (m *MovieTicketBookingSystem) CancelBooking(booking_id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	booking, exisits := m.bookings[booking_id]
	if !exisits {
		return fmt.Errorf("booking not found")
	}
	if booking.GetStatus() == BookingStatusCancelled {
		return fmt.Errorf("booking is already cancelled")
	}

	booking.SetStatus(BookingStatusCancelled)
	//Release seats
	m.markSeatsAsAvailable(booking.show, booking.Seats)

	return nil
}

func (m *MovieTicketBookingSystem) markSeatsAsAvailable(show *Show, selected_seats []*Seat) {
	for _, seat := range selected_seats {
		show.Seats[seat.ID].SetStatus(SeatStatusAvailable)
	}
}

// Create utility function for demo
func CreateSeats(rows, columns int) map[string]*Seat {
	seats := make(map[string]*Seat)
	for row := 1; row <= rows; row++ {
		for col := 1; col <= columns; col++ {
			seatID := fmt.Sprintf("%d-%d", row, col)
			seatType := SeatTypeNormal
			price := 100.0

			if row <= 2 {
				seatType = SeatTypePremium
				price = 150.0
			}

			seats[seatID] = NewSeat(
				seatID,
				row,
				col,
				seatType,
				price,
				SeatStatusAvailable,
			)
		}
	}
	return seats
}
