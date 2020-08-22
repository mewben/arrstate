package enums

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/markbates/pkger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// I18nBundle -
var I18nBundle *i18n.Bundle

// InitI18n -
func InitI18n() {
	I18nBundle = i18n.NewBundle(language.English)
	I18nBundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	loadTranslations("en", "/web/static/locales/en/global.json")
}

// loadTranslations
func loadTranslations(lang, path string) {
	box, err := pkger.Open(path)
	if err != nil {
		log.Fatalln("error open", path, err)
	}
	defer box.Close()

	raw, err := ioutil.ReadAll(box)
	if err != nil {
		log.Fatalln("error read", path, err)
	}

	I18nBundle.MustParseMessageFileBytes(raw, lang+".json")
}
