package product

import (
	"strconv"
)

type Product struct {
	Name   string `json:"product" csv:"Product"`
	Price  int
	Rating int
}

func (p Product) String() string {
	return p.Name + " " + strconv.Itoa(p.Price) + " " + strconv.Itoa(p.Rating)
}

func (p Product) LessGood(other Product) bool {
	return (p.Price < other.Price) || ((p.Price == other.Price) && (p.Rating < other.Rating))
}
