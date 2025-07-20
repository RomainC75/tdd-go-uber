package gateways

type ITrip interface {
	GetBasePrice(startAddr string, endAddr string) float32
	// GetDistance() float32
}
