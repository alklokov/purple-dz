package request

import (
	"encoding/json"
	"io"
)

func decode[T any](body io.ReadCloser) (T, error) {
	var object T
	err := json.NewDecoder(body).Decode(&object)
	if err != nil {
		return object, err
	}
	return object, nil
}
