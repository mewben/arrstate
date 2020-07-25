package businesses

import (
	"github.com/gofiber/fiber"
	"github.com/mewben/arrstate/internal/enums"
)

// GetCountries -
func (h *Handler) GetCountries() (fiber.Map, error) {
	return fiber.Map{
		"countries": enums.Countries,
	}, nil
}
