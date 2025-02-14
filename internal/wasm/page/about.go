package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
)

func main() {

	cardTitle := js.Global().Get("document").Call("createElement", "title")
	cardTitle.Set("className", "pf-v6-c-card__title")

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")
	card.Call("appendChild", cardTitle)

	component.CreateApp(card)
	select {}
}
