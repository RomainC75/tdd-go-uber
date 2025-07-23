package gateways

import "tdd-go-uber/src/rideBooking-context/business-logic/models"

type IRideRepo interface {
	Save(ride models.Ride) error
}
