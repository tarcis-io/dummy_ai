package main

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/component"
	"dummy_ai/internal/wasm/util"
)

func main() {

	cardTitleText := js.Global().Get("document").Call("createElement", "h2")
	cardTitleText.Set("className", "pf-v6-c-card__title-text")
	cardTitleText.Set("innerText", util.App())

	cardTitle := js.Global().Get("document").Call("createElement", "div")
	cardTitle.Set("className", "pf-v6-c-card__title")
	cardTitle.Call("appendChild", cardTitleText)

	cardBody := js.Global().Get("document").Call("createElement", "div")
	cardBody.Set("className", "pf-v6-c-card__body")
	cardBody.Set("innerText", util.AppDescription())

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")
	card.Call("appendChild", cardTitle)
	card.Call("appendChild", cardBody)

	component.CreateApp(card)
	select {}
}
