package main

import (
	"fmt"
	"log"
	"projects/projector-go/src/projector-go/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}
	fmt.Printf("opts: %v", opts)
	// fmt.Println("Hello Fucking World")
}
