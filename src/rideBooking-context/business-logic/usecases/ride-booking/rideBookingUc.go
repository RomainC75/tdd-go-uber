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
	riderId   uuid.UUID
	startAddr TAdressInput
	endAddr   TAdressInput
	isUberX   bool
}

type BookOptions struct {
	isUberX bool
}

type RideBookingUc struct {
	rideRpo           gateways.IRideRepo
	tripScanner       gateways.ITripScanner
	riderRepo         gateways.RiderRepo
	uuidGenerator     gateways.IUUIDGenerator
	deterministicTime gateways.IDeterministicTime
}

func NewRideBookingUc(
	rideRepo gateways.IRideRepo,
	riderRepo gateways.RiderRepo,
	tripProvider gateways.ITripScanner,
	uuidGenerator gateways.IUUIDGenerator,
	deterministicTime gateways.IDeterministicTime,
) *RideBookingUc {
	return &RideBookingUc{
		rideRpo:           rideRepo,
		riderRepo:         riderRepo,
		tripScanner:       tripProvider,
		uuidGenerator:     uuidGenerator,
		deterministicTime: deterministicTime,
	}
}

func (rbuc *RideBookingUc) Book(args TBook) (models.Ride, error) {
	foundRider, err := rbuc.riderRepo.GetRider(args.riderId)
	if err != nil {
		return models.Ride{}, err
	}
	isBirthday := foundRider.IsBirthday(rbuc.deterministicTime.Now())

	startAddr := valueobjects.NewAddressVA(args.startAddr.number, args.startAddr.street, args.startAddr.code, args.startAddr.city)
	endAddr := valueobjects.NewAddressVA(args.endAddr.number, args.endAddr.street, args.endAddr.code, args.endAddr.city)
	distance := rbuc.tripScanner.GetTotalDistance(*startAddr, *endAddr)

	trip, err := valueobjects.NewTrip(*startAddr, *endAddr, distance, foundRider.GetForfait(), args.isUberX, isBirthday)
	if err != nil {
		return models.Ride{}, err
	}

	newUuid := rbuc.uuidGenerator.Generate()
	ride := models.BookNewRide(newUuid, foundRider, trip, args.isUberX, rbuc.deterministicTime.Now())
	err = rbuc.rideRpo.Save(ride)
	if err != nil {
		return models.Ride{}, err
	}
	return ride, nil
}
