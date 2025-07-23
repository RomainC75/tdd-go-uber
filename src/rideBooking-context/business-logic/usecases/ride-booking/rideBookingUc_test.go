package ridebooking

import (
	"fmt"
	"testing"
	"time"

	"tdd-go-uber/src/rideBooking-context/adapters/secondary/providers"
	"tdd-go-uber/src/rideBooking-context/adapters/secondary/repositories"
	"tdd-go-uber/src/rideBooking-context/business-logic/models"
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RideTestSuite struct {
	suite.Suite
}

func dependenciesInstanciator() (*repositories.FakeRideRepo, *repositories.FakeRiderRepo, *providers.FakeTripProvider, *providers.FakeUuidGenerator, *providers.FakeDeterministicTime) {
	fdt := providers.NewDeterministicTime()
	fug := providers.NewFakeUuidGenerator()
	ftsp := providers.NewFakeTripScannerProvider()
	fRiderR := repositories.NewFakeRiderRepo()
	fRideR := repositories.NewFakeRideRepo()
	return fRideR, fRiderR, ftsp, fug, fdt
}

func (suite *RideTestSuite) TestRide() {
	riderUUID := "0510c938-138b-4860-b5a7-c1bcb71719df"
	rideUUID := "d57f6854-c4ea-45d6-bbee-5d395002a279"
	nowDate, _ := time.Parse(time.RFC3339, "2025-01-01T15:04:05Z")
	defaultBirhday, _ := time.Parse(time.RFC3339, "1980-01-10T15:04:05Z")
	defaultInscription, _ := time.Parse(time.RFC3339, "2023-12-10T15:04:05Z")
	defaultStartAddr := TAdressInput{11, "boulevard poissonière", 75002, "paris"}
	defaultEndAddr := TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}

	testCases := []struct {
		startAddr         TAdressInput
		endAddr           TAdressInput
		distance          float32
		expectedBasePrice float32
		Forfait           valueobjects.Forfait
		isUberX           bool
	}{
		{TAdressInput{11, "boulevard poissonière", 75002, "Paris"}, TAdressInput{11, "boulevard poissonière", 75002, "paris"}, 3, 31.5, valueobjects.ForfaitBasic, false},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{11, "boulevard poissonière", 75002, "paris"}, 3, 11.5, valueobjects.ForfaitBasic, false},
		{TAdressInput{11, "boulevard poissonière", 75002, "paris"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 21.5, valueobjects.ForfaitBasic, false},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 51.5, valueobjects.ForfaitBasic, false},
		{TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 6, 50.5, valueobjects.ForfaitPremium, false},
		{TAdressInput{11, "boulevard poissonière", 75002, "paris"}, TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}, 3, 20, valueobjects.ForfaitPremium, false},
	}
	suite.T().Run("should calculate price", func(t *testing.T) {
		for _, testCase := range testCases {
			fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()

			fdt.ExpectedTime = nowDate
			fug.ExpectedUuid = uuid.MustParse(rideUUID)
			ftsp.Distance = testCase.distance
			fRiderR.ExpectedRider = *models.NewRider(uuid.MustParse(riderUUID), "blop", testCase.Forfait, defaultBirhday, defaultInscription)

			rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)
			ride, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), testCase.startAddr, testCase.endAddr, false})
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedBasePrice, ride.GetTotalPrice())

		}
	})

	suite.T().Run("should return an error if the rider is not found", func(t *testing.T) {
		fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()

		fdt.ExpectedTime = nowDate
		ftsp.Distance = 10
		fRiderR.ShouldReturnAnError = true
		rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)

		_, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), TAdressInput{}, TAdressInput{}, false})
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf("rider %s not found", riderUUID))
	})

	suite.T().Run("should book a UberX when distance is more than 3 km price", func(t *testing.T) {
		fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()

		fdt.ExpectedTime = nowDate

		var distance float32 = 3.0
		var expectedPrice float32 = 30

		ftsp.Distance = distance

		fRiderR.ExpectedRider = *models.NewRider(uuid.MustParse(riderUUID), "blop", valueobjects.ForfaitPremium, defaultBirhday, defaultInscription)

		rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)

		ride, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), defaultStartAddr, defaultEndAddr, true})
		assert.Nil(t, err)
		assert.Equal(t, expectedPrice, ride.GetTotalPrice())
	})

	suite.T().Run("should return an error is distance is less than 3 km for an UberX ride", func(t *testing.T) {
		fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()
		fdt.ExpectedTime = nowDate

		var distance float32 = 2.0
		ftsp.Distance = distance

		fRiderR.ExpectedRider = *models.NewRider(uuid.MustParse(riderUUID), "blop", valueobjects.ForfaitPremium, defaultBirhday, defaultInscription)

		rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)

		_, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), defaultStartAddr, defaultEndAddr, true})

		assert.EqualError(t, err, "distance cannot be < 3 when uberX")
	})

	suite.T().Run("uberX free if today is the birthday of the rider", func(t *testing.T) {
		fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()
		fdt.ExpectedTime = nowDate
		customBirthday, _ := time.Parse(time.RFC3339, "1980-01-01T15:04:05Z")

		fug.ExpectedUuid = uuid.MustParse(rideUUID)

		var distance float32 = 3.0
		var expectedPrice float32 = 20

		ftsp.Distance = distance

		fRiderR.ExpectedRider = *models.NewRider(uuid.MustParse(riderUUID), "blop", valueobjects.ForfaitPremium, customBirthday, defaultInscription)

		rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)

		ride, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), defaultStartAddr, defaultEndAddr, true})
		assert.Nil(t, err)
		assert.Equal(t, expectedPrice, ride.GetTotalPrice())
	})

	suite.T().Run("should get 5% because of a new rider (<1year)", func(t *testing.T) {
		fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()
		fdt.ExpectedTime = nowDate

		fug.ExpectedUuid = uuid.MustParse(rideUUID)

		var expectedPrice float32 = 28.5
		var distance float32 = 3.0

		ftsp.Distance = distance

		inscriptionDate, _ := time.Parse(time.RFC3339, "2024-01-10T15:04:05Z")
		fRiderR.ExpectedRider = *models.NewRider(uuid.MustParse(riderUUID), "blop", valueobjects.ForfaitPremium, defaultBirhday, inscriptionDate)

		rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)

		ride, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), defaultStartAddr, defaultEndAddr, true})
		assert.Nil(t, err)
		assert.Equal(t, expectedPrice, ride.GetTotalPrice())
	})
	suite.T().Run("should save the ride in db", func(t *testing.T) {
		fRideR, fRiderR, ftsp, fug, fdt := dependenciesInstanciator()
		fdt.ExpectedTime = nowDate
		ftsp.Distance = 3.0
		fug.ExpectedUuid = uuid.MustParse(rideUUID)
		fRiderR.ExpectedRider = *models.NewRider(uuid.MustParse(riderUUID), "blop", valueobjects.ForfaitPremium, defaultBirhday, defaultInscription)

		rideBookingUc := NewRideBookingUc(fRideR, fRiderR, ftsp, fug, fdt)

		_, err := rideBookingUc.Book(TBook{uuid.MustParse(riderUUID), defaultStartAddr, defaultEndAddr, true})
		assert.NoError(t, err)
		assert.Equal(t, len(fRideR.Rides), 1)
		assert.Equal(t, fRideR.Rides[0].GetUuid().String(), rideUUID)
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RideTestSuite))
}
