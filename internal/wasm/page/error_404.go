package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
	"dummy_ai/internal/wasm/util"
)

func init() {

	html := js.Global().Get("document").Get("documentElement")
	html.Set("className", "pf-v6-theme-dark")
	html.Set("lang", util.Language())
}

func main() {

	pageContent := js.Global().Get("document").Call("createElement", "div")

	app := component.CreateApp(pageContent)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)

	select {}
}
