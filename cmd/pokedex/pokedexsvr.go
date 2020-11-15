package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pokedex/internal/config"
	"github.com/Pokedex/internal/service/pokedex"
)

func main() {
	cfg := readConfig()
	// uso m de pokeMon, p esta reservada para Pokedex
	service, _ := pokedex.New(cfg) // inyecto a mi servicio una config
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
