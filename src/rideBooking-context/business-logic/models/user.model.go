package models

import "github.com/google/uuid"

type Forfait string

const (
	ForfaitBasic   Forfait = "BASIC"
	ForfaitPremium         = "PREMIUM"
)

type User struct {
	id      uuid.UUID
	name    string
	forfait Forfait
}

func NewUser(id uuid.UUID, name string, forfait Forfait) *User {
	return &User{
		id:      id,
		name:    name,
		forfait: forfait,
	}
}

func (u *User) GetForfait() Forfait {
	return u.forfait
}
