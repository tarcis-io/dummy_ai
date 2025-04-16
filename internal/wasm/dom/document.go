package dom

type (
	Document struct {
		DOM
	}
)

func (document Document) CreateElement(element string) DOM {

	return DOM{
		jsObject: document.jsObject.Call("createElement", element),
	}
}
