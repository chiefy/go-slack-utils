package blockui

// NewSlackBlockDivider creates a new BlockUI divider element
func NewSlackBlockDivider() *SlackBlockDivider {
	return &SlackBlockDivider{
		Type: slackBlockDividerType,
	}
}

// SlackBlockDivider represents a BlockUI divider element
type SlackBlockDivider struct {
	Type string `json:"type"`
}

// GetType implements the SlackBlock interface and returns the type as string
func (d SlackBlockDivider) GetType() string {
	return d.Type
}
