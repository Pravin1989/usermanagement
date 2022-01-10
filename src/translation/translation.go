package translation

import (
	"strconv"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	defaultLocalizer *i18n.Localizer
)

func TraanslateToDefaultlanguage(messageCode string, params ...interface{}) (string, error) {
	return translate(defaultLocalizer, messageCode, params...)
}
func translate(localizer *i18n.Localizer, messageCode string, params ...interface{}) (string, error) {
	config := i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: messageCode,
		},
	}
	if len(params) != 0 {
		templateData := make(map[string]interface{}, len(params))
		for index, param := range params {
			key := strings.Join([]string{"arg", strconv.Itoa(index)}, "")
			templateData[key] = param
		}
		config.TemplateData = templateData
	}
	return localizer.Localize(&config)
}
