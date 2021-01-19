package main

import (
	"./wayback"
	"flag"
	"fmt"
	"log"
)

var (
	raw = flag.Bool("raw", false, "Image URls when the target is an image")
)

func main() {
	flag.Parse()
	for _, u := range flag.Args() {
		closest, err := wayback.FindClosestURL(u, *raw)
		if err != nil {
			log.Fatalf("Error finding closest URL for %s: %v", u, err)
		}
		fmt.Println(closest)
	}
}
