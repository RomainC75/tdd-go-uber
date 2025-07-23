package repositories

import (
	"fmt"

	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
)

type FakeRiderRepo struct {
	ExpectedRider       models.Rider
	ShouldReturnAnError bool
}

func NewFakeRiderRepo() *FakeRiderRepo {
	return &FakeRiderRepo{}
}

func (fur *FakeRiderRepo) GetRider(id uuid.UUID) (models.Rider, error) {
	if fur.ShouldReturnAnError {
		return models.Rider{}, fmt.Errorf("rider %s not found", id.String())
	}
	return fur.ExpectedRider, nil
}
