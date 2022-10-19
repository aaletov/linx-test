package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aaletov/linx-test/pkg/product"
	"github.com/aaletov/linx-test/pkg/utils"
)

var (
	pathPtr  = flag.String("path", "", "Path to data file")
	decoders = map[string]func(io.Reader) (product.Product, error){
		"csv":  utils.GetBestProductCSV,
		"json": utils.GetBestProductJSON,
	}
)

func main() {
	flag.Parse()

	splitPath := strings.Split(*pathPtr, ".")
	fmt.Println(*pathPtr)
	extension := splitPath[len(splitPath)-1]
	if _, ok := decoders[extension]; !ok {
		fmt.Printf("Not implemented support for: .%v\n", extension)
		os.Exit(1)
	}

	ioreader, err := os.Open(*pathPtr)
	if err != nil {
		fmt.Printf("Incorrect file path: %v\n", *pathPtr)
		os.Exit(1)
	}

	bestProduct, err := decoders[extension](ioreader)
	if err != nil {
		fmt.Printf("Internal file processing error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(bestProduct)
}
