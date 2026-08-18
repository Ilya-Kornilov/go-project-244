package formatters

import (
	"encoding/json"

	"code/internal/differ"
)

func JSON(nodes []differ.Node) (string, error) {
	data, err := json.Marshal(nodes)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
