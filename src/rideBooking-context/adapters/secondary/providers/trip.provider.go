package providers

import "strings"

type FakeProvider struct{}

func NewFakeProvider() *FakeProvider {
	return &FakeProvider{}
}

func (fp *FakeProvider) GetBasePrice(start string, end string) float32 {
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
