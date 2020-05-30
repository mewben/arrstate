package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Countries slice
var Countries = make([]string, 0)

func init() {
	raw, err := ioutil.ReadFile("./countries.json")
	if err != nil {
		log.Fatalln("Error in countries", err)
	}
	json.Unmarshal(raw, &Countries)
}
