package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"

	"github.com/aaletov/linx-test/pkg/product"
	"github.com/gocarina/gocsv"
)

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

func GetBestProductJSON(r io.Reader) (product.Product, error) {
	d := json.NewDecoder(r)
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

func GetBestProductCSV(r io.Reader) (product.Product, error) {
	u, _ := GetCSVUmarshaller(r)
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
