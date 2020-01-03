package blockui

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var basicActionsBlock = json.RawMessage(`{
	"blocks": [
		{
			"type": "actions",
			"block_id": "actions1",
			"elements": []
		}
	]
}`)

func TestActionsBlock(t *testing.T) {
	assert := assert.New(t)
	b := NewSlackBlocks()
	a := NewActionsBlock("actions1")
	b.Push(a)
	j, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}
	assert.Equal(string(basicActionsBlock), string(j))
}
