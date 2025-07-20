package models

type Forfait string

const (
	ForfaitBasic   Forfait = "BASIC"
	ForfaitPremium         = "PREMIUM"
)

type User struct {
	name    string
	forfait Forfait
}

func NewUser(name string, forfait Forfait) *User {
	return &User{
		name:    name,
		forfait: forfait,
	}
}
