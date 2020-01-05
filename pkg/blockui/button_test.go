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

	button := NewBlockButton()
	button.SetValue("click_me_123")

	tt := NewBlockTitleText(slackTextPlainText)
	tt.Text = "Button"
	tt.Emoji = true
	button.SetText(tt)

	s := NewBlockSectionWithButton(button)
	s.SetText(
		slackTextMarkdown,
		"You can add a button alongside text in your message. ",
	)
	b.Push(s)

	j, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}
	assert.Equal(string(sectionWithButtonBlock), string(j))
}
