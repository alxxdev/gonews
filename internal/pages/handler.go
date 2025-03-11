package pages

import "github.com/gofiber/fiber/v2"

type PageHandler struct {
	router fiber.Router
}

type Rubric struct {
	Id   int
	Name string
}

// NewHandler
func NewHandler(router fiber.Router) {
	h := &PageHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h *PageHandler) home(c *fiber.Ctx) error {
	rubrics := []Rubric{
		{Id: 1, Name: "#Еда"},
		{Id: 2, Name: "#Животные"},
		{Id: 3, Name: "#Машины"},
		{Id: 4, Name: "#Спорт"},
		{Id: 5, Name: "#Музыка"},
		{Id: 6, Name: "#Технологии"},
		{Id: 7, Name: "#Прочее"},
	}
	return c.Render("home", fiber.Map{"Rubrics": rubrics})
}
