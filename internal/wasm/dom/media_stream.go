package dom

type (
	MediaStream struct {
		DOM
	}
)

func GetUserMedia(constraints map[string]any) (MediaStream, error) {

	value, err := GetMediaDevices().Call("getUserMedia").Await()

	if err != nil {

		return MediaStream{}, err
	}

	mediaStream := MediaStream{
		DOM: value,
	}

	return mediaStream, nil
}
