package services

import (
	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/mewben/arrstate/internal/enums"
)

// T - translates input to locale
// @input string
// @locale string
func T(input string) string {
	localizer := i18n.NewLocalizer(enums.I18nBundle, language.English.String())
	tn := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: input,
		// DefaultMessage: &i18n.Message{
		// 	ID:    input,
		// 	One:   "{{.Name}} has {{.Count}} cat.",
		// 	Other: "{{.Name}} has {{.Count}} cats.",
		// },
	})

	// TODO
	return tn
}
