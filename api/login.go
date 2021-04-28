package api

import (
	"fmt"

	"github.com/mwafa/share-me/utils"
	"github.com/valyala/fasthttp"
)

func Login(ctx *fasthttp.RequestCtx, pw string) {
	password := ctx.QueryArgs().Peek(utils.CookieKey())
	if pw == string(password) {
		fmt.Println("-------")
		fmt.Println("Login Success !")
		cook := fasthttp.Cookie{}
		cook.SetKey(utils.CookieKey())
		cook.SetValue(string(password))
		cook.HTTPOnly()
		cook.SetMaxAge(3600000)
		cook.SetPath(("/"))
		ctx.Response.Header.SetCookie(&cook)
		ctx.Redirect(string(ctx.Path()), fasthttp.StatusTemporaryRedirect)
	}
}

func IsLogin(ctx *fasthttp.RequestCtx, pw string) bool {
	password := ctx.Request.Header.Cookie(utils.CookieKey())
	if string(password) == pw {
		return true
	}
	return false
}
