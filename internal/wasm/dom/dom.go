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
	dom.jsObject.Set(property, unwrapValue(value))
}

func (dom DOM) Call(method string, arguments ...any) DOM {
	return DOM{
		jsObject: dom.jsObject.Call(method, unwrapValues(arguments)...),
	}
}

func GetGlobal() DOM {
	return DOM{
		jsObject: js.Global(),
	}
}

func unwrapValue(value any) any {
	if domObject, ok := value.(DOM); ok {
		return domObject.jsObject
	}
	return value
}

func unwrapValues(values []any) []any {
	unwrappedValues := make([]any, len(values))
	for index, value := range values {
		unwrappedValues[index] = unwrapValue(value)
	}
	return unwrappedValues
}
