package main

import (
	"time"
)

type Show struct {
	id         string
	movie      *Movie
	theatre    *Theatre
	start_time time.Time
	end_time   time.Time
	Seats      map[string]*Seat
}

func NewShow(id string, movie *Movie, theatre *Theatre, stime time.Time, etime time.Time) *Show {
	return &Show{
		id:         id,
		movie:      movie,
		theatre:    theatre,
		start_time: stime,
		end_time:   etime,
		Seats:      theatre.Seats,
	}
}

func (s *Show) GetID() string {
	return s.id
}
