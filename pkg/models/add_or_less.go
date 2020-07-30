package models

import (
	"strconv"
	"strings"

	"github.com/mewben/arrstate/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddOrLessModel -
// UseCase:
// Tax:
// 		Name: VAT					PWD Discount
//		Less: false				true
//		Value: 10%				1000
//		Amount: 1500			-1000
//		FromBase: true		false
type AddOrLessModel struct {
	ID    *primitive.ObjectID `bson:"_id" json:"_id"`
	Name  string              `bson:"name" json:"name"`
	Less  bool                `bson:"less" json:"less"`
	Value string              `bson:"value" json:"value" validate:"numberOrPercentage"`
	// Amount int64              `bson:"amount" json:"amount"`
	// FromBase sets which to calculate from. Base or preceding amount
	FromBase bool `bson:"fromBase" json:"fromBase"`
}

// ComputeAddOrLess -
// computes the additional charges or deductions
func ComputeAddOrLess(baseAmount int64, inputs []AddOrLessModel) (int64, error) {
	var result int64
	var accumulator int64
	accumulator = baseAmount

	for _, item := range inputs {
		// Parse Value
		v, isPercent, err := ParseNumberOrPercentage(item.Value)
		if err != nil {
			return result, err
		}

		// get amount to be added/deducted
		if item.FromBase {
			if isPercent {
				v = baseAmount * v / 10000
			}
		} else {
			if isPercent {
				v = accumulator * v / 10000
			}
		}

		if item.Less {
			result -= v
		} else {
			result += v
		}

		accumulator = baseAmount + result
	}

	return result, nil
}

// ParseNumberOrPercentage -
// returns the number part
// ex. 10% = 1000, true
// 15.35% = 1535, true
// 200 = 20000, false
func ParseNumberOrPercentage(input string) (int64, bool, error) {
	sp := strings.Split(input, "%")
	v := input
	var isPercent bool

	if len(sp) == 2 {
		v = sp[0]
		isPercent = true
	}

	v = v + "00"
	d, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		return 0, isPercent, err
	}
	if d < 0 {
		return 0, isPercent, errors.NewHTTPError(errors.ErrMin0)
	}
	return d, isPercent, err
}
