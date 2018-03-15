package filter

import (
	"github.com/astaxie/beego/context"
	//"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego"
)

var FilterAll = func(ctx *context.Context) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", beego.AppConfig.String("header_origin"))       //允许访问源
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")    //允许post访问
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers:x-requested-with", "Content-Type,Authorization") // header的类型
	ctx.ResponseWriter.Header().Set("Access-Control-Request-Headers:x-requested-with", "Content-Type,Authorization")
	ctx.ResponseWriter.Header().Set("Access-Control-Expose-Headers:x-requested-with","content-type,Authorization") // 暴露的字段

	ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", beego.AppConfig.String("max_age"))
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
	ctx.ResponseWriter.Header().Set("Authorization", "") //返回数据格式是json
}
