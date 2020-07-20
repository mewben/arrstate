package utils

import (
	"encoding/json"
	"log"
)

func PrettyJSON(input interface{}) {
	m, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(string(m))
}
