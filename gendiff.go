package code

import (
	"fmt"

	"code/internal/differ"
	"code/internal/formatters"
)

func GenDiff(path1, path2, format string) (string, error) {
	nodes, err := differ.GenDiff(path1, path2)
	if err != nil {
		return "", err
	}

	result, err := formatters.Format(nodes, format)
	if err != nil {
		return "", fmt.Errorf("format diff: %w", err)
	}

	return result, nil
}
