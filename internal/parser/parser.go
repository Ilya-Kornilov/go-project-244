package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Data map[string]interface{}

// reads the file into a byte slice
func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// transforms a JSON file into `Data` type
func ParseJSON(data []byte) (Data, error) {
	d := Data{}
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&d)

	return d, err
}

// transforms a YAML file into `Data` type
func ParseYAML(data []byte) (Data, error) {
	d := Data{}
	err := yaml.Unmarshal(data, &d)

	return d, err
}

// parses the file
func ParseFile(path string) (Data, error) {
	ext := filepath.Ext(path)

	f, err := ReadFile(path)
	if err != nil {
		return nil, err
	}

	d, err := Parse(f, ext)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// checks the file extension
func Parse(data []byte, ext string) (Data, error) {
	switch ext {
	case ".json":
		return ParseJSON(data)

	case ".yaml", ".yml":
		return ParseYAML(data)
	}

	return nil, errors.New("unsupported file format")
}
