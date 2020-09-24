package blocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/utils"
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

	// decode and validate blockType
	block, err := h.Decode(data)
	if err != nil {
		return nil, err
	}
	err = block.Prepare(h.Ctx, h.DB)
	return block, err

}
