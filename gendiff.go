package code

import (
	"code/internal/differ"
)

func GenDiff(path1, path2 string, format string) (string, error) {
	_ = format 
	return differ.GenDiff(path1, path2)
}
