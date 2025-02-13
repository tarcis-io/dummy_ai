package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
)

func main() {

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")

	component.CreateApp(card)
	select {}
}
