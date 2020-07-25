package blocks

import (
	"log"

	"github.com/gofiber/fiber"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
)

// Create block
// blockI is specified before calling this function
func (h *Handler) Create(data fiber.Map) (BlockI, error) {
	block, err := h.Prepare(data)
	if err != nil {
		return nil, err
	}

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollBlocks, block)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}
	r := doc.(fiber.Map)

	block, err = h.Decode(r)
	if err != nil {
		return nil, err
	}

	// block hook
	err = block.AfterCreate(h.Ctx, h.DB)
	if err != nil {
		return nil, err
	}

	return block, nil
}
