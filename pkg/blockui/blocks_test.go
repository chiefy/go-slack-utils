package blockui

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var basicBlocks = json.RawMessage(`{
	"blocks": []
}`)

func TestBlocks(t *testing.T) {
	assert := assert.New(t)
	b := NewSlackBlocks()
	j, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}
	assert.Equal(string(basicBlocks), string(j))
}
