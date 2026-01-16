package config

import (
  "os"
  "fmt"
  "github.com/joho/godotenv"
)

type Config struct {
  port      string
  dbUrl     string
  JWTSecret string
}

func LoadConfig() *Config {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Warning: env file not found.\nUsing system environment variables.")
  }
  return &Config {
    port: getEnv("PORT", "8080"),
    dbUrl: getEnv("DB_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
    JWTSecret: getEnv("JWT_SECRET", "your-secret-key")
  }
}

func getEnv(key, defaultValue string) {
  if value := os.Getenv(key); value != "" {
    return value
  }
  return defaultValue
}
