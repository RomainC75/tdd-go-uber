package gateways

import "tdd-go-uber/src/rideBooking-context/business-logic/models"

type UserRepo interface {
	GetUser(email string) models.User
}
