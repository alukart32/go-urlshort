package env

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dev"
)

// GetDBConnCfg returns default db connection config.
func GetDBConnCfg() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
