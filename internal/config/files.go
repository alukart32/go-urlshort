package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	YAML = "yaml"
	JSON = "json"
)

// NewFileConfigReader create a new NewFileConfigReader using the passed extension type.
func NewFileConfigReader(ext string) (Reader, error) {
	switch ext {
	case YAML:
		return &YAMLReader{
			ReadRecords{
				LogWriter: os.Stdout,
			},
		}, nil
	case JSON:
		return &JSONReader{
			ReadRecords{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("file format '%s' not found", ext)
	}
}

// JSONReader reads JSON config files.
type JSONReader struct {
	ReadRecords
}

func (j *JSONReader) Read() (*Config, error) {
	c := new(Config)

	j.LogWriter.Write([]byte("start decode the json file...\n"))
	err := json.NewDecoder(j.Reader).Decode(c)
	if err != nil {
		return nil, err
	}
	j.LogWriter.Write([]byte("the json file has been decoded...\n"))

	return c, nil
}

// YAMLReader reads YAML config files.
type YAMLReader struct {
	ReadRecords
}

func (y *YAMLReader) Read() (*Config, error) {
	c := new(Config)

	y.LogWriter.Write([]byte("start decode the yaml file...\n"))
	err := yaml.NewDecoder(y.Reader).Decode(c)
	if err != nil {
		return nil, err
	}
	y.LogWriter.Write([]byte("the yaml file has been decoded...\n"))

	return c, nil
}
