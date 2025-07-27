package ridecancel

import (
	"errors"

	"tdd-go-uber/src/rideBooking-context/business-logic/gateways"

	"github.com/google/uuid"
)

var ErrorUserNotAuthorizedToDeleteRide = errors.New("this user cannot cancel the ride")

type RideCancelUC struct {
	rideRepo gateways.IRideRepo
}

func NewRideCancel(rideRepo gateways.IRideRepo) *RideCancelUC {
	return &RideCancelUC{
		rideRepo: rideRepo,
	}
}

func (crr *RideCancelUC) Execute(userId uuid.UUID, rideId uuid.UUID) error {
	foundRide, err := crr.rideRepo.GetById(rideId)
	if err != nil {
	}
	if !foundRide.IdUserTheOwner(userId) {
		return ErrorUserNotAuthorizedToDeleteRide
	}
	err = crr.rideRepo.DeleteById(rideId)
	if err != nil {
	}
	return nil
}
