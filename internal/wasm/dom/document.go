package dom

type (
	Document struct {
		DOM
	}
)

func GetDocument() Document {
	return GetWindow().Document()
}
