package lang

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func Init() *i18n.Locale {
	lang := beego.AppConfig.String("lang")
	locale := new(i18n.Locale)
	locale.Lang = lang
	return locale
}
