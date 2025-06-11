package dom

type (
	Document struct {
		DOM
	}
)

func (document Document) CreateElement(element string) DOM {

	return document.Call("createElement", element)
}

func GetDocument() Document {

	return GetWindow().Document()
}
