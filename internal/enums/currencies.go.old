package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/markbates/pkger"
)

// Currency -
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Currencies slice
var Currencies = make([]Currency, 0)

func init3() {
	f, err := pkger.Open("/assets/currencies.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Error in currencies", err)
	}
	json.Unmarshal(raw, &Currencies)
}
