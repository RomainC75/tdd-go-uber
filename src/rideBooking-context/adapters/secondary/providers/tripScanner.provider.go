package providers

import valueobjects "tdd-go-uber/src/rideBooking-context/business-logic/valueObjects"

type FakeTripProvider struct {
	Distance float32
}

func NewFakeTripScannerProvider() *FakeTripProvider {
	return &FakeTripProvider{}
}

func (fp *FakeTripProvider) GetTotalDistance(startAddr valueobjects.Adress, endAddr valueobjects.Adress) float32 {
	return fp.Distance
}
