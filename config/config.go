package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	log.Println(".env file loaded")
}

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: getSrting("PORT", ""),
	}
}

type LogConfig struct {
	Level slog.Level
}

func NewLogConfig() *LogConfig {
	var l slog.Level
	switch strings.ToUpper(getSrting("LOG_LEVEL", "INFO")) {
	case "DEBUG":
		l = slog.LevelDebug
	case "INFO":
		l = slog.LevelInfo
	case "WARN":
		l = slog.LevelWarn
	default:
		l = slog.LevelInfo
	}
	return &LogConfig{
		Level: l,
	}
}

func getSrting(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}
