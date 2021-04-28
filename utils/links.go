package utils

import "fmt"

func GenLinks(pw string, port int) []string {
	ips := GetIP()

	var links []string
	for _, ip := range ips {
		l := fmt.Sprintf("http://%s:%s?pw=%s", ip, fmt.Sprint(port), pw)
		links = append(links, l)
	}
	return links
}
