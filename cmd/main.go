package main

import (
	"log/slog"
	"os"

	"github.com/alxxdev/gonews/config"
	"github.com/alxxdev/gonews/internal/home"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	conf := config.NewConfig()
	// Logger
	logConf := config.NewLogConfig()
	opts := &slog.HandlerOptions{
		Level: logConf.Level,
	}
	slogHandler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(slogHandler)
	slog.SetDefault(logger)
	// Fiber
	app := fiber.New()
	app.Static("/public", "public")
	home.NewPageHandler(app)
	app.Listen(conf.Port)
}
