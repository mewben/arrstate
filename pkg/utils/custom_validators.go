package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

// ValidateNumberOrPercentage -
func ValidateNumberOrPercentage(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	if v == "" {
		return true
	}
	return checkNumberOrPercentage(fl.Field().String())
}

// we separate this function to be able to test
func checkNumberOrPercentage(input string) bool {
	r, err := regexp.Compile(`^\d?(\.?\d)+%?$`)
	if err != nil {
		log.Println("err checknumberorpercentage", err)
		return false
	}

	if r.MatchString(input) {
		// check further for the percentage
		sp := strings.Split(input, "%")

		if len(sp) == 2 {
			// percentage must be <= 100
			v, err := strconv.ParseFloat(sp[0], 64)
			if err != nil {
				return false
			}
			if v > 100 {
				return false
			}
		}
		return true
	}

	return false
}
