package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubdomain(t *testing.T) {
	t.Run("It should get subdomain from url", func(t *testing.T) {
		assert := assert.New(t)
		cases := []struct {
			In  string
			Out string
		}{
			{
				"http://melvin.example.lh:8000",
				"melvin",
			},
			{
				"http://dean.example.lh",
				"dean",
			},
			{
				"https://theo.example.lh",
				"theo",
			},
			{
				"https://vinia.example.com",
				"vinia",
			},
			{
				"http://google.com",
				"google",
			},
		}

		for _, cas := range cases {
			assert.Equal(cas.Out, GetSubdomain(cas.In))
		}

	})
}
