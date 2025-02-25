package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	AuthServiceHost string
	AuthServicePort string

	CoreServiceHost string
	CoreServicePort string

	FinanceServiceHost string
	FinanceServicePort string

	LogLevel string
	HttpPort string

	JWTSigningKey string

	RedisHost string
	RedisPort string
}

func Load() Config {
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "staging"))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "8000"))
	config.AuthServiceHost = cast.ToString(getOrReturnDefault("AUTH_SERVICE_HOST", "localhost"))
	config.AuthServicePort = cast.ToString(getOrReturnDefault("AUTH_SERVICE_PORT", 8080))
	config.CoreServiceHost = cast.ToString(getOrReturnDefault("CORE_SERVICE_HOST", "localhost"))
	config.CoreServicePort = cast.ToString(getOrReturnDefault("CORE_SERVICE_PORT", 8080))
	config.FinanceServiceHost = cast.ToString(getOrReturnDefault("FINANCE_SERVICE_HOST", "localhost"))
	config.FinanceServicePort = cast.ToString(getOrReturnDefault("FINANCE_SERVICE_PORT", 8080))
	config.JWTSigningKey = cast.ToString(getOrReturnDefault("JWT_SIGNING_KEY", "secret"))
	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
