package differ

import (
	"code/internal/parser"
	"maps"
	"reflect"
	"slices"
)

type Node struct {
	Key      string
	Status   string
	OldValue interface{}
	NewValue interface{}
	Children []Node
}

func toData(value interface{}) (parser.Data, bool) {
	switch data := value.(type) {
	case parser.Data:
		return data, true
	case map[string]interface{}:
		return parser.Data(data), true
	default:
		return nil, false
	}
}

func buildDiff(left, right parser.Data) []Node {
	keys := map[string]struct{}{}

	for key := range left {
		keys[key] = struct{}{}
	}

	for key := range right {
		keys[key] = struct{}{}
	}

	allKeys := slices.Collect(maps.Keys(keys))
	slices.Sort(allKeys)

	nodes := make([]Node, 0, len(allKeys))

	for _, key := range allKeys {
		l, okLeft := left[key]
		r, okRight := right[key]

		if okLeft && okRight {
			lData, lOk := toData(l)
			rData, rOk := toData(r)

			switch {
			case lOk && rOk:
				nodes = append(nodes, Node{
					Key:      key,
					Children: buildDiff(lData, rData),
				})

			case lOk:
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "removed",
					Children: buildDiff(lData, parser.Data{}),
				})
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "added",
					NewValue: r,
				})

			case rOk:
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "removed",
					OldValue: l,
				})
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "added",
					Children: buildDiff(parser.Data{}, rData),
				})

			case !reflect.DeepEqual(l, r):
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "updated",
					OldValue: l,
					NewValue: r,
				})

			default:
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "unchanged",
					OldValue: l,
				})
			}
		} else if okLeft {
			if lData, ok := toData(l); ok {
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "removed",
					Children: buildDiff(lData, parser.Data{}),
				})
			} else {
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "removed",
					OldValue: l,
				})
			}
		} else {
			if rData, ok := toData(r); ok {
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "added",
					Children: buildDiff(parser.Data{}, rData),
				})
			} else {
				nodes = append(nodes, Node{
					Key:      key,
					Status:   "added",
					NewValue: r,
				})
			}
		}
	}

	return nodes
}

// parces the two files into a `[]differ.Node`
func GenDiff(path1, path2 string) ([]Node, error) {
	left, err := parser.ParseFile(path1)
	if err != nil {
		return nil, err
	}

	right, err := parser.ParseFile(path2)
	if err != nil {
		return nil, err
	}

	return buildDiff(left, right), nil
}
