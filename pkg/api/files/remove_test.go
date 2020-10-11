package files_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestRemoveFile(t *testing.T) {
	log.Println("-- test file.delete")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/files"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	file1 := CreateFixture(app, token1, testData1)
	file2 := CreateFixture(app, token2, testData1)

	t.Run("It should remove file", func(t *testing.T) {
		assert := assert.New(t)

		req := helpers.DoRequest("DELETE", path+"/"+file1.ID.Hex(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseDelete(res)
		assert.Nil(err)
		assert.Equal(file1.ID.Hex(), response["file"])

	})

	t.Run("It should not allow removing from other business", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("DELETE", path+"/"+file2.ID.Hex(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(services.T(errors.ErrNotFoundFile), response.Message, response)

	})
}
