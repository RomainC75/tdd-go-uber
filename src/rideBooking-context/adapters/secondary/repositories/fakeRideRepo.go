package repositories

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type FakeRideRepo struct {
	Rides []models.Ride
}

func NewFakeRideRepo() *FakeRideRepo {
	return &FakeRideRepo{}
}

func (frr *FakeRideRepo) Save(ride models.Ride) error {
	frr.Rides = append(frr.Rides, ride)
	return nil
}

func (frr *FakeRideRepo) GetById(rideId uuid.UUID) (models.Ride, error) {
	return frr.Rides[0], nil
}

func (frr *FakeRideRepo) DeleteById(rideId uuid.UUID) error {
	frr.Rides = []models.Ride{}
	return nil
}
