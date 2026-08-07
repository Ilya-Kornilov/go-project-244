package differ

import (
	"code/internal/parser"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"
)

func GenDiff(path1, path2 string) (string, error) {
	left, err := parser.ParseFile(path1)
	if err != nil {
		return "", err
	}
	leftKeys := slices.Collect(maps.Keys(left))
	slices.Sort(leftKeys)
	
	right, err := parser.ParseFile(path2)
	if err != nil {
		return "", err 
	}
	rightKeys := slices.Collect(maps.Keys(right))

	allKeys := map[string]struct{}{} 
	for _, v := range rightKeys {
		allKeys[v] = struct{}{}
	}

	for _, v := range leftKeys {
		allKeys[v] = struct{}{}
	}

	keys := slices.Collect(maps.Keys(allKeys))
	slices.Sort(keys)	

	result := []string{}
	for _, k := range keys {
		l, okLeft := left[k]
		r, okRight := right[k]
		
		if okLeft && !okRight {
			result = append(result, fmt.Sprintf("  - %v: %v", k, l))
		} else if !okLeft && okRight {
			result = append(result, fmt.Sprintf("  + %v: %v", k, r))
		} else if okLeft && okRight && !reflect.DeepEqual(l, r) {
			result = append(result, fmt.Sprintf("  - %v: %v", k, l))
			result = append(result, fmt.Sprintf("  + %v: %v", k, r))
		} else {
			result = append(result, fmt.Sprintf("    %v: %v", k, l))
		}
	} 

	output := "{\n" + strings.Join(result, "\n") + "\n}"
	return output, nil 
}