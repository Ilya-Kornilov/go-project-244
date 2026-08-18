package main

import (
	"code"
	"testing"
)

func TestGenDiffJSON(t *testing.T) {
	res, err := code.GenDiff(
		"../../testdata/file1.json", 
		"../../testdata/file2.json",
		"stylish",
	)
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

func TestGenDiffYAML(t *testing.T) {
	res, err := code.GenDiff(
		"../../testdata/file1.yaml", 
		"../../testdata/file2.yaml",
		"stylish",
	)
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
