package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/aaletov/linx-test/pkg/product"
	"github.com/aaletov/linx-test/pkg/utils"
)

var (
	pathPtr = flag.String("path", "", "Path to data file")
)

func main() {
	jsonstring := `[
    {"product": "Печенье", "price": 34, "rating": 3},
    {"product": "Сахар", "price": 45, "rating": 2},
    {"product": "Варенье", "price": 200, "rating": 5}
	]`
	stringreader := strings.NewReader(jsonstring)
	d := json.NewDecoder(stringreader)
	fmt.Println(utils.GetBestProductJSON(d))
	csvstring := `Product;Price;Rating
		Печенье;3;5
		Яблоки;1;2
		Тыква;2;3`
	stringreader = strings.NewReader(csvstring)
	u, _ := utils.GetCSVUmarshaller(stringreader)
	fmt.Println(utils.GetBestProductCSV(u))
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
