package main

type BookingStatus int
type SeatType int
type SeatStatus int

const (
	BookingStatusPending BookingStatus = iota
	BookingStatusConfirmed
	BookingStatusCancelled
)

const (
	SeatTypeNormal SeatType = iota
	SeatTypePremium
)

const (
	SeatStatusAvailable SeatStatus = iota
	SeatStatusBooked
)
