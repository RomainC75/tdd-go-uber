package ridebooking

import "tdd-go-uber/src/rideBooking-context/business-logic/gateways"

type TBook struct {
	startAddr string
	endAddr   string
}

type RideBookingUc struct {
	trip gateways.ITrip
}

func NewRideBookingUc(tripProvider gateways.ITrip) *RideBookingUc {
	return &RideBookingUc{
		trip: tripProvider,
	}
}

func (rbuc *RideBookingUc) Book(args TBook) float32 {
	return rbuc.trip.GetBasePrice(args.startAddr, args.endAddr)
}
