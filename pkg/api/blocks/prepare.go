package blocks

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/utils"
)

// Prepare block -
func (h *Handler) Prepare(data fiber.Map) (BlockI, error) {
	entityTypeI, ok := data["entityType"]
	if !ok {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}
	entityType := entityTypeI.(string)

	blockTypeI, ok := data["type"]
	if !ok {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}
	blockType := blockTypeI.(string)

	// check if blockType is allowed for this entity
	var allowedBlocks []string
	switch entityType {
	case enums.EntityInvoice:
		allowedBlocks = allowedInvoiceBlocks
		break
	default:
		break
	}
	if !utils.Contains(allowedBlocks, blockType) {
		return nil, errors.NewHTTPError(errors.ErrBlockTypeNotAllowed)
	}

	log.Println("before decode")
	// decode and validate blockType
	block, err := h.Decode(data)
	if err != nil {
		return nil, err
	}
	log.Println("after decode", err)
	err = block.Prepare(h.Ctx, h.DB)
	log.Println("after block.prepare", err)
	return block, err

}
