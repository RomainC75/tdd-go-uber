package models

import (
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

type Ride struct {
	id      uuid.UUID
	user    User
	trip    valueobjects.Trip
	isUberX bool
}

func BookNewRide(id uuid.UUID, user User, trip valueobjects.Trip, isUberX bool) Ride {
	return Ride{
		id:      id,
		user:    user,
		trip:    trip,
		isUberX: isUberX,
	}
}

func (r Ride) GetTotalPrice() float32 {
	return r.trip.GetTotalPrice()
}
