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
	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}
	fmt.Printf("opts: %v", config)
	// fmt.Println("Hello Fucking World")

}
