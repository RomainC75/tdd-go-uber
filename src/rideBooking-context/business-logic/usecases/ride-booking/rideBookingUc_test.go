package ridebooking

import (
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
	}{
		{TAdressInput{11, "boulevard poissonière", 75002, "Paris"}, TAdressInput{11, "boulevard poissonière", 75002, "paris"}, 3, 31.5, models.ForfaitBasic},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{11, "boulevard poissonière", 75002, "paris"}, 3, 11.5, models.ForfaitBasic},
		{TAdressInput{11, "boulevard poissonière", 75002, "paris"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 21.5, models.ForfaitBasic},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 51.5, models.ForfaitBasic},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 6, 50.5, models.ForfaitPremium},
		{TAdressInput{11, "boulevard poissonière", 75002, "paris"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 20, models.ForfaitPremium},
	}
	suite.T().Run("should calculate price", func(t *testing.T) {
		for _, testCase := range testCases {
			fakeTripProvider := providers.NewFakeTripScannerProvider()
			fakeTripProvider.Distance = testCase.distance

			fakeUserRepo := repositories.NewFakeUserRepo()
			fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(UUID), "blop", testCase.Forfait)

			rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider)

			trip, err := rideBookingUc.Book(TBook{uuid.MustParse(UUID), testCase.startAddr, testCase.endAddr})
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedBasePrice, trip.GetTotalPrice())

		}
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RideTestSuite))
}
