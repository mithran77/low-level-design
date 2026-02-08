package main

type Guest struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber string
}

func NewGuest(id string, name string, email string, phoneNumber string) *Guest {
	return &Guest{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func (g *Guest) GetID() string {
	return g.ID
}

func (g *Guest) GetName() string {
	return g.Name
}

func (g *Guest) GetEmail() string {
	return g.Email
}

func (g *Guest) GetPhone() string {
	return g.PhoneNumber
}
