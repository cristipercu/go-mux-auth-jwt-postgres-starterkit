package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
  DbUser string
  DbName string
  DbPassword string
  Sslmode string
}

var Envs = initConfig()

func initConfig() Config {
  godotenv.Load()
  return Config{
    DbUser: getEnv("POSTGRES_USER", "postgres"),
    DbName: getEnv("PPOSTGRES_DB", "postgres"),
    DbPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
    Sslmode: getEnv("SSL_MODE", "disable"),
  }
}

func getEnv(key, fallback string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
  if value, ok := os.LookupEnv(key); ok {
    i, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
      return fallback
    }
    return i
  }
  return fallback
}
