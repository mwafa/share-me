package view

import (
	"embed"
	"path"
	"text/template"

	"github.com/mwafa/share-me/utils"
	"github.com/valyala/fasthttp"
)

//go:embed template/*
var FS embed.FS

type link struct {
	Label string
	Path  string
}

func Index(ctx *fasthttp.RequestCtx, files []string, file string) {

	ctx.SetContentType("text/html")
	var ff []link
	for _, f := range files {
		ff = append(ff, link{
			Label: f,
			Path:  path.Join("/", file, f),
		})
	}
	var data = map[string]interface{}{
		"qr":    utils.GenQR(),
		"files": ff,
	}
	loginPage, _ := template.ParseFS(FS, "template/index.html")
	err := loginPage.Execute(ctx, data)
	if err != nil {
		panic(err)
	}
}
