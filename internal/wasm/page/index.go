package main

import (
	"syscall/js"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "index.wasm")

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", h2)

	select {}
}
