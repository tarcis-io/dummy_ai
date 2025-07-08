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

func GetGlobal() DOM {
	return DOM{
		jsObject: js.Global(),
	}
}
