package providers

import "github.com/google/uuid"

type FakeUuidGenerator struct {
	expectedUuid uuid.UUID
}

func NewFakeUuidGenerator(expectedUuid uuid.UUID) *FakeUuidGenerator {
	return &FakeUuidGenerator{
		expectedUuid: expectedUuid,
	}
}

func (fug *FakeUuidGenerator) Generate() uuid.UUID {
	return fug.expectedUuid
}
