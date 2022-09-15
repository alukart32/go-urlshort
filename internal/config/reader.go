package config

type ConfigReader interface {
	Read() (*Config, error)
}
