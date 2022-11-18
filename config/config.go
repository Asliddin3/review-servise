package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment         string
	PostgresHost        string
	PostgresPort        int
	PostgresDatabase    string
	PostgresUser        string
	PostgresPassword    string
	CustomerServiceHost string
	CustomerServicePort int
	LogLevel            string
	RPCPort             string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DB", "review"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "asliddin2001"))
	c.CustomerServiceHost = cast.ToString(getOrReturnDefault("CUSTOMER_SERVICE_HOST", "customer-servis"))
	c.CustomerServicePort = cast.ToInt(getOrReturnDefault("CUSTOMER_SEVICE_PORT", 8810))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8840"))
	return c
}

func getOrReturnDefault(key string, defaulValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaulValue
}
