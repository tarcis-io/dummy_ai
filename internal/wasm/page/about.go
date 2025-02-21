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

	cardBodyAppDescription := js.Global().Get("document").Call("createElement", "div")
	cardBodyAppDescription.Set("className", "pf-v6-c-card__body")
	cardBodyAppDescription.Set("innerText", util.AppDescription())

	cardBodyAppVersion := js.Global().Get("document").Call("createElement", "div")
	cardBodyAppVersion.Set("className", "pf-v6-c-card__body")
	cardBodyAppVersion.Set("innerText", util.AppVersion())

	cardFooter := js.Global().Get("document").Call("createElement", "div")
	cardFooter.Set("className", "pf-v6-c-card__footer")
	cardFooter.Set("innerText", util.AppDevelopedBy())

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")
	card.Call("appendChild", cardTitle)
	card.Call("appendChild", cardBodyAppDescription)
	card.Call("appendChild", cardBodyAppVersion)
	card.Call("appendChild", cardFooter)

	component.CreateApp(card)

	select {}
}
