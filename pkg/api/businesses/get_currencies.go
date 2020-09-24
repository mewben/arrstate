package businesses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/internal/enums"
)

// GetCurrencies -
func (h *Handler) GetCurrencies() (fiber.Map, error) {
	return fiber.Map{
		"currencies": enums.Currencies,
	}, nil
}
