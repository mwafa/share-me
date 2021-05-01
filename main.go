package main

import (
	"fmt"
	"log"

	"github.com/mwafa/share-me/api"
	"github.com/mwafa/share-me/utils"
	"github.com/valyala/fasthttp"
)

func main() {
	port := utils.GetPort()
	ips := utils.GetIP()
	pw := utils.Password

	fmt.Printf("Password: %s\n", pw)
	fmt.Println("Server Running di:")
	for _, ip := range ips {
		fmt.Printf("http://%s:%d\n", ip, port)
	}

	link := fmt.Sprintf(":%d", port)
	utils.OpenBrowser(fmt.Sprintf("http://localhost:%d?pw=%s", port, pw))

	s := &fasthttp.Server{
		Handler: func(c *fasthttp.RequestCtx) {
			api.MainHandler(c, pw)
		},
		Name:               "Share Me",
		MaxRequestBodySize: 4 * 1024 * 1024 * 1024,
	}
	log.Fatal(s.ListenAndServe(link))

}
