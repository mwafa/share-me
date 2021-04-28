package view

import (
	"text/template"

	"github.com/mwafa/share-me/utils"
	"github.com/valyala/fasthttp"
)

func LoginPage(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	loginPage, _ := template.ParseFS(FS, "template/login.html")
	err := loginPage.Execute(ctx, map[string]interface{}{
		"Key": utils.CookieKey(),
	})
	if err != nil {
		panic(err)
	}
}
