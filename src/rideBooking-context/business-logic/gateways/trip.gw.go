package gateways

type ITrip interface {
	GetPriceTotalPrice(start string, end string) float32
}
