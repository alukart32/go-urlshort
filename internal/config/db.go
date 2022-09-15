package config

import (
	"log"

	"alukart32.com/urlshort/internal/dal"
)

const (
	DB = "db"
)

// DBReader reads configuration for determine http redirects rules from db.
type DBReader struct {
	RedirectDAO dal.RedirectDAO
}

func (p *DBReader) Read() (*Config, error) {
	log.Printf("fetch all redirects from db...")
	entities, err := p.RedirectDAO.FetchAll()
	if err != nil {
		return nil, err
	}

	cfg := new(Config)

	for _, v := range entities {
		cfg.AddRedirect(v.Path, v.Url)
	}
	return cfg, nil
}

func GetDBConfig(reader ConfigReader) (*Config, error) {
	conf, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	return conf, nil
}
