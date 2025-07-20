package ridebooking

import (
	"testing"

	"tdd-go-uber/src/rideBooking-context/adapters/secondary/providers"
	"tdd-go-uber/src/rideBooking-context/adapters/secondary/repositories"
	"tdd-go-uber/src/rideBooking-context/business-logic/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RideTestSuite struct {
	suite.Suite
}

func (suite *RideTestSuite) TestRide() {
	testCases := []struct {
		startAddr         string
		endAddr           string
		distance          float32
		expectedBasePrice float32
		Forfait           models.Forfait
	}{
		{"PARIS", "PARIS", 3, 31.5, models.ForfaitBasic},
		{"OUT", "PARIS", 3, 11.5, models.ForfaitBasic},
		{"PARIS", "OUT", 3, 21.5, models.ForfaitBasic},
		{"OUT", "OUT", 3, 51.5, models.ForfaitBasic},
		{"OUT", "OUT", 6, 50.5, models.ForfaitPremium},
		{"PARIS", "OUT", 3, 20, models.ForfaitBasic},
	}
	suite.T().Run("should calculate price", func(t *testing.T) {
		for _, testCase := range testCases {
			fakeTripProvider := providers.NewFakeTripScannerProvider()
			fakeTripProvider.Distance = testCase.distance

			fakeUserRepo := repositories.NewFakeUserRepo()
			fakeUserRepo.ExpectedUser = *models.NewUser("blop", testCase.Forfait)

			rideBookingUc := NewRideBookingUc(fakeTripProvider)

			trip := rideBookingUc.Book(TBook{testCase.startAddr, testCase.endAddr})
			assert.Equal(t, testCase.expectedBasePrice, trip.GetTotalPrice())

		}
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RideTestSuite))
}
