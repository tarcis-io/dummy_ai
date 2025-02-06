package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
)

func main() {

	pageContent := js.Global().Get("document").Call("createElement", "div")

	app := component.CreateApp(pageContent)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)

	select {}
}
