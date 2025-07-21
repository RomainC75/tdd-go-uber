package repositories

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type FakeUserRepo struct {
	ExpectedUser models.User
}

func NewFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{}
}

func (fur *FakeUserRepo) GetUser(id uuid.UUID) (models.User, error) {
	return fur.ExpectedUser, nil
}
