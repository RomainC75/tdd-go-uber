package gateways

import "github.com/google/uuid"

type IUUIDGenerator interface {
	Generate() uuid.UUID
}
