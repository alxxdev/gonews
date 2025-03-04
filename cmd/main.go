package main

import (
	"github.com/alxxdev/gonews/config"
	"github.com/alxxdev/gonews/internal/pages"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.NewConfig()
	app := fiber.New()
	pages.NewHandler(app)
	app.Listen(conf.Port)
}
