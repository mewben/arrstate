package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

// Countries slice
var Countries = make([]string, 0)

func init() {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No caller information")
	}

	raw, err := ioutil.ReadFile(path.Join(path.Dir(currentFile), "./countries.json"))
	if err != nil {
		log.Fatalln("Error in countries", err)
	}
	json.Unmarshal(raw, &Countries)
}
