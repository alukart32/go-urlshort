package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetConfig gets config from config.Reader.
// The method of getting the config depends on the inType parameter: yaml/json.
// If there is an error, it will return a nil Config value and error description.
func GetConfig(reader Reader, inType string, fp string) (*Config, error) {
	ext := filepath.Ext(fp)
	switch inType {
	case YAML:
		if ext != ".yaml" {
			return nil, fmt.Errorf("for input type 'yaml' config file has a wrong extension: %s", fp)
		}

		f, err := os.Open(fp)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		reader.SetReader(f)
	case JSON:
		if ext != ".json" {
			return nil, fmt.Errorf("for input type 'json' config file has a wrong extension: %s", fp)
		}

		f, err := os.Open(fp)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		reader.SetReader(f)
	}

	conf, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	return conf, nil
}
