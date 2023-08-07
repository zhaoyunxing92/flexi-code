package translator

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	LanguageChinese Language = "zh_CN"
	LanguageEnglish Language = "en_US"

	DefaultLanguage = LanguageChinese
)

type Language string

func Tr(lang Language, reason string) string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.LoadMessageFileFS(I18n, "locale.es.yaml")
	return ""
}
