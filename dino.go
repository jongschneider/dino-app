package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jongschneider/Dino/dynowebportal"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Now listening on port ", config.ServerAddress)
	dynowebportal.RunWebPortal(config.ServerAddress)
}
