package main

import (
	"code/internal/differ"
	"testing"
)

func TestGenDiff(t *testing.T) {
	res, err := differ.GenDiff("../../testdata/file1.json", "../../testdata/file2.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	exp :=  "{\n" +
		"  - follow: false\n" +
		"    host: hexlet.io\n" +
		"  - proxy: 123.234.53.22\n" +
		"  - timeout: 50\n" +
		"  + timeout: 20\n" +
		"  + verbose: true\n" +
		"}"
	if res != exp {
		t.Errorf("expected: `%s`; got: `%s`", exp, res)
	}
}
