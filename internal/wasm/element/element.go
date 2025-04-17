package element

import (
	"dummy_ai/internal/wasm/dom"
)

type (
	Element struct {
		domObject dom.DOM
	}
)

func Create(element string) Element {

	return Element{
		domObject: dom.GetWindow().Document().CreateElement(element),
	}
}
