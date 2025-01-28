package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
	"dummy_ai/internal/wasm/util"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "index.wasm")

	app := component.CreateApp(h2)

	html := js.Global().Get("document").Get("documentElement")
	html.Set("lang", util.Language())
	html.Set("className", "pf-v6-theme-dark")

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)

	select {}
}
