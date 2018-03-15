package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"study-go/helper/filter"
	"study-go/helper/response"
	_ "study-go/initial"
	_ "study-go/routers"
	//"study-go/service/user"
)

func main() {
	beego.BConfig.ServerName = beego.AppConfig.String("server_name")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterAll)
	beego.ErrorHandler("404", pageNotFound)
	beego.ErrorHandler("401", pageNotPermission)
	beego.Run()
}

func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	//userInfo, err := user.Login("18610373739", "123456")
	//fmt.Fprint(rw, userInfo, err)

	resp, _ := response.NewErrorRespJsonStr(404)
	fmt.Fprint(rw, resp)
}

func pageNotPermission(rw http.ResponseWriter, r *http.Request) {
	resp, _ := response.NewErrorRespJsonStr(401)
	fmt.Fprint(rw, resp)
}
