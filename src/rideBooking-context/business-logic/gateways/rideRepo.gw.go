package gateways

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type IRideRepo interface {
	Save(ride models.Ride) error
	GetById(rideId uuid.UUID) (models.Ride, error)
	DeleteById(rideId uuid.UUID) error
}
