package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func (dom DOM) Call(method string, args ...any) DOM {

	return DOM{
		jsObject: dom.jsObject.Call(method, args...),
	}
}

func GetGlobal() js.Value {

	return js.Global()
}
