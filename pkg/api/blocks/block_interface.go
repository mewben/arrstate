package blocks

import (
	"context"
	"encoding/json"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services/database"
	"github.com/mewben/realty278/pkg/utils"
)

// BlockI -
type BlockI interface {
	Prepare(ctx context.Context, db *database.Service) error
	AfterCreate(ctx context.Context, db *database.Service) error
}

// BaseBlock -
type BaseBlock struct {
	models.BlockModel `bson:",inline"`
}

// InvoiceItemBlock -
type InvoiceItemBlock struct {
	models.InvoiceItemBlockModel `bson:",inline"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// CheckEntity -
func CheckEntity(ctx context.Context, db *database.Service, entityType string, entityID, businessID primitive.ObjectID) error {
	collectionName := utils.GetCollectionName(entityType)
	if collectionName == "" {
		return errors.NewHTTPError(errors.ErrNotFound)
	}
	// check entityID
	_, err := db.FindByID(ctx, collectionName, entityID, businessID)
	if err != nil {
		return err
	}
	return nil
}

// Prepare generic block
func (block *BaseBlock) Prepare(ctx context.Context, db *database.Service) error {
	if err := CheckEntity(ctx, db, block.EntityType, block.EntityID, block.BusinessID); err != nil {
		return err
	}
	return nil
}

// Prepare -
func (block *InvoiceItemBlock) Prepare(ctx context.Context, db *database.Service) error {
	log.Println("preaprea invoice itemblock")
	var err error
	if err = CheckEntity(ctx, db, block.EntityType, block.EntityID, block.BusinessID); err != nil {
		return err
	}
	log.Println("after checkentity")

	// validate
	validate.RegisterValidation("numberOrPercentage", utils.ValidateNumberOrPercentage)
	if err = validate.Struct(block); err != nil {
		log.Println("errr---", err)
		return errors.NewHTTPError(errors.ErrInputInvalid, err)
	}
	log.Println("after validatestruct")

	// defaults
	if block.Quantity == 0 {
		block.Quantity = 1
	}
	block.TaxAmount, block.DiscountAmount, block.Total, err = utils.CalculateItem(block.Amount, block.Quantity, block.Tax, block.Discount)
	if err != nil {
		return err
	}

	return nil
}

// AfterCreate -
func (block *BaseBlock) AfterCreate(ctx context.Context, db *database.Service) error {
	log.Println("after create baseblock")
	collectionName := utils.GetCollectionName(block.EntityType)
	if collectionName == "" {
		return errors.NewHTTPError(errors.ErrNotFound)
	}
	upd := bson.D{
		{
			Key: "$push",
			Value: bson.D{
				{
					Key:   "blocks",
					Value: block.ID,
				},
			},
		},
	}
	_, err := db.FindByIDAndUpdate(ctx, collectionName, block.EntityID, upd)
	if err != nil {
		return err
	}

	return nil
}

// AfterCreate -
func (block *InvoiceItemBlock) AfterCreate(ctx context.Context, db *database.Service) error {
	log.Println("after create invoiceitemblock")
	collectionName := utils.GetCollectionName(block.EntityType)
	if collectionName == "" {
		return errors.NewHTTPError(errors.ErrNotFound)
	}

	// Calculate the amounts here
	// 1. Get the invoice doc here
	doc, err := db.FindByID(ctx, collectionName, block.EntityID, block.BusinessID)
	if err != nil {
		return err
	}
	invoice := doc.(*models.InvoiceModel)

	// 2. Get the blocks
	filter := bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: invoice.Blocks,
				},
			},
		},
	}
	blocksCursor, err := db.Find(ctx, enums.CollBlocks, filter)
	if err != nil {
		return err
	}
	blocks := make([]fiber.Map, 0)
	if err = blocksCursor.All(ctx, &blocks); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(blocks, "", "  ")
	log.Println("currentblocks: ", string(j))

	j, _ = json.MarshalIndent(block, "", "  ")
	log.Println("thisblock: ", string(j))

	// 3. calculate total invoice amounts by looping through the blocks type=item
	subTotal := block.Total
	for _, block := range blocks {
		if block["type"].(string) == enums.InvoiceBlockItem {
			subTotal += block["total"].(float64)
		}
	}

	log.Println("subTotal", subTotal)

	taxAmount, discountAmount, total, err := utils.CalculateItem(subTotal, 1, invoice.Tax, invoice.Discount)
	if err != nil {
		return err
	}

	set := fiber.Map{
		"subTotal":       subTotal,
		"taxAmount":      taxAmount,
		"discountAmount": discountAmount,
		"total":          total,
	}

	upd := bson.D{
		{
			Key:   "$set",
			Value: set,
		},
		{
			Key: "$push",
			Value: bson.D{
				{
					Key:   "blocks",
					Value: block.ID,
				},
			},
		},
	}
	_, err = db.FindByIDAndUpdate(ctx, collectionName, block.EntityID, upd)
	if err != nil {
		return err
	}

	return nil
}
