package api

import (
	"fmt"
	"os"

	"github.com/mwafa/share-me/view"
	"github.com/valyala/fasthttp"
)

func ListHandler(ctx *fasthttp.RequestCtx) {
	fmt.Printf("%s - %s \n", ctx.Method(), ctx.Path())
	file := string(ctx.Path())[1:]
	info, err := os.Stat(file)

	if err == nil || os.IsExist(err) {
		if info.IsDir() {
			files, _ := os.ReadDir(file)
			var ff []string
			for _, f := range files {
				ff = append(ff, f.Name())
			}
			view.Index(ctx, ff, file)
		} else {
			fasthttp.ServeFileUncompressed(ctx, file)
		}
		return
	}
	files, _ := os.ReadDir(".")
	var ff []string
	for _, f := range files {
		ff = append(ff, f.Name())
	}
	view.Index(ctx, ff, "")
}
