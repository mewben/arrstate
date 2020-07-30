package models

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeAddOrLess(t *testing.T) {
	t.Run("It should compute additional charges or deductions correctly", func(t *testing.T) {
		assert := assert.New(t)
		log.Println("assert", assert)
		cases := []struct {
			result     int64
			baseAmount int64
			inputs     []AddOrLessModel
			i          int
		}{
			{
				100,
				1000,
				[]AddOrLessModel{
					{
						Name:     "tax",
						Value:    "10%",
						Less:     false,
						FromBase: false,
					},
				},
				1,
			},
			{
				200,
				1000,
				[]AddOrLessModel{
					{
						Name:     "deposit",
						Value:    "2", // actual value must be 200
						Less:     false,
						FromBase: false,
					},
				},
				2,
			},
			{
				-20,
				1000,
				[]AddOrLessModel{
					{
						Name:     "discount",
						Value:    "2%",
						Less:     true,
						FromBase: false,
					},
				},
				3,
			},
			{
				2100,
				7000,
				[]AddOrLessModel{
					{
						Name:     "tax1",
						Value:    "20%",
						Less:     false,
						FromBase: false,
					},
					{
						Name:     "tax2",
						Value:    "10%",
						Less:     false,
						FromBase: true,
					},
				},
				4,
			},
			{
				2240,
				7000,
				[]AddOrLessModel{
					{
						Name:     "tax1",
						Value:    "20%",
						Less:     false,
						FromBase: false,
					},
					{
						Name:     "tax2",
						Value:    "10%",
						Less:     false,
						FromBase: false,
					},
				},
				5,
			},
			{
				3640,
				7000,
				[]AddOrLessModel{
					{
						Name:     "tax1",
						Value:    "20%", // 7000 * .20 = 1400
						Less:     false,
						FromBase: false,
					},
					{
						Name:     "tax2",
						Value:    "10%",
						Less:     false, // (7000 + 1400) * .10 = 840
						FromBase: false,
					},
					{
						Name:     "tax3",
						Value:    "20%", // 7000 * .20 = 1400
						Less:     false,
						FromBase: true,
					},
				},
				6,
			},
			{
				240,
				7000,
				[]AddOrLessModel{
					{
						Name:     "tax1",
						Value:    "20%", // 7000 * .20 = 1400
						Less:     false,
						FromBase: false,
					},
					{
						Name:     "tax2",
						Value:    "10%",
						Less:     false, // (7000 + 1400) * .10 = 840
						FromBase: false,
					},
					{
						Name:     "Discount",
						Value:    "20", // (7000+1400+840) - 2000 = 7240
						Less:     true,
						FromBase: false,
					},
				},
				6,
			},
			{
				900,
				20000,
				[]AddOrLessModel{
					{
						Name:     "tax1",
						Value:    "10%", // 2000
						Less:     false,
						FromBase: true,
					},
					{
						Name:     "discount",
						Value:    "5%", // 1000
						Less:     true,
						FromBase: false,
					},
				},
				7,
			},
		}

		for _, item := range cases {
			result, err := ComputeAddOrLess(item.baseAmount, item.inputs)
			assert.Nil(err)
			assert.Equal(item.result, result, item.i)
		}
	})
}
