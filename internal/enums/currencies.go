package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Currency -
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Currencies slice
var Currencies = make([]Currency, 0)

func init() {
	raw, err := ioutil.ReadFile("./currencies.json")
	if err != nil {
		log.Fatalln("Error in currencies", err)
	}
	json.Unmarshal(raw, &Currencies)
}
