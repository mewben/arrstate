package blocks

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
)

// Decode map interface into specific block type
func (h *Handler) Decode(data fiber.Map) (BlockI, error) {
	blockTypeI, ok := data["type"]
	if !ok {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}
	blockType := blockTypeI.(string)

	var block BlockI
	switch blockType {
	case enums.InvoiceBlockItem:
		block = NewInvoiceItemBlock(h.User.ID, h.Business.ID)
		break
	default:
		block = NewBaseBlock(h.User.ID, h.Business.ID)
		break
	}

	m, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, &block)
	return block, err
}
