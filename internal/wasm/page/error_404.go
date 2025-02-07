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

	cardTitleText := js.Global().Get("document").Call("createElement", "h2")
	cardTitleText.Set("className", "pf-v6-c-card__title-text")
	cardTitleText.Set("innerText", util.Error404())

	cardTitle := js.Global().Get("document").Call("createElement", "div")
	cardTitle.Set("className", "pf-v6-c-card__title")
	cardTitle.Call("appendChild", cardTitleText)

	cardBody := js.Global().Get("document").Call("createElement", "div")
	cardBody.Set("className", "pf-v6-c-card__body")
	cardBody.Set("innerText", util.Error404Message())

	buttonText := js.Global().Get("document").Call("createElement", "span")
	buttonText.Set("className", "pf-v6-c-button__text")
	buttonText.Set("innerText", util.ButtonReturnToHome())

	button := js.Global().Get("document").Call("createElement", "a")
	button.Set("className", "pf-v6-c-button pf-m-primary")
	button.Set("href", "/")
	button.Call("appendChild", buttonText)

	cardFooter := js.Global().Get("document").Call("createElement", "div")
	cardFooter.Set("className", "pf-v6-c-card__footer")
	cardFooter.Call("appendChild", button)

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")
	card.Call("appendChild", cardTitle)
	card.Call("appendChild", cardBody)
	card.Call("appendChild", cardFooter)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", component.CreateApp(card))

	select {}
}
