package gateways

import valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

type ITripScanner interface {
	GetTotalDistance(startAddr valueobjects.Adress, endAddr valueobjects.Adress) float32
}
