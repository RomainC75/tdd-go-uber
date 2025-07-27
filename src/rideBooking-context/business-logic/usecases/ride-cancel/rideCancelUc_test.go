package ridecancel

import (
	"testing"
	"time"

	"tdd-go-uber/src/rideBooking-context/adapters/secondary/repositories"
	"tdd-go-uber/src/rideBooking-context/business-logic/models"
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCancelRide(t *testing.T) {
	rideUUID := uuid.MustParse("d57f6854-c4ea-45d6-bbee-5d395002a279")
	riderUUID := uuid.MustParse("0510c938-138b-4860-b5a7-c1bcb71719df")

	t.Run("should cancel a ride if in running status", func(t *testing.T) {
		rider := models.NewRider(riderUUID, "bob@gmail.com", valueobjects.ForfaitBasic, time.Now(), time.Now())
		ride := models.BookNewRide(rideUUID, *rider, valueobjects.Trip{}, false, time.Now())
		fakeRideRepo := repositories.NewFakeRideRepo()
		fakeRideRepo.Rides = []models.Ride{ride}

		rideCancel := NewRideCancel(fakeRideRepo)
		rideCancel.Execute(riderUUID, rideUUID)

		assert.Equal(t, len(fakeRideRepo.Rides), 0)
	})
}
