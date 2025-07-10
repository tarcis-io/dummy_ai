package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func (dom DOM) Bool() bool {
	return dom.jsObject.Bool()
}

func (dom DOM) String() string {
	return dom.jsObject.String()
}

func (dom DOM) Int() int {
	return dom.jsObject.Int()
}

func (dom DOM) Float() float64 {
	return dom.jsObject.Float()
}

func (dom DOM) Truthy() bool {
	return dom.jsObject.Truthy()
}

func (dom DOM) Get(property string) DOM {
	return DOM{
		jsObject: dom.jsObject.Get(property),
	}
}

func (dom DOM) Set(property string, value any) {
	if domObject, ok := value.(DOM); ok {
		value = domObject.jsObject
	}
	dom.jsObject.Set(property, value)
}

func (dom DOM) Call(method string, arguments ...any) DOM {
	for index, argument := range arguments {
		if domObject, ok := argument.(DOM); ok {
			arguments[index] = domObject.jsObject
		}
	}
	return DOM{
		jsObject: dom.jsObject.Call(method, arguments...),
	}
}

func GetGlobal() DOM {
	return DOM{
		jsObject: js.Global(),
	}
}
