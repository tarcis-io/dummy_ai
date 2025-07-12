package dom

type (
	LocalStorage struct {
		DOM
	}
)

func (localStorage LocalStorage) SetItem(key string, value string) {
	localStorage.Call("setItem", key, value)
}

func (localStorage LocalStorage) GetItem(key string) (string, bool) {
	if value := localStorage.Call("getItem", key); value.Truthy() {
		return value.String(), true
	}
	return "", false
}

func GetLocalStorage() LocalStorage {
	return GetWindow().LocalStorage()
}
