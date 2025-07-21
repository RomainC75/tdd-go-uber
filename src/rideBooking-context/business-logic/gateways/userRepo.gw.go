package gateways

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUser(id uuid.UUID) (models.User, error)
}
