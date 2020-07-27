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

	log.Println("currentFile:", currentFile)
	log.Println("path.dir:", path.Dir(currentFile))
	log.Println("join:", path.Join(path.Dir(currentFile), "../../assets/countries.json"))

	raw, err := ioutil.ReadFile(path.Join(path.Dir(currentFile), "../../assets/countries.json"))
	if err != nil {
		log.Fatalln("Error in countries", err)
	}
	json.Unmarshal(raw, &Countries)
}
