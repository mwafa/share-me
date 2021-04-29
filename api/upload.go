package api

import (
	"fmt"
	"os"
	"path"

	"github.com/valyala/fasthttp"
)

func Upload(ctx *fasthttp.RequestCtx) bool {
	p := string(ctx.Path())[1:]
	loc := path.Join(".", p)
	stat, err := os.Stat(loc)

	if err == nil || os.IsExist(err) {
		if stat.IsDir() {
			fh, err := ctx.FormFile("file")
			fmt.Printf("Upload to %s\n", loc)
			if err != nil {
				panic(err)
			}
			if err := fasthttp.SaveMultipartFile(fh, path.Join(loc, fh.Filename)); err != nil {
				panic(err)
			}
			return true
		}
	}
	return false

}
