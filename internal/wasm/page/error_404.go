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

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", component.CreateApp(card))

	select {}
}
