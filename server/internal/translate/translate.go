package translate

import (
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"

	"github.com/zhaoyunxing92/flexi-code/server/i18n"
)

var bundle *goi18n.Bundle

func init() {
	bundle = goi18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	fs, err := bundle.LoadMessageFileFS(i18n.I18n, "")
	if err != nil {
		return
	}
	err = bundle.AddMessages(language.Chinese, fs.Messages...)
	if err != nil {
		return
	}
}

func Tr(lang string, key string) string {
	local := goi18n.NewLocalizer(bundle, lang)
	return local.MustLocalize(&goi18n.LocalizeConfig{
		DefaultMessage: &goi18n.Message{
			ID: key,
		}})
}
