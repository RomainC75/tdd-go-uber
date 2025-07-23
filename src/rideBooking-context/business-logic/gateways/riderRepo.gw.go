package gateways

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type RiderRepo interface {
	GetRider(id uuid.UUID) (models.Rider, error)
}
