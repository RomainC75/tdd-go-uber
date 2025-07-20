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
		expectedBasePrice float32
	}{
		{"PARIS", "PARIS", 30},
		{"OUT", "PARIS", 10},
		{"PARIS", "OUT", 20},
		{"OUT", "OUT", 50},
	}
	suite.T().Run("should calculate price", func(t *testing.T) {
		for _, testCase := range testCases {
			fakeTripProvider := providers.NewFakeProvider()
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
