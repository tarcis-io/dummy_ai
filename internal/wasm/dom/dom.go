package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func (dom DOM) Get(property string) DOM {

	return DOM{
		jsObject: dom.jsObject.Get(property),
	}
}

func (dom DOM) GetString(property string) string {

	return dom.Get(property).jsObject.String()
}

func (dom DOM) Set(property string, value any) {

	dom.jsObject.Set(property, value)
}

func (dom DOM) Call(method string, arguments ...any) DOM {

	return DOM{
		jsObject: dom.jsObject.Call(method, arguments...),
	}
}

func GetGlobal() DOM {

	return DOM{
		jsObject: js.Global(),
	}
}
