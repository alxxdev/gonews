package main

import (
	"log/slog"
	"os"

	"github.com/alxxdev/gonews/config"
	"github.com/alxxdev/gonews/internal/pages"
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	conf := config.NewConfig()

	// Настройка логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Создание Fiber приложения
	app := fiber.New()
	// Подключение slog-fiber как middleware
	app.Use(slogfiber.New(logger))

	pages.NewHandler(app)
	app.Listen(conf.Port)
}
