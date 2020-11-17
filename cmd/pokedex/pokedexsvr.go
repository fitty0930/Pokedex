package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pokedex/internal/config"
	"github.com/Pokedex/internal/database"
	"github.com/Pokedex/internal/service/pokedex"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()
	// uso m de pokeMon, p esta reservada para Pokedex

	db, err := database.NewDatabase(cfg) // instanciado de la database
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := pokedex.New(db, cfg)               // inyecto a mi servicio una config y 1 db
	httpService := pokedex.NewHTTPTransport(service) // servicio http que usa mi servicio
	r := gin.Default()
	httpService.Register(r)
	r.Run()

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

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS pokedex (
		id integer primary key,
		name varchar);`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
