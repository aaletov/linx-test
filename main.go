package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/aaletov/linx-test/pkg/product"
	"github.com/gocarina/gocsv"
)

func GetBestProductJSON(d *json.Decoder) (product.Product, error) {
	var bestProduct product.Product
	d.Token()
	for d.More() {
		var currProduct product.Product
		err := d.Decode(&currProduct)
		// I suppose that JSON was previously validated and contains correct data.
		// In this case, err != nil means decoder reached closing bracket
		uerr := &json.UnmarshalFieldError{}
		if err != nil {
			if errors.As(err, uerr) {
				break
			} else {
				return product.Product{}, err
			}
		}
		if bestProduct.LessGood(currProduct) {
			bestProduct = currProduct
		}
	}
	return bestProduct, nil
}

func GetBestProductCSV(u *gocsv.Unmarshaller) (product.Product, error) {
	var bestProduct product.Product
	for {
		record, err := u.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return product.Product{}, err
			}
		}
		// Considering file as correct, type assertion with "ok" is not needed
		currProduct := record.(product.Product)
		if bestProduct.LessGood(currProduct) {
			bestProduct = currProduct
		}
	}
	return bestProduct, nil
}

func GetCSVUmarshaller(r io.Reader) (*gocsv.Unmarshaller, error) {
	csvreader := csv.NewReader(r)
	csvreader.Comma = ';'
	csvreader.TrimLeadingSpace = true
	u, err := gocsv.NewUnmarshaller(csvreader, product.Product{})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func main() {
	jsonstring := `[
    {"product": "Печенье", "price": 34, "rating": 3},
    {"product": "Сахар", "price": 45, "rating": 2},
    {"product": "Варенье", "price": 200, "rating": 5}
	]`
	stringreader := strings.NewReader(jsonstring)
	d := json.NewDecoder(stringreader)
	fmt.Println(GetBestProductJSON(d))
	csvstring := `Product;Price;Rating
		Печенье;3;5
		Яблоки;1;2
		Тыква;2;3`
	stringreader = strings.NewReader(csvstring)
	u, _ := GetCSVUmarshaller(stringreader)
	fmt.Println(GetBestProductCSV(u))
	for {
		record, err := u.Read()
		if err != nil {
			break
		}
		if p, ok := record.(product.Product); ok {
			fmt.Println(p)
		}
	}

}
