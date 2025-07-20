package models

import "strings"

var PRICE_PER_KILOMETER float32 = 0.5

type Trip struct {
	startAddr  string
	endAddr    string
	distance   float32
	totalPrice float32
	forfait    Forfait
}

func NewTrip(startAddr string, endAddr string, distance float32, forfait Forfait) Trip {
	trip := Trip{
		startAddr: startAddr,
		endAddr:   endAddr,
		distance:  distance,
		forfait:   forfait,
	}
	trip.setTotalCost()
	return trip
}

func (fp *Trip) getBasePrice() float32 {
	if strings.Contains(fp.startAddr, "PARIS") {
		if strings.Contains(fp.endAddr, "PARIS") {
			return 30
		}
		return 20
	}
	if strings.Contains(fp.endAddr, "PARIS") {
		return 10
	}
	return 50
}

func (fp *Trip) setTotalCost() {
	basePrice := fp.getBasePrice()
	fp.totalPrice = basePrice + fp.getBasePrice()
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
