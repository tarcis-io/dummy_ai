package dom

type (
	LocalStorage struct {
		DOM
	}
)

func (localStorage LocalStorage) SetItem(key string, value string) {

	localStorage.Call("setItem", key, value)
}

func (localStorage LocalStorage) GetItem(key string) string {

	return localStorage.Call("getItem", key).String()
}

func GetLocalStorage() LocalStorage {

	return GetWindow().LocalStorage()
}
