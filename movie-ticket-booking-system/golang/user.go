package main

type User struct {
	id    string
	name  string
	email string
}

func NewUser(id string, name string, email string) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

func (u *User) GetID() string {
	return u.id
}
