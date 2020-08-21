package config

import "fmt"

type Config struct {
	Postgres *PostgresConfig
	Port     uint16
}

type PostgresConfig struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
}

func (pc *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pc.Host, pc.Port, pc.User, pc.Password, pc.Database)
}
