package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/internal/env"
)

var (
	staticFiles = map[string]string{
		"/manifest.json":                   "./static/config/manifest.json",
		"/robots.txt":                      "./static/config/robots.txt",
		"/sitemap.xml":                     "./static/config/sitemap.xml",
		"/apple_touch_icon.png":            "./static/img/favicon/apple_touch_icon.png",
		"/favicon.ico":                     "./static/img/favicon/favicon.ico",
		"/favicon.svg":                     "./static/img/favicon/favicon.svg",
		"/favicon_192.png":                 "./static/img/favicon/favicon_192.png",
		"/favicon_512.png":                 "./static/img/favicon/favicon_512.png",
		"/favicon_512_maskable.png":        "./static/img/favicon/favicon_512_maskable.png",
		"/logo.svg":                        "./static/img/logo/logo.svg",
		"/logo_white.svg":                  "./static/img/logo/logo_white.svg",
		"/fa_solid_900.woff2":              "./static/lib/patternfly/font/fa_solid_900.woff2",
		"/pf_v6_pficon.woff2":              "./static/lib/patternfly/font/pf_v6_pficon.woff2",
		"/red_hat_display_vf.woff2":        "./static/lib/patternfly/font/red_hat_display_vf.woff2",
		"/red_hat_display_vf_italic.woff2": "./static/lib/patternfly/font/red_hat_display_vf_italic.woff2",
		"/red_hat_mono_vf.woff2":           "./static/lib/patternfly/font/red_hat_mono_vf.woff2",
		"/red_hat_mono_vf_italic.woff2":    "./static/lib/patternfly/font/red_hat_mono_vf_italic.woff2",
		"/red_hat_text_vf.woff2":           "./static/lib/patternfly/font/red_hat_text_vf.woff2",
		"/red_hat_text_vf_italic.woff2":    "./static/lib/patternfly/font/red_hat_text_vf_italic.woff2",
		"/patternfly.css":                  "./static/lib/patternfly/patternfly.css",
		"/patternfly_addons.css":           "./static/lib/patternfly/patternfly_addons.css",
		"/wasm_exec.js":                    "./static/lib/wasm/wasm_exec.js",
		"/wasm_start.js":                   "./static/lib/wasm/wasm_start.js",
		"/about.wasm":                      "./static/wasm/about.wasm",
		"/error_404.wasm":                  "./static/wasm/error_404.wasm",
		"/home.wasm":                       "./static/wasm/home.wasm",
	}
)

var (
	pages = map[string]string{}
)

var (
	serverTemplate = template.Must(template.ParseFiles("./static/server/server.html"))
)

func Start() {

	for route, staticFile := range staticFiles {

		handleFile(route, staticFile)
	}

	for route, page := range pages {

		handlePage(route, page)
	}

	listenAndServe()
}

func handleFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func handlePage(route string, wasmRoute string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.URL.Path != route {

			handlePageError404(responseWriter)
			return
		}

		executeServerTemplate(responseWriter, wasmRoute)
	})
}

func handlePageError404(responseWriter http.ResponseWriter) {

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
