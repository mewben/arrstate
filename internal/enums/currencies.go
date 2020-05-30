package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"runtime"
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
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No caller information")
	}

	raw, err := ioutil.ReadFile(path.Join(path.Dir(currentFile), "./currencies.json"))
	if err != nil {
		log.Fatalln("Error in currencies", err)
	}
	json.Unmarshal(raw, &Currencies)
}
