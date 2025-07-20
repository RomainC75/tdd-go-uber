package providers

import "strings"

var PRICE_PER_KILO float32 = 0.5

type FakeProvider struct {
	Distance float32
}

func NewFakeProvider() *FakeProvider {
	return &FakeProvider{}
}

func (fp *FakeProvider) GetPriceTotalPrice(startAddr string, endAddr string) float32 {
	basicPrice := fp.getBasePrice(startAddr, endAddr)
	distancePrice := fp.getDistancePrice(startAddr, endAddr)
	return basicPrice + distancePrice
}

func (fp *FakeProvider) getBasePrice(start string, end string) float32 {
	if strings.Contains(start, "PARIS") {
		if strings.Contains(end, "PARIS") {
			return 30
		}
		return 20
	}
	if strings.Contains(end, "PARIS") {
		return 10
	}
	return 50
}

func (fp *FakeProvider) getDistancePrice(start string, end string) float32 {
	return fp.Distance * PRICE_PER_KILO
}
