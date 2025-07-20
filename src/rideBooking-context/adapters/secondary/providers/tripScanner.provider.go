package providers

type FakeTripProvider struct {
	Distance float32
}

func NewFakeTripScannerProvider() *FakeTripProvider {
	return &FakeTripProvider{}
}

func (fp *FakeTripProvider) GetTotalDistance(start string, end string) float32 {
	return fp.Distance
}
