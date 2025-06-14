package dom

type (
	Navigator struct {
		DOM
	}
)

func GetNavigator() Navigator {

	return GetWindow().Navigator()
}
