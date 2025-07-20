package repositories

import "tdd-go-uber/src/rideBooking-context/business-logic/models"

type FakeUserRepo struct {
	ExpectedUser models.User
}

func NewFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{}
}

func (fur *FakeUserRepo) GetUser(email string) models.User {
	return fur.ExpectedUser
}
