package formatters

import (
	"testing"

	"code/internal/differ"

	"github.com/stretchr/testify/assert"
)

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
