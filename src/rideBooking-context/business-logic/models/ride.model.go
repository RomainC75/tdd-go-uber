package models

import (
	"time"

	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

type Ride struct {
	id      uuid.UUID
	rider   Rider
	trip    valueobjects.Trip
	isUberX bool
}

func BookNewRide(id uuid.UUID, rider Rider, trip valueobjects.Trip, isUberX bool, now time.Time) Ride {
	trip.SetTotalCost(rider.forfait, isUberX, rider.IsBirthday(now), rider.IsNewRider(now))
	return Ride{
		id:      id,
		rider:   rider,
		trip:    trip,
		isUberX: isUberX,
	}
}

func (r Ride) GetTotalPrice() float32 {
	return r.trip.GetTotalPrice()
}

func (r Ride) GetUuid() uuid.UUID {
	return r.id
}
