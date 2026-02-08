package main

import (
	"fmt"
	"time"
)

func Run() {
	bookingSystem := GetMovieTicketBookingSystem()

	// Add movies
	movie1 := NewMovie("M1", "Movie 1", "Description 1", 120)
	movie2 := NewMovie("M2", "Movie 2", "Description 2", 135)
	bookingSystem.AddMovie(movie1)
	bookingSystem.AddMovie(movie2)

	// Add theaters
	theatre1 := NewTheatre("T1", "Theater 1", "Location 1", CreateSeats(10, 10))
	theatre2 := NewTheatre("T2", "Theater 2", "Location 2", CreateSeats(8, 8))
	bookingSystem.AddTheatre(theatre1)
	bookingSystem.AddTheatre(theatre2)

	// Add shows
	show1 := NewShow(
		"S1",
		movie1,
		theatre1,
		time.Now(),
		time.Now().Add(time.Duration(movie1.duration_in_minutes)*time.Minute),
	)
	show2 := NewShow(
		"S2",
		movie2,
		theatre2,
		time.Now(),
		time.Now().Add(time.Duration(movie2.duration_in_minutes)*time.Minute),
	)

	bookingSystem.AddShow(show1)
	bookingSystem.AddShow(show2)

	// Create user
	user := NewUser("U1", "John Doe", "john@example.com")

	// Select seats
	selectedSeats := []*Seat{
		show1.Seats["1-5"],
		show1.Seats["1-6"],
	}

	// Book tickets
	booking, err := bookingSystem.BookTickets(user, show1, selectedSeats)
	if err != nil {
		fmt.Printf("Booking failed: %v\n", err)
		return
	}

	fmt.Printf("Booking successful. Booking ID: %s\n", booking.GetID())

	// Confirm booking
	// Make Payment
	if err := bookingSystem.ConfirmBooking(booking.GetID()); err != nil {
		fmt.Printf("Failed to confirm booking: %v\n", err)
		return
	}
	fmt.Println("Booking confirmed")

	// Cancel booking
	if err := bookingSystem.CancelBooking(booking.GetID()); err != nil {
		fmt.Printf("Failed to cancel booking: %v\n", err)
		return
	}
	fmt.Printf("Booking canceled. Booking ID: %s\n", booking.GetID())
}

func main() {
	Run()
}
