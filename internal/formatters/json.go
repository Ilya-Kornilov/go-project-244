package formatters

import (
	"encoding/json"

	"code/internal/differ"
)

// transforms `[]differ.Node` into a JSON representation string
func JSON(nodes []differ.Node) (string, error) {
	data, err := json.Marshal(nodes)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
