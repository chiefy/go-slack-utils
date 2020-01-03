package blockui

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var basicDivider = json.RawMessage(`{
	"blocks": [
		{
			"type": "divider"
		}
	]
}`)

func TestDividerBlock(t *testing.T) {
	assert := assert.New(t)
	b := NewSlackBlocks()
	b.Push(NewSlackBlockDivider())
	j, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}
	assert.Equal(string(basicDivider), string(j))
}
