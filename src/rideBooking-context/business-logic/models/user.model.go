package models

import (
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

type User struct {
	id      uuid.UUID
	name    string
	forfait valueobjects.Forfait
}

func NewUser(id uuid.UUID, name string, forfait valueobjects.Forfait) *User {
	return &User{
		id:      id,
		name:    name,
		forfait: forfait,
	}
}

func (u *User) GetForfait() valueobjects.Forfait {
	return u.forfait
}
