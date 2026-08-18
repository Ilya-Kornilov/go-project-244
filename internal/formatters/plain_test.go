package formatters

import (
	"testing"

	"code/internal/differ"
	"github.com/stretchr/testify/assert"
)

func TestPlain(t *testing.T) {
	nodes := []differ.Node{
		{
			Key:    "common",
			Status: "",
			Children: []differ.Node{
				{
					Key:      "follow",
					Status:   "added",
					NewValue: false,
				},
				{
					Key:      "setting2",
					Status:   "removed",
					OldValue: 200,
				},
				{
					Key:      "setting3",
					Status:   "updated",
					OldValue: true,
					NewValue: nil,
				},
			},
		},
	}

	want := "Property 'common.follow' was added with value: false\n" +
		"Property 'common.setting2' was removed\n" +
		"Property 'common.setting3' was updated. From true to null"

	assert.Equal(t, want, Plain(nodes))
}
