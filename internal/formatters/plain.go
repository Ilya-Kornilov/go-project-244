package formatters

import (
	"code/internal/differ"
	"fmt"
	"strings"
)

func Plain(nodes []differ.Node) string {
	return formatNodes(nodes, "")
}

func formatNodes(nodes []differ.Node, path string) string {
	lines := []string{}

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]

		currentPath := node.Key
		if path != "" {
			currentPath = path + "." + node.Key
		}

		if i+1 < len(nodes) &&
			nodes[i+1].Key == node.Key &&
			node.Status == "removed" &&
			nodes[i+1].Status == "added" {
			next := nodes[i+1]

			lines = append(lines, fmt.Sprintf(
				"Property '%s' was updated. From %s to %s",
				currentPath,
				valueFromNode(node),
				valueFromNode(next),
			))

			i++
			continue
		}

		switch node.Status {
		case "added":
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was added with value: %s",
				currentPath,
				valueFromNode(node),
			))

		case "removed":
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was removed",
				currentPath,
			))

		case "updated":
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was updated. From %s to %s",
				currentPath,
				formatValue(node.OldValue),
				formatValue(node.NewValue),
			))

		default:
			if len(node.Children) > 0 {
				childOutput := formatNodes(node.Children, currentPath)
				if childOutput != "" {
					lines = append(lines, childOutput)
				}
			}
		}
	}

	return strings.Join(lines, "\n")
}

func valueFromNode(node differ.Node) string {
	if len(node.Children) > 0 {
		return "[complex value]"
	}

	if node.Status == "added" {
		return formatValue(node.NewValue)
	}

	return formatValue(node.OldValue)
}

func formatValue(value interface{}) string {
	if value == nil {
		return "null"
	}

	if _, ok := value.(map[string]interface{}); ok {
		return "[complex value]"
	}

	switch value := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}
