package main

import (
	"log"

	"github.com/joaoalmeidarx/jra-go-pub/config"
)

func main() {
	err := config.LoadDefault("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Appname :", config.Default.Appname)
	log.Println("Version :", config.Default.Version)
	log.Println("Hostname:", config.Default.Hostname)
	log.Println("Port    :", config.Default.Port)
}
