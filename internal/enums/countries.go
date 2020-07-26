package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/markbates/pkger"
)

// Countries slice
var Countries = make([]string, 0)

func init() {
	f, err := pkger.Open("/assets/countries.json")
	if err != nil {
		log.Println("Error opening countries.json")
		log.Fatalln(err)
	}
	defer f.Close()

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Error in countries", err)
	}
	json.Unmarshal(raw, &Countries)
}
