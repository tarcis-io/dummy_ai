package component

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/util"
)

func CreateCamera() js.Value {

	return CreateCameraLoading()
}

func CreateCameraLoading() js.Value {

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")

	return card
}

func CreateCameraError(title string, text string) js.Value {

	cardTitleText := js.Global().Get("document").Call("createElement", "h2")
	cardTitleText.Set("className", "pf-v6-c-card__title-text")
	cardTitleText.Set("innerText", title)

	cardTitle := js.Global().Get("document").Call("createElement", "div")
	cardTitle.Set("className", "pf-v6-c-card__title")
	cardTitle.Call("appendChild", cardTitleText)

	cardBody := js.Global().Get("document").Call("createElement", "div")
	cardBody.Set("className", "pf-v6-c-card__body")
	cardBody.Set("innerText", text)

	buttonText := js.Global().Get("document").Call("createElement", "span")
	buttonText.Set("className", "pf-v6-c-button__text")
	buttonText.Set("innerText", util.ButtonReloadPage())

	button := js.Global().Get("document").Call("createElement", "button")
	button.Set("className", "pf-v6-c-button pf-m-primary")
	button.Set("type", "button")
	button.Call("appendChild", buttonText)
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {

		js.Global().Get("location").Call("reload")
		return nil
	}))

	cardFooter := js.Global().Get("document").Call("createElement", "div")
	cardFooter.Set("className", "pf-v6-c-card__footer")
	cardFooter.Call("appendChild", button)

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")
	card.Call("appendChild", cardTitle)
	card.Call("appendChild", cardBody)
	card.Call("appendChild", cardFooter)

	return card
}
