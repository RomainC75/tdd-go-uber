package models

import (
	"fmt"
	"time"

	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

type Ride struct {
	id      uuid.UUID
	user    User
	trip    valueobjects.Trip
	isUberX bool
}

func BookNewRide(id uuid.UUID, user User, trip valueobjects.Trip, isUberX bool, now time.Time) Ride {
	trip.SetTotalCost(user.forfait, isUberX, user.IsBirthday(now), user.IsNewUser(now))
	fmt.Printf("-----------------IS NEW : %t \n", user.IsNewUser(now))
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
