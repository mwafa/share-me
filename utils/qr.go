package utils

import (
	"encoding/base64"
	"strings"

	"github.com/skip2/go-qrcode"
)

func GenQR() string {
	pw := Password
	port := GetPort()
	links := GenLinks(pw, port)
	payload := strings.Join(links, "#\n")
	png, _ := qrcode.Encode(payload, qrcode.Medium, 256)

	imgUrl := base64.StdEncoding.EncodeToString(png)
	return imgUrl
}
