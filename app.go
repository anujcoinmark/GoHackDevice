package main

import (
	"fmt"
	"net/http"
	"time"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	deviceType := GetType(r)
	operSys := GOOSS()
	if operSys == "linux" {
		fmt.Fprintf(w, "<h1>Operating System: %s<h1>", operSys)
	}
	if operSys == "windows" {
		fmt.Fprintf(w, "<h1>Operating System: %s<h1>", operSys)
	}
	if operSys == "darwin" {
		fmt.Fprintf(w, "<h1>Operating System: %s<h1>", operSys)
	}

	if deviceType == "Mobile" {
		fmt.Fprintf(w, "<h1>Device : Mobile</h1>")
	} else if deviceType == "Web" {
		fmt.Fprintf(w, "<h1>Device : Web</h1>")
	} else if deviceType == "Tab" {
		fmt.Fprintf(w, "<h1>Device : Tablet</h1>")
	}

	elapsed := time.Since(start)
	fmt.Printf("Time take to serve static file %s\n\n", elapsed)
}

func main() {
	//ShellScript()
	http.HandleFunc("/", RouteHandler)
	error := http.ListenAndServe(":9999", nil)
	fmt.Println(error)
}
