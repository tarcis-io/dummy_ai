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

func GetGlobal() DOM {
	return DOM{
		jsObject: js.Global(),
	}
}
