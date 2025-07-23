package repositories

import "tdd-go-uber/src/rideBooking-context/business-logic/models"

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
