package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "index.wasm")

	app := component.CreateApp(h2)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)

	select {}
}
