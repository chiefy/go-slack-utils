package blockui

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var basicSection = json.RawMessage(`{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "This is a mrkdwn section block :ghost: *this is bold*, and ~this is crossed out~"
			}
		}
	]
}`)

func TestSectionBlock(t *testing.T) {
	assert := assert.New(t)
	b := &SlackBlocks{}
	s := NewBlockSection()
	s.SetText(
		"mrkdwn",
		"This is a mrkdwn section block :ghost: *this is bold*, and ~this is crossed out~",
	)
	b.Push(s)

	j, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}
	assert.Equal(string(basicSection), string(j))
}
