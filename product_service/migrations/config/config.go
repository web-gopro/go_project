package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   string = "debug"
	TestMode    string = "test"
	ReleaseMode string = "release"
)

type GeneralConfig struct {
	AppName     string
	Environment string
	Version     string
	HTTPPort    string
	HTTPScheme  string
	SignInKey   string
}

type PgConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

type Config struct {
	GeneralConfig GeneralConfig
	PgConfig      PgConfig
}

func NewConfig() Config {
	return Config{GeneralConfig: GeneralConfig{AppName: "gitlab.com/e-market724/back-end/product_service"}}
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	var config = NewConfig()

	// General config
	config.GeneralConfig.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.GeneralConfig.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))
	config.GeneralConfig.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.GeneralConfig.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))
	config.GeneralConfig.SignInKey = cast.ToString(getOrReturnDefaultValue("SIGN_IN_KEY", "ASJDKLFJASasdFASE2SD2dafa"))

	// Postgres config
	config.PgConfig.Username = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "superuser"))
	config.PgConfig.Password = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "superuser"))
	config.PgConfig.Host = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PgConfig.Port = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PgConfig.DatabaseName = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.GeneralConfig.AppName))

	return config

}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
