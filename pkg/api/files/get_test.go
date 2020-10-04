package files_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/test/helpers"
)

func TestGetFiles(t *testing.T) {
	log.Println("-- test files.get --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/files"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project1 := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 0)
	CreateFixture(app, token1, testData1) // global1
	CreateFixture(app, token2, testData1) // global2

	t.Run("It should get the list of files by entityType and entityID", func(t *testing.T) {
		// setup
		assert := assert.New(t)
		file1 := testData1
		file1["entityType"] = enums.EntityProject
		file1["entityID"] = project1.ID
		CreateFixture(app, token1, file1)
		file2 := testData2
		file2["entityType"] = enums.EntityProject
		file2["entityID"] = project1.ID
		CreateFixture(app, token1, file2)
		file3 := testData1
		file3["entityType"] = enums.EntityProject
		file3["entityID"] = project2.ID
		CreateFixture(app, token1, file3)

		query := url.Values{}
		query.Add("entityType", enums.EntityProject)
		query.Add("entityID", project1.ID.Hex())
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := GetResponses(res)
		assert.Nil(err)
		assert.Len(response.Data, 2)
		assert.Equal(2, response.Total)
	})

	t.Run("It should get list of global files", func(t *testing.T) {
		assert := assert.New(t)

		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := GetResponses(res)
		assert.Nil(err)
		assert.Len(response.Data, 4)
		assert.Equal(4, response.Total)

	})
}
