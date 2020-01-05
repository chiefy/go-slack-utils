package blockui

// BlockButton represents a Block Kit button UI element
type BlockButton struct {
	Type     string `json:"type"`
	ActionID string `json:"action_id,omitempty"`
	ActionTS string `json:"action_ts,omitempty"`
	BlockID  string `json:"block_id,omitempty"`
	*BlockOption
}

// NewBlockButton creates a new Block Kit button UI element
func NewBlockButton() *BlockButton {
	return &BlockButton{
		Type: slackAccessoryButtonType,
		BlockOption: &BlockOption{
			Text:  nil,
			Value: "",
		},
	}
}

// SetText implements the SectionAccessory interface, sets the text block of element
func (b *BlockButton) SetText(text *BlockTitleText) {
	b.BlockOption.Text = text
}

// SetValue implements the SectionAccessory interface, sets the button's value
func (b *BlockButton) SetValue(val string) {
	b.BlockOption.Value = val
}

// GetValue implements the ActionElement interface, gets the button's value
func (b BlockButton) GetValue() string {
	return b.BlockOption.Value
}

// HasInteraction implements the SectionAccessory interface
func (b BlockButton) HasInteraction() bool {
	return true
}

// HasOptions implements the ActionElement interface
func (b BlockButton) HasOptions() bool {
	return true
}

// UsesPlaceholder implements the ActionElement interface
func (b BlockButton) UsesPlaceholder() bool {
	return false
}

// IsImage implements the SectionAccessory interface
func (b BlockButton) IsImage() bool {
	return false
}

// GetActionID implements the ActionElement interface
func (b BlockButton) GetActionID() string {
	return b.ActionID
}

// GetOptions implements the ActionElement interface
func (b BlockButton) GetOptions() []*BlockOption {
	return []*BlockOption{b.BlockOption}
}
