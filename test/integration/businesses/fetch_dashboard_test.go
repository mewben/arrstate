package businesses

import (
	"log"
	"os"
	"testing"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestFetchDashboard(t *testing.T) {
	log.Println("-- BUSINESS.DASHBOARD --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	project1 := helpers.ProjectFixture(app, token1, 0)
	helpers.ProjectFixture(app, token1, 1)
	helpers.PropertyFixture(app, token1, nil, 0)
	helpers.PropertyFixture(app, token1, nil, 0)
	helpers.PropertyFixture(app, token1, &project1.ID, 0)
	helpers.PersonFixture(app, token1, 0)

	t.Run("It should refresh data to dashbooard", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", "/api/dashboard", nil, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "business")
		assert.Nil(err)
		response := ress.(*models.BusinessModel)

		assert.NotZero(response.Dashboard, response.Dashboard)
		assert.Equal(response.Dashboard["projects"], models.DashboardModel{
			Total: 2,
			Label: "Projects",
		})
		assert.Equal(response.Dashboard["properties"], models.DashboardModel{
			Total: 3,
			Label: "Properties",
		})
		assert.Equal(response.Dashboard["people"], models.DashboardModel{
			Total: 2,
			Label: "People",
		})
	})
}
