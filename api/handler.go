package api

import (
	"github.com/mwafa/share-me/view"
	"github.com/valyala/fasthttp"
)

func MainHandler(ctx *fasthttp.RequestCtx, pw string) {
	if IsLogin(ctx, pw) {
		if ctx.IsPost() {
			Upload(ctx)
		}
		ListHandler(ctx)
		return
	}
	Login(ctx, pw)
	view.LoginPage(ctx)
}
