package formatters

import (
	"testing"

	"code/internal/differ"

	"github.com/stretchr/testify/assert"
)

// tests the case when a `[]differ.Node` transforms into a JSON string
func TestJSON(t *testing.T) {
	nodes := []differ.Node{
		{
			Key:    "name",
			Status: "added",
			NewValue: "John",
		},
	}

	got, err := JSON(nodes)

	assert.NoError(t, err)

	want := `[{"Key":"name","Status":"added","OldValue":null,"NewValue":"John","Children":null}]`

	assert.Equal(t, want, got)
}
