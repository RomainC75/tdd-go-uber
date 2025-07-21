package valueobjects

import "strings"

type Adress struct {
	number int
	street string
	code   int
	city   string
}

func NewAddressVA(number int, street string, code int, city string) *Adress {
	return &Adress{number, street, code, city}
}

func (ava *Adress) IsInParis() bool {
	if strings.Contains(strings.ToLower(ava.city), "paris") {
		return true
	}
	return false
}
