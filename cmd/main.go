package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"alukart32.com/urlshort/env"
	"alukart32.com/urlshort/internal/config"
	"alukart32.com/urlshort/internal/dal"
	"alukart32.com/urlshort/internal/db"
	"alukart32.com/urlshort/internal/httpx"
	"github.com/google/uuid"
)

var (
	input    = flag.String("input", "yaml", "The input to use between 'yaml/json' files or read from db")
	filepath = flag.String("filepath", "../assets/config.yaml", "The filepath of the yaml/json/... config file")
	dbconn   = flag.String("dbconn", env.GetDBConnCfg(), "Enable the database connection configuration for PostgreSQL")
)

func main() {
	flag.Parse()

	var cfg *config.Config

	switch *input {
	case config.YAML, config.JSON:
		reader, err := config.NewFileConfigReader(*input)
		if err != nil {
			log.Fatal(err)
		}
		cfg, err = config.GetFileConfig(reader, *input, *filepath)
		if err != nil {
			log.Fatal(err)
		}
	case config.DB:
		db, err := db.GetDB(*dbconn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		reader := &config.DBReader{
			RedirectDAO: &dal.Redirect{
				DB: db,
				GenerateUUID: func() string {
					return uuid.New().String()
				},
			},
		}

		cfg, err = config.GetDBConfig(reader)
		if err != nil {
			log.Fatal(err)
		}
	}

	cfg.Print()

	fmt.Println("\ninit a new http handler...")
	mux := httpx.AddHandlerFuncs(http.NewServeMux())
	handler := httpx.GetRedirectHandlerFunc(cfg.PathsToMap(), mux)

	fmt.Println("start http server at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
