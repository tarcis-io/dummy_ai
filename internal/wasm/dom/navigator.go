package dom

type (
	Navigator struct {
		DOM
	}
)

func (navigator Navigator) MediaDevices() MediaDevices {

	return MediaDevices{
		DOM: navigator.Get("mediaDevices"),
	}
}

func GetNavigator() Navigator {

	return GetWindow().Navigator()
}
