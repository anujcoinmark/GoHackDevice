package main

import (
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

type DeviceType string

const (
	MOBILE DeviceType = "Mobile"
	TABLET DeviceType = "Tab"
	WEB    DeviceType = "Web"
	TV     DeviceType = "TV"
)

func GetType(r *http.Request) DeviceType {

	if isUserAgent(r, "Android", "webOS", "iPhone", "BlackBerry", "Windows Phone") {
		return MOBILE
	}
	if isUserAgent(r, "iPad", "iPod", "tablet", "RX-34", "FOLIO") ||
		(isUserAgent(r, "Kindle", "Mac OS") && isUserAgent(r, "Silk")) ||
		(isUserAgent(r, "AppleWebKit") && isUserAgent(r, "Silk")) {
		return TABLET
	}
	if isUserAgent(r, "TV", "NetCast", "boxee", "Kylo", "Roku", "DLNADOC") {
		return TV
	}

	return WEB
}

func isUserAgent(r *http.Request, userAgents ...string) bool {
	userAgent := r.Header.Get("User-Agent")
	for _, v := range userAgents {
		if strings.Contains(userAgent, v) {
			return true
		}
	}
	return false
}

func GOOSS() string {
	os := runtime.GOOS
	return os
}

func ShellScript() {
	out, err := exec.Command("/bin/sh", "sample.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
