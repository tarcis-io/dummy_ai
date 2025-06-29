package main

import (
	"dummy_ai/internal/wasm/dom"
)

func main() {

	h2 := dom.GetDocument().CreateElement("h2")
	h2.Set("innerText", "home.go")

	body := dom.GetDocument().Get("body")
	body.Call("appendChild", h2)
}
