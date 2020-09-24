package reports

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services/database"
	"github.com/mewben/arrstate/pkg/utils"
)

// Handler -
type Handler struct {
	DB       *database.Service
	Ctx      context.Context
	User     *models.UserModel
	Business *models.BusinessModel
}

// ResponseList -
type ResponseList struct {
	Total int                    `json:"total"`
	Data  []*models.InvoiceModel `json:"list"`
}

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/reports/income", func(c *fiber.Ctx) error {
		log.Println("get income reports")
		var err error
		h.Ctx = c.Context()
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			return c.Status(400).JSON(err)
		}

		response, err := h.GetIncome(c.Query("from"), c.Query("to"))
		if err != nil {
			log.Println("errrrrr", err)
			return c.Status(400).JSON(err)
		}
		return c.Status(200).JSON(response)

	})
}
