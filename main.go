package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aaletov/linx-test/pkg/utils"
)

var (
	pathPtr = flag.String("path", "", "Path to data file")
)

func main() {
	flag.Parse()

	ioreader, decoder, err := utils.OpenWithCheck(*pathPtr)
	if err != nil {
		fmt.Printf("Opening error: %v", err)
		os.Exit(1)
	}
	bestProduct, err := decoder(ioreader)
	if err != nil {
		fmt.Printf("Internal file processing error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(bestProduct)
}
