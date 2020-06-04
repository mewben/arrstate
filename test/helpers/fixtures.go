package helpers

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/auth"
	"github.com/mewben/realty278/pkg/models"
)

// CleanupFixture -
func CleanupFixture(db *mongo.Database) {
	collections := []string{
		enums.CollBusinesses,
		enums.CollUsers,
		enums.CollPeople,
		enums.CollProjects,
	}
	for _, col := range collections {
		db.Collection(col).DeleteMany(context.Background(), bson.D{{}})
	}
}

// SignupFixture -
func SignupFixture(app *fiber.App) (*auth.SignupPayload, *models.AuthSuccessResponse) {
	payload := &auth.SignupPayload{
		GivenName:  "testgn",
		FamilyName: "testfn",
		Business:   "Test Business",
		Domain:     "test-domain",
		Email:      "test@email.com",
		Password:   "password",
	}
	req := DoRequest("POST", "/auth/signup", payload, "")
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test signup", err)
	}

	response, err := GetResponseAuth(res)
	if err != nil {
		log.Fatalln("err get response auth", err)
	}

	return payload, response
}
