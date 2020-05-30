package enums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencies(t *testing.T) {
	t.Run("It should load the currencies.json on init", func(t *testing.T) {
		assert := assert.New(t)
		assert.Greater(len(Currencies), 0)
	})
}
