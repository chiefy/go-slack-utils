package payload

import (
	"github.com/chiefy/go-slack-utils/pkg/blockui"
)

// NewMessagePayload creates a new MessagePayload with provided response type
func NewMessagePayload(responseType string) *MessagePayload {
	return &MessagePayload{
		ResponseType:    responseType,
		ReplaceOriginal: false,
		SlackBlocks:     blockui.NewSlackBlocks(),
	}
}

// MessagePayload is the payload structure to send to the Slack API
type MessagePayload struct {
	ResponseType    string `json:"response_type,omitempty"`
	ReplaceOriginal bool   `json:"replace_original"`
	*blockui.SlackBlocks
}

// AddBlock adds a new block to the payload
func (mp *MessagePayload) AddBlock(block blockui.SlackBlock) {
	mp.SlackBlocks.Push(block)
}

// AddDivider adds a new block to the payload
func (mp *MessagePayload) AddDivider() {
	mp.SlackBlocks.Push(blockui.NewSlackBlockDivider())
}
