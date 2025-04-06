package home

import (
	"github.com/alxxdev/gonews/pkg/tadapter"
	"github.com/alxxdev/gonews/views/components"
	"github.com/alxxdev/gonews/views/pages"
	"github.com/gofiber/fiber/v2"
)

type PageHandler struct {
	router fiber.Router
}

var mockData = pages.MainProps{
	Tags: []components.Tag{
		{ID: 1, Name: "#Еда", Enum: "food"},
		{ID: 2, Name: "#Животные", Enum: "animals"},
		{ID: 3, Name: "#Машины", Enum: "cars"},
		{ID: 4, Name: "#Спорт", Enum: "sport"},
		{ID: 5, Name: "#Музыка", Enum: "music"},
		{ID: 6, Name: "#Технологии", Enum: "tech"},
		{ID: 7, Name: "#Прочее", Enum: "other"},
	},
}

func NewPageHandler(router fiber.Router) {
	h := &PageHandler{
		router: router,
	}
	h.router.Get("/", h.main)
}

func (h *PageHandler) main(c *fiber.Ctx) error {
	return tadapter.Render(c, pages.Main(mockData))
}
