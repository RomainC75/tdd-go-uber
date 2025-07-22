package valueobjects

import (
	"errors"
)

var PRICE_PER_KILOMETER float32 = 0.5

// value Object
type Trip struct {
	startAddr  Adress
	endAddr    Adress
	distance   float32
	totalPrice float32
}

func NewTrip(startAddr Adress, endAddr Adress, distance float32, forfait Forfait, isUberX bool, isBirthday bool) (Trip, error) {
	if forfait == ForfaitPremium && distance < 3 {
		return Trip{}, errors.New("distance cannot be < 3 when uberX")
	}
	trip := Trip{
		startAddr: startAddr,
		endAddr:   endAddr,
		distance:  distance,
	}
	trip.setTotalCost(forfait, isUberX, isBirthday)
	return trip, nil
}

func (fp *Trip) getBasePrice(isUberX bool, isBirthday bool) float32 {
	var basePrice float32 = 0
	if isUberX && !isBirthday {
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

func (fp *Trip) setTotalCost(forfait Forfait, isUberX bool, isBirthday bool) {
	basePrice := fp.getBasePrice(isUberX, isBirthday)
	fp.totalPrice = basePrice + fp.getDistancePrice(forfait)
}

func (fp *Trip) getDistancePrice(forfait Forfait) float32 {
	if forfait == ForfaitPremium {
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
