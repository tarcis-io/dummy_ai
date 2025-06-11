package dom

type (
	Window struct {
		dom DOM
	}
)

func GetWindow() Window {

	return Window{
		dom: GetGlobal(),
	}
}
