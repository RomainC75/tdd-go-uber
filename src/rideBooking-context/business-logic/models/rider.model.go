package models

import (
	"time"

	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

const ONE_YEAR_IN_SECONDS = 31622400

type Rider struct {
	id          uuid.UUID
	name        string
	forfait     valueobjects.Forfait
	birthday    time.Time
	inscription time.Time
}

func NewRider(id uuid.UUID, name string, forfait valueobjects.Forfait, birthday time.Time, inscription time.Time) *Rider {
	return &Rider{
		id:          id,
		name:        name,
		forfait:     forfait,
		birthday:    birthday,
		inscription: inscription,
	}
}

func (u *Rider) GetForfait() valueobjects.Forfait {
	return u.forfait
}

func (u *Rider) IsBirthday(now time.Time) bool {
	if u.birthday.Day() == now.Day() && u.birthday.Month() == now.Month() {
		return true
	}
	return false
}

func (u *Rider) IsNewRider(now time.Time) bool {
	d := now.Sub(u.inscription)
	return d.Seconds() < ONE_YEAR_IN_SECONDS
}

func (u *Rider) GetUuid() uuid.UUID {
	return u.id
}
