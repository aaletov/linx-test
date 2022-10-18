package main

import (
	"encoding/json"

	"github.com/aaletov/linx-test/pkg/product"
)

func GetBestProductJSON(d *json.Decoder) {
	var bestProduct product.Product
	d.Token()
	for d.More() {
		var product product.Product
		err := d.Decode(&product)
		// I suppose that JSON was previously validated and contains correct data.
		// In this case, err != nil means decoder reached closing bracket
		if err != nil {
			break
		}
		if bestProduct.LessGood(product) {
			bestProduct = product
		}
	}
}

func main() {
	_ = `[
    {"product": "Печенье", "price": 34, "rating": 3},
    {"product": "Сахар", "price": 45, "rating": 2},
    {"product": "Варенье", "price": 200, "rating": 5}
	]`
}
