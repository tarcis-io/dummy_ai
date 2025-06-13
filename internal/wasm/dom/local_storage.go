package dom

type (
	LocalStorage struct {
		DOM
	}
)

func GetLocalStorage() LocalStorage {

	return GetWindow().LocalStorage()
}
