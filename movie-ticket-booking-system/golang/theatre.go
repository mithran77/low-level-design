package main

type Theatre struct {
	id       string
	name     string
	location string
	Seats    map[string]*Seat
}

func NewTheatre(id string, name string, location string, seats map[string]*Seat) *Theatre {
	return &Theatre{
		id:       id,
		name:     name,
		location: location,
		Seats:    seats,
	}
}

func (t *Theatre) GetID() string {
	return t.id
}
