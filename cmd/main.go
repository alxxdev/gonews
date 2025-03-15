package main

import (
	"log/slog"
	"os"

	"github.com/alxxdev/gonews/config"
	"github.com/alxxdev/gonews/internal/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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
	// HTML template engine
	engine := html.New("./html", ".html")
	// Создание Fiber приложения
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	pages.NewHandler(app)
	app.Listen(conf.Port)
}
