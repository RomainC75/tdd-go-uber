package gateways

type ITripScanner interface {
	GetTotalDistance(start string, end string) float32
}
