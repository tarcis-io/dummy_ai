package dom

type (
	Window struct {
		DOM
	}
)

func GetWindow() Window {

	return Window{
		DOM: GetGlobal(),
	}
}
