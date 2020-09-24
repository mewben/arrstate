package utils

import (
	"math"
	"strconv"
	"strings"
)

// CalculateItem invoice
func CalculateItem(amount, tax int64, qty float64, discountStr string) (taxAmount, discountAmount, totalAmount int64, err error) {
	withQuantity := int64(math.Round(float64(amount) * qty))
	taxAmount = withQuantity * tax / 10000
	withTax := withQuantity + taxAmount
	// parse discount
	// example 5%
	sp := strings.Split(discountStr, "%")
	d := discountStr
	if len(sp) == 2 {
		d = sp[0] // 5
	}
	var discount int64
	if d != "" {
		discount, err = strconv.ParseInt(d, 0, 64) // 5
		if err != nil {
		}
	}

	if len(sp) == 2 {
		discountAmount = withTax * discount / 100
	} else {
		discountAmount = discount
	}

	totalAmount = withTax - discountAmount
	return
}
