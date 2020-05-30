package enums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountries(t *testing.T) {
	t.Run("It should load the countries.json on init", func(t *testing.T) {
		assert := assert.New(t)
		assert.Greater(len(Countries), 0)
	})
}
