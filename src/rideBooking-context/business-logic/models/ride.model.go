package models

import (
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"
)

type Ride struct {
	user User
	trip valueobjects.Trip
}

func BookNewRide(user User, trip valueobjects.Trip) Ride {
	return Ride{
		user: user,
		trip: trip,
	}
}

func (r Ride) GetTotalPrice() float32 {
	return r.trip.GetTotalPrice()
}
