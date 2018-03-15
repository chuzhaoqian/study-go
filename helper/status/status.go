package status

import (
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/config"
	"github.com/beego/i18n"
	"strconv"
	"study-go/helper/lang"
)

//var conf config.Configer
var locale *i18n.Locale
var err error

func init() {
	//conf, err = config.NewConfig("ini", beego.AppPath+"/conf/status.conf")
	locale = lang.Init()
}

func Status(code int) (int, string, error) {
	if nil != err {
		return 1001, "", err
	}
	//str := conf.String(strconv.Itoa(code))
	str := locale.Tr(strconv.Itoa(code))
	return code, str, nil
}

func InValid() int {
	return 1000
}

func Error() int {
	return 1002
}

func Ok() (int, string, error) {
	var code int = 200
	if nil != err {
		return 1001, "", err
	}
	//str := conf.String(strconv.Itoa(code))
	str := locale.Tr(strconv.Itoa(code))
	return code, str, nil
}
