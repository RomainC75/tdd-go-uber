package ridebooking

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/gateways"
	"tdd-go-uber/src/rideBooking-context/business-logic/models"
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

	"github.com/google/uuid"
)

type TAdressInput struct {
	number int
	street string
	code   int
	city   string
}

type TBook struct {
	userId    uuid.UUID
	startAddr TAdressInput
	endAddr   TAdressInput
	isUberX   bool
}

type BookOptions struct {
	isUberX bool
}

type RideBookingUc struct {
	tripScanner       gateways.ITripScanner
	userRepo          gateways.UserRepo
	uuidGenerator     gateways.IUUIDGenerator
	deterministicTime gateways.IDeterministicTime
}

func NewRideBookingUc(userRepo gateways.UserRepo, tripProvider gateways.ITripScanner, uuidGenerator gateways.IUUIDGenerator, deterministicTime gateways.IDeterministicTime) *RideBookingUc {
	return &RideBookingUc{
		userRepo:          userRepo,
		tripScanner:       tripProvider,
		uuidGenerator:     uuidGenerator,
		deterministicTime: deterministicTime,
	}
}

func (rbuc *RideBookingUc) Book(args TBook) (models.Ride, error) {
	foundUser, err := rbuc.userRepo.GetUser(args.userId)
	if err != nil {
		return models.Ride{}, err
	}
	isBirthday := foundUser.IsBirthday(rbuc.deterministicTime.Now())

	startAddr := valueobjects.NewAddressVA(args.startAddr.number, args.startAddr.street, args.startAddr.code, args.startAddr.city)
	endAddr := valueobjects.NewAddressVA(args.endAddr.number, args.endAddr.street, args.endAddr.code, args.endAddr.city)
	distance := rbuc.tripScanner.GetTotalDistance(*startAddr, *endAddr)

	trip, err := valueobjects.NewTrip(*startAddr, *endAddr, distance, foundUser.GetForfait(), args.isUberX, isBirthday)
	if err != nil {
		return models.Ride{}, err
	}
	newUuid := rbuc.uuidGenerator.Generate()
	ride := models.BookNewRide(newUuid, foundUser, trip, args.isUberX)
	return ride, nil
}
