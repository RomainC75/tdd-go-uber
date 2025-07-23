package models

import (
	"time"

	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

const ONE_YEAR_IN_SECONDS = 31622400

type User struct {
	id          uuid.UUID
	name        string
	forfait     valueobjects.Forfait
	birthday    time.Time
	inscription time.Time
}

func NewUser(id uuid.UUID, name string, forfait valueobjects.Forfait, birthday time.Time, inscription time.Time) *User {
	return &User{
		id:          id,
		name:        name,
		forfait:     forfait,
		birthday:    birthday,
		inscription: inscription,
	}
}

func (u *User) GetForfait() valueobjects.Forfait {
	return u.forfait
}

func (u *User) IsBirthday(now time.Time) bool {
	if u.birthday.Day() == now.Day() && u.birthday.Month() == now.Month() {
		return true
	}
	return false
}

func (u *User) IsNewUser(now time.Time) bool {
	d := now.Sub(u.inscription)
	return d.Seconds() < ONE_YEAR_IN_SECONDS
}
