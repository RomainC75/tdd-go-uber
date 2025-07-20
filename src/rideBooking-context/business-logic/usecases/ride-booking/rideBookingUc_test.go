package ridebooking

import (
	"testing"

	"tdd-go-uber/src/rideBooking-context/adapters/secondary/providers"

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
	}{
		{"PARIS", "PARIS", 3, 31.5},
		{"OUT", "PARIS", 3, 11.5},
		{"PARIS", "OUT", 3, 21.5},
		{"OUT", "OUT", 3, 51.5},
	}
	suite.T().Run("should calculate price", func(t *testing.T) {
		for _, testCase := range testCases {
			fakeTripProvider := providers.NewFakeProvider()
			fakeTripProvider.Distance = testCase.distance
			rideBookingUc := NewRideBookingUc(fakeTripProvider)

			// basePrice := getBasePrice(testCase.startAddr, testCase.endAddr)
			basePrice := rideBookingUc.Book(TBook{testCase.startAddr, testCase.endAddr})
			assert.Equal(t, testCase.expectedBasePrice, basePrice)

		}
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RideTestSuite))
}
