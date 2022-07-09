package main

import (
	"Aoi/cmd"
	"log"
)

var name string

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

}
