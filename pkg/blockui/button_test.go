package blockui

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sectionWithButtonBlock = json.RawMessage(`{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "You can add a button alongside text in your message. "
			},
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Button",
					"emoji": true
				},
				"value": "click_me_123"
			}
		}
	]
}`)

func TestButtonAccessory(t *testing.T) {
	assert := assert.New(t)
	b := NewSlackBlocks()

	s, _ := NewSlackBlockSectionWithAccessory(slackAccessoryButtonType)
	s.SetText(
		slackTextMarkdown,
		"You can add a button alongside text in your message. ",
	)

	tt := NewBlockTitleText(slackTextPlainText)
	tt.Text = "Button"
	tt.Emoji = true

	s.Accessory.SetText(tt)
	s.Accessory.SetValue("click_me_123")
	b.Push(s)

	j, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}
	assert.Equal(string(sectionWithButtonBlock), string(j))
}
