package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/markbates/pkger"
)

// Init some enums
func Init() {
	// Countries
	box, err := pkger.Open("/assets/countries.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer box.Close()

	raw, err := ioutil.ReadAll(box)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(raw, &Countries)

	// Currencies
	box2, err := pkger.Open("/assets/currencies.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer box2.Close()

	raw2, err := ioutil.ReadAll(box2)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(raw2, &Currencies)

	// init i18n
	InitI18n()
}
