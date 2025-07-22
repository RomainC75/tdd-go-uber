package providers

import "time"

type FakeDeterministicTime struct {
	ExpectedTime time.Time
}

func NewDeterministicTime() *FakeDeterministicTime {
	return &FakeDeterministicTime{}
}

func (fdt *FakeDeterministicTime) Now() time.Time {
	return fdt.ExpectedTime
}
