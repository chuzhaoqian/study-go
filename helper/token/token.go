package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego/context"
	//"github.com/astaxie/beego"
)
// 建立 token
func New()string {
	claims := make(jwt.MapClaims)
	claims["sub"] = 0	// 面向的用户
	claims["exp"] = 0	// 过期时间
	claims["iat"] = 0	// 签发时间
	claims["jti"] = 0	// 唯一身份标识
	claims["aud"] = 0	// 接收方 UserAgent
	claims["ip"] = 0	// 接收方 ip
	return ""
}

func Set(ctx  *context.Context, jwt string){
	ctx.ResponseWriter.Header().Set("Authorization", jwt)
}

// 验证 token
func Valid() {

}

// 获取 Head 中 token
func Get(ctx  *context.Context) string {
	return ctx.Input.Header("Authorization")
}

