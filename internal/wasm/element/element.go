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

func (element Element) Id(id string) Element {

	element.domObject.Set("id", id)
	return element
}

func (element Element) Class(class string) Element {

	element.domObject.Set("className", class)
	return element
}
