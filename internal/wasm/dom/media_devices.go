package dom

type (
	MediaDevices struct {
		DOM
	}
)

func (mediaDevices MediaDevices) GetUserMedia(constraints map[string]any) (MediaStream, error) {

	value, err := mediaDevices.Call("getUserMedia", constraints).Await()

	if err != nil {

		return MediaStream{}, err
	}

	mediaStream := MediaStream{
		DOM: value,
	}

	return mediaStream, nil
}

func GetMediaDevices() MediaDevices {

	return GetNavigator().MediaDevices()
}
