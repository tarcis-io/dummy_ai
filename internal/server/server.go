package server

import (
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

var (
	staticFiles = map[string]string{
		"/manifest.json":            "./static/configuration/manifest.json",
		"/robots.txt":               "./static/configuration/robots.txt",
		"/sitemap.xml":              "./static/configuration/sitemap.xml",
		"/apple_touch_icon.png":     "./static/image/favicon/apple_touch_icon.png",
		"/favicon_192.png":          "./static/image/favicon/favicon_192.png",
		"/favicon_512_maskable.png": "./static/image/favicon/favicon_512_maskable.png",
		"/favicon_512.png":          "./static/image/favicon/favicon_512.png",
		"/favicon.ico":              "./static/image/favicon/favicon.ico",
		"/favicon.svg":              "./static/image/favicon/favicon.svg",
		"/logo_white.svg":           "./static/image/logo/logo_white.svg",
		"/logo.svg":                 "./static/image/logo/logo.svg",
	}
)

var (
	pages = map[string]string{}
)

var (
	serverTemplate = template.Must(template.ParseFiles("./static/server/template.html"))
)

func ListenAndServe() {
	for route, file := range staticFiles {
		serveFile(route, file)
	}
	for route, wasmRoute := range pages {
		servePage(route, wasmRoute)
	}
	listenAndServe()
}

func serveFile(route string, file string) {
	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {
		http.ServeFile(responseWriter, request, file)
	})
}

func servePage(route string, wasmRoute string) {
	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.URL.Path != route {
			servePageError404(responseWriter)
			return
		}
		executeServerTemplate(responseWriter, wasmRoute)
	})
}

func servePageError404(responseWriter http.ResponseWriter) {
	responseWriter.WriteHeader(http.StatusNotFound)
	executeServerTemplate(responseWriter, "/error_404.wasm")
}

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {
	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {
		panic(err)
	}
}

func listenAndServe() {
	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {
		panic(err)
	}
}
