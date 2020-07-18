package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomValidators(t *testing.T) {
	t.Run("ValidateNumberOrPercentage", func(t *testing.T) {
		assert := assert.New(t)
		cases := []struct {
			output bool
			input  string
		}{
			{
				true,
				"0",
			},
			{
				true,
				"10",
			},
			{
				true,
				"10.5",
			},
			{
				true,
				"5%",
			},
			{
				true,
				"0.5%",
			},
			{
				true,
				"100%",
			},
			{
				true,
				"98.6%",
			},
			{
				true,
				"101",
			},
			{
				false,
				"101%",
			},
			{
				false,
				"noperc",
			},
			{
				false,
				"%",
			},
		}

		for _, item := range cases {
			assert.Equal(item.output, checkNumberOrPercentage(item.input), item.input)
		}

	})
}
