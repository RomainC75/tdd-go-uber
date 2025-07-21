package models

import (
	valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"
)

var PRICE_PER_KILOMETER float32 = 0.5

type Trip struct {
	startAddr  valueobjects.Adress
	endAddr    valueobjects.Adress
	distance   float32
	totalPrice float32
	forfait    Forfait
	isUberX    bool
}

func NewTrip(startAddr valueobjects.Adress, endAddr valueobjects.Adress, distance float32, forfait Forfait, isUberX bool) Trip {
	trip := Trip{
		startAddr: startAddr,
		endAddr:   endAddr,
		distance:  distance,
		forfait:   forfait,
		isUberX:   isUberX,
	}
	trip.setTotalCost()
	return trip
}

func (fp *Trip) getBasePrice() float32 {
	var basePrice float32 = 0
	if fp.isUberX {
		basePrice = 10
	}
	if fp.startAddr.IsInParis() {
		if fp.endAddr.IsInParis() {
			return basePrice + 30
		}
		return basePrice + 20
	}
	if fp.endAddr.IsInParis() {
		return basePrice + 10
	}
	return basePrice + 50
}

func (fp *Trip) setTotalCost() {
	basePrice := fp.getBasePrice()
	fp.totalPrice = basePrice + fp.getDistancePrice()
}

func (fp *Trip) getDistancePrice() float32 {
	if fp.forfait == ForfaitPremium {
		if fp.distance < 5 {
			return 0
		}
		return (fp.distance - 5) * PRICE_PER_KILOMETER
	}
	return fp.distance * PRICE_PER_KILOMETER
}

func (fp *Trip) GetTotalPrice() float32 {
	return fp.totalPrice
}
