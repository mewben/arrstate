package utils

import (
	"strconv"
	"strings"
)

// CalculateItem invoice
func CalculateItem(amount, qty, tax float64, discountStr string) (taxAmount, discountAmount, totalAmount float64, err error) {
	taxAmount = amount * qty * (tax / 100.0)
	withTax := amount*qty + taxAmount
	// parse discount
	sp := strings.Split(discountStr, "%")
	d := discountStr
	if len(sp) == 2 {
		d = sp[0]
	}
	var discount float64
	if d != "" {
		discount, err = strconv.ParseFloat(d, 64)
		if err != nil {
			return
		}
	}

	if len(sp) == 2 {
		discountAmount = withTax * discount / 100.0
	} else {
		discountAmount = discount
	}

	totalAmount = withTax - discountAmount
	return
}
