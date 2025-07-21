package repositories

import (
	"fmt"

	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type FakeUserRepo struct {
	ExpectedUser        models.User
	ShouldReturnAnError bool
}

func NewFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{}
}

func (fur *FakeUserRepo) GetUser(id uuid.UUID) (models.User, error) {
	if fur.ShouldReturnAnError {
		return models.User{}, fmt.Errorf("user %s not found", id.String())
	}
	return fur.ExpectedUser, nil
}
