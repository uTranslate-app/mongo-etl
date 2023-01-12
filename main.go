package main

import (
	etl "etl/cmd"
	"log"
)

func main() {
	config, err := etl.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	etl.Start(config)
}
