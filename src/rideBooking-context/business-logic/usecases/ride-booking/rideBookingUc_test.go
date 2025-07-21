package ridebooking

import (
	"fmt"
	"testing"

	"tdd-go-uber/src/rideBooking-context/adapters/secondary/providers"
	"tdd-go-uber/src/rideBooking-context/adapters/secondary/repositories"
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RideTestSuite struct {
	suite.Suite
}

func (suite *RideTestSuite) TestRide() {
	UUID := "0510c938-138b-4860-b5a7-c1bcb71719df"
	testCases := []struct {
		startAddr         TAdressInput
		endAddr           TAdressInput
		distance          float32
		expectedBasePrice float32
		Forfait           models.Forfait
		isUberX           bool
	}{
		{TAdressInput{11, "boulevard poissonière", 75002, "Paris"}, TAdressInput{11, "boulevard poissonière", 75002, "paris"}, 3, 31.5, models.ForfaitBasic, false},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{11, "boulevard poissonière", 75002, "paris"}, 3, 11.5, models.ForfaitBasic, false},
		{TAdressInput{11, "boulevard poissonière", 75002, "paris"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 21.5, models.ForfaitBasic, false},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 51.5, models.ForfaitBasic, false},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 6, 50.5, models.ForfaitPremium, false},
		{TAdressInput{11, "boulevard poissonière", 75002, "paris"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 20, models.ForfaitPremium, false},
	}
	suite.T().Run("should calculate price", func(t *testing.T) {
		for _, testCase := range testCases {
			fakeTripProvider := providers.NewFakeTripScannerProvider()
			fakeTripProvider.Distance = testCase.distance

			fakeUserRepo := repositories.NewFakeUserRepo()
			fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(UUID), "blop", testCase.Forfait)

			rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider)

			trip, err := rideBookingUc.Book(TBook{uuid.MustParse(UUID), testCase.startAddr, testCase.endAddr, false})
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedBasePrice, trip.GetTotalPrice())

		}
	})

	suite.T().Run("should return an error if the user is not found", func(t *testing.T) {
		fakeTripProvider := providers.NewFakeTripScannerProvider()
		fakeTripProvider.Distance = 10

		fakeUserRepo := repositories.NewFakeUserRepo()
		fakeUserRepo.ShouldReturnAnError = true

		rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider)

		_, err := rideBookingUc.Book(TBook{uuid.MustParse(UUID), TAdressInput{}, TAdressInput{}, false})
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf("user %s not found", UUID))
	})

	suite.T().Run("should book a UberX when distance is more than 3 km price", func(t *testing.T) {
		startAddr := TAdressInput{11, "boulevard poissonière", 75002, "paris"}
		endAddr := TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}
		var distance float32 = 3.0
		var expectedPrice float32 = 30

		fakeTripProvider := providers.NewFakeTripScannerProvider()
		fakeTripProvider.Distance = distance

		fakeUserRepo := repositories.NewFakeUserRepo()
		fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(UUID), "blop", models.ForfaitPremium)

		rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider)

		trip, err := rideBookingUc.Book(TBook{uuid.MustParse(UUID), startAddr, endAddr, true})
		assert.Nil(t, err)
		assert.Equal(t, expectedPrice, trip.GetTotalPrice())
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RideTestSuite))
}
