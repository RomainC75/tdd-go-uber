package ridebooking

import (
	"tdd-go-uber/src/rideBooking-context/business-logic/gateways"
	"tdd-go-uber/src/rideBooking-context/business-logic/models"
)

type TBook struct {
	startAddr string
	endAddr   string
}

type RideBookingUc struct {
	tripScanner gateways.ITripScanner
}

func NewRideBookingUc(tripProvider gateways.ITripScanner) *RideBookingUc {
	return &RideBookingUc{
		tripScanner: tripProvider,
	}
}

func (rbuc *RideBookingUc) Book(args TBook) models.Trip {
	distance := rbuc.tripScanner.GetTotalDistance(args.startAddr, args.endAddr)
	return models.NewTrip(args.startAddr, args.endAddr, distance)
}
