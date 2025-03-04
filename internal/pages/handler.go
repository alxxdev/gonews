package pages

import "github.com/gofiber/fiber/v2"

type PageHandler struct {
	router fiber.Router
}

// NewHandler
func NewHandler(router fiber.Router) {
	h := &PageHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h *PageHandler) home(c *fiber.Ctx) error {
	return c.SendString("Homepage")
}
