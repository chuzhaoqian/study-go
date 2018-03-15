package initial

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func InitI18n() {
	lang := beego.AppConfig.String("lang")
	i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini")
}
