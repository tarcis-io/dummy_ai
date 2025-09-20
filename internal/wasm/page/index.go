package main

import (
	"syscall/js"
)

func main() {
	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerText", "index.wasm")
	js.Global().Get("document").Get("body").Call("appendChild", h2)
}
