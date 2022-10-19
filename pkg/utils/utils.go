package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aaletov/linx-test/pkg/product"
	"github.com/gocarina/gocsv"
)

type DecoderType = func(io.Reader) (product.Product, error)

var (
	Decoders = map[string]DecoderType{
		"csv":  GetBestProductCSV,
		"json": GetBestProductJSON,
	}
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

func OpenWithCheck(path string) (io.Reader, DecoderType, error) {
	var (
		ioreader io.Reader
		decoder  DecoderType
		err      error
	)

	for {
		if path == "" {
			err = errors.New("Path cannot be empty")
			break
		}
		splitPath := strings.Split(path, ".")
		if len(splitPath) == 1 {
			err = errors.New(fmt.Sprintf("No file specified: %v\n", path))
			break
		}
		extension := splitPath[len(splitPath)-1]
		if strings.Contains(extension, "/") {
			err = errors.New(fmt.Sprintf("Incorrect path: %v\n", path))
			break
		}
		var ok bool
		if decoder, ok = Decoders[extension]; !ok {
			err = errors.New(fmt.Sprintf("Not implemented support for: .%v\n", extension))
			break
		}
		ioreader, err = os.Open(path)
		if err != nil {
			err = fmt.Errorf("Incorrect file path: %v\n", err)
			break
		}

		return ioreader, decoder, nil
	}

	return nil, nil, err
}
