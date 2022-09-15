package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	YAML = "yaml"
	JSON = "json"
)

// FileConfigReader combines the roles of Reader of records from config files.
type FileConfigReader struct {
	Reader io.Reader
}

func (r *FileConfigReader) SetReader(reader io.Reader) {
	r.Reader = reader
}

// NewFileConfigReader create a new NewFileConfigReader using the passed extension type.
func NewFileConfigReader(ext string) (ConfigReader, error) {
	switch ext {
	case YAML:
		return &YAMLReader{
			FileConfigReader{},
		}, nil
	case JSON:
		return &JSONReader{
			FileConfigReader{},
		}, nil
	default:
		return nil, fmt.Errorf("file format '%s' not found", ext)
	}
}

// JSONReader reads JSON config files.
type JSONReader struct {
	FileConfigReader
}

func (j *JSONReader) Read() (*Config, error) {
	c := new(Config)

	log.Println("start decode the json file...")
	err := json.NewDecoder(j.Reader).Decode(c)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("the json file has been decoded...")

	return c, nil
}

// YAMLReader reads YAML config files.
type YAMLReader struct {
	FileConfigReader
}

func (y *YAMLReader) Read() (*Config, error) {
	c := new(Config)

	log.Println("start decode the yaml file...")
	err := yaml.NewDecoder(y.Reader).Decode(c)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("the yaml file has been decoded...")

	return c, nil
}

// GetConfig gets config from config.Reader.
// The method of getting the config depends on the inType parameter: yaml/json.
// If there is an error, it will return a nil Config value and error description.
func GetFileConfig(reader ConfigReader, inType string, fp string) (*Config, error) {
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
		reader.(*YAMLReader).SetReader(f)
	case JSON:
		if ext != ".json" {
			return nil, fmt.Errorf("for input type 'json' config file has a wrong extension: %s", fp)
		}

		f, err := os.Open(fp)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		reader.(*JSONReader).SetReader(f)
	}

	conf, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	return conf, nil
}
