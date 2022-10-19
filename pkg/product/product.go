package product

type Product struct {
	Name   string `json:"product" csv:"Product"`
	Price  int
	Rating int
}

func (p Product) LessGood(other Product) bool {
	return (p.Price < other.Price) || ((p.Price == other.Price) && (p.Rating < other.Rating))
}
