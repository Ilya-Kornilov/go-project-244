package formatters

import (
	"code/internal/differ"
	"fmt"
	"strings"
)

// transforms a `[]differ.Node` into a formatted string
// showing formatted differences, if any, between two files
func Stylish(nodes []differ.Node) string {
	return "{\n" + formatStylishNodes(nodes, 0) + "\n}"
}

func formatStylishNodes(nodes []differ.Node, depth int) string {
	lines := []string{}

	indent := strings.Repeat("    ", depth)
	propertyIndent := indent + "    "
	signIndent := indent + "  "

	for _, node := range nodes {
		switch node.Status {
		case "removed":
			lines = append(lines, formatStylishChangedNode("-", node, depth))

		case "added":
			lines = append(lines, formatStylishChangedNode("+", node, depth))

		case "updated":
			lines = append(lines, fmt.Sprintf(
				"%s- %s: %s",
				signIndent,
				node.Key,
				formatStylishValue(node.OldValue),
			))
			lines = append(lines, fmt.Sprintf(
				"%s+ %s: %s",
				signIndent,
				node.Key,
				formatStylishValue(node.NewValue),
			))

		case "unchanged":
			lines = append(lines, fmt.Sprintf(
				"%s%s: %s",
				propertyIndent,
				node.Key,
				formatStylishValue(node.OldValue),
			))

		default:
			lines = append(lines, fmt.Sprintf(
				"%s%s: {",
				propertyIndent,
				node.Key,
			))

			lines = append(lines, formatStylishNodes(node.Children, depth+1))

			lines = append(lines, propertyIndent+"}")
		}
	}

	return strings.Join(lines, "\n")
}

func formatStylishChangedNode(sign string, node differ.Node, depth int) string {
	indent := strings.Repeat("    ", depth)
	signIndent := indent + "  "
	childIndent := signIndent + "    "

	if len(node.Children) == 0 {
		value := node.OldValue
		if sign == "+" {
			value = node.NewValue
		}

		return fmt.Sprintf(
			"%s%s %s: %s",
			signIndent,
			sign,
			node.Key,
			formatStylishValue(value),
		)
	}

	lines := []string{
		fmt.Sprintf("%s%s %s: {", signIndent, sign, node.Key),
		formatStylishPlainNodes(node.Children, childIndent),
		signIndent + "}",
	}

	return strings.Join(lines, "\n")
}

func formatStylishPlainNodes(nodes []differ.Node, indent string) string {
	lines := []string{}

	for _, node := range nodes {
		if len(node.Children) > 0 {
			lines = append(lines, fmt.Sprintf(
				"%s%s: {",
				indent,
				node.Key,
			))

			lines = append(lines, formatStylishPlainNodes(
				node.Children,
				indent+"    ",
			))

			lines = append(lines, indent+"}")
			continue
		}

		value := node.OldValue
		if value == nil {
			value = node.NewValue
		}

		lines = append(lines, fmt.Sprintf(
			"%s%s: %s",
			indent,
			node.Key,
			formatStylishValue(value),
		))
	}

	return strings.Join(lines, "\n")
}

func formatStylishValue(value interface{}) string {
	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}
