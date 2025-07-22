package gateways

import "time"

type IDeterministicTime interface {
	Now() time.Time
}
