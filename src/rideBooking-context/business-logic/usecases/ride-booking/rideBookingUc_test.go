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

func (suite *RideTestSuite) TestRide() {
	userUUID := "0510c938-138b-4860-b5a7-c1bcb71719df"
	rideUUID := "d57f6854-c4ea-45d6-bbee-5d395002a279"
	nowDate, _ := time.Parse(time.RFC3339, "2025-01-01T15:04:05Z")
	defaultBirhday, _ := time.Parse(time.RFC3339, "2025-01-10T15:04:05Z")

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
			fakeDeterministicTime := providers.NewDeterministicTime()

			fakeUuidGenerator := providers.NewFakeUuidGenerator()
			fakeUuidGenerator.ExpectedUuid = uuid.MustParse(rideUUID)

			fakeTripProvider := providers.NewFakeTripScannerProvider()
			fakeTripProvider.Distance = testCase.distance

			fakeUserRepo := repositories.NewFakeUserRepo()
			fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(userUUID), "blop", testCase.Forfait, defaultBirhday)

			rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider, fakeUuidGenerator, fakeDeterministicTime)

			ride, err := rideBookingUc.Book(TBook{uuid.MustParse(userUUID), testCase.startAddr, testCase.endAddr, false})
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedBasePrice, ride.GetTotalPrice())

		}
	})

	suite.T().Run("should return an error if the user is not found", func(t *testing.T) {
		fakeDeterministicTime := providers.NewDeterministicTime()

		fakeUuidGenerator := providers.NewFakeUuidGenerator()

		fakeTripProvider := providers.NewFakeTripScannerProvider()
		fakeTripProvider.Distance = 10

		fakeUserRepo := repositories.NewFakeUserRepo()
		fakeUserRepo.ShouldReturnAnError = true

		rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider, fakeUuidGenerator, fakeDeterministicTime)

		_, err := rideBookingUc.Book(TBook{uuid.MustParse(userUUID), TAdressInput{}, TAdressInput{}, false})
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf("user %s not found", userUUID))
	})

	suite.T().Run("should book a UberX when distance is more than 3 km price", func(t *testing.T) {
		fakeDeterministicTime := providers.NewDeterministicTime()

		fakeUuidGenerator := providers.NewFakeUuidGenerator()

		startAddr := TAdressInput{11, "boulevard poissonière", 75002, "paris"}
		endAddr := TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}
		var distance float32 = 3.0
		var expectedPrice float32 = 30

		fakeTripProvider := providers.NewFakeTripScannerProvider()
		fakeTripProvider.Distance = distance

		fakeUserRepo := repositories.NewFakeUserRepo()
		fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(userUUID), "blop", valueobjects.ForfaitPremium, defaultBirhday)

		rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider, fakeUuidGenerator, fakeDeterministicTime)

		ride, err := rideBookingUc.Book(TBook{uuid.MustParse(userUUID), startAddr, endAddr, true})
		assert.Nil(t, err)
		assert.Equal(t, expectedPrice, ride.GetTotalPrice())
	})

	suite.T().Run("should return an error is distance is less than 3 km for an UberX ride", func(t *testing.T) {
		fakeDeterministicTime := providers.NewDeterministicTime()

		fakeUuidGenerator := providers.NewFakeUuidGenerator()

		startAddr := TAdressInput{11, "boulevard poissonière", 75002, "paris"}
		endAddr := TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}
		var distance float32 = 2.0

		fakeTripProvider := providers.NewFakeTripScannerProvider()
		fakeTripProvider.Distance = distance

		fakeUserRepo := repositories.NewFakeUserRepo()
		fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(userUUID), "blop", valueobjects.ForfaitPremium, defaultBirhday)

		rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider, fakeUuidGenerator, fakeDeterministicTime)

		_, err := rideBookingUc.Book(TBook{uuid.MustParse(userUUID), startAddr, endAddr, true})

		assert.EqualError(t, err, "distance cannot be < 3 when uberX")
	})

	suite.T().Run("uberX free if today is the birthday of the user", func(t *testing.T) {
		fakeDeterministicTime := providers.NewDeterministicTime()
		customBirthday, _ := time.Parse(time.RFC3339, "1980-01-01T15:04:05Z")
		fakeDeterministicTime.ExpectedTime = nowDate

		fakeUuidGenerator := providers.NewFakeUuidGenerator()
		fakeUuidGenerator.ExpectedUuid = uuid.MustParse(rideUUID)

		startAddr := TAdressInput{11, "boulevard poissonière", 75002, "paris"}
		endAddr := TAdressInput{7, "chemin du trou de l'hotel", 91300, "Massy"}
		var distance float32 = 3.0
		var expectedPrice float32 = 20

		fakeTripProvider := providers.NewFakeTripScannerProvider()
		fakeTripProvider.Distance = distance

		fakeUserRepo := repositories.NewFakeUserRepo()
		fakeUserRepo.ExpectedUser = *models.NewUser(uuid.MustParse(userUUID), "blop", valueobjects.ForfaitPremium, customBirthday)

		rideBookingUc := NewRideBookingUc(fakeUserRepo, fakeTripProvider, fakeUuidGenerator, fakeDeterministicTime)

		ride, err := rideBookingUc.Book(TBook{uuid.MustParse(userUUID), startAddr, endAddr, true})
		assert.Nil(t, err)
		assert.Equal(t, expectedPrice, ride.GetTotalPrice())
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RideTestSuite))
}
