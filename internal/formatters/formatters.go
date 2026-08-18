package formatters

import (
	"fmt"

	"code/internal/differ"
)

func Format(nodes []differ.Node, format string) (string, error) {
	switch format {
	case "stylish":
		return Stylish(nodes), nil
	case "plain":
		return Plain(nodes), nil
	case "json":
		return JSON(nodes)
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
