package main

type Movie struct {
	id                  string
	name                string
	description         string
	duration_in_minutes int
}

func NewMovie(id string, name string, description string, duration int) *Movie {
	return &Movie{
		id:                  id,
		name:                name,
		description:         description,
		duration_in_minutes: duration,
	}
}

func (m *Movie) GetID() string {
	return m.id
}
