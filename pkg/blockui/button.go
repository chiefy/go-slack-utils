package blockui

// BlockButton represents a Block Kit button UI element
type BlockButton struct {
	Type string `json:"type"`
	*BlockOption
}

// NewBlockButton creates a new Block Kit button UI element
func NewBlockButton() *BlockButton {
	return &BlockButton{
		slackAccessoryButtonType,
		&BlockOption{
			Text:  nil,
			Value: "",
		},
	}
}

// SetText implements the SectionAccessory interface, sets the text block of element
func (b *BlockButton) SetText(text *BlockTitleText) {
	b.BlockOption = &BlockOption{
		Text:  text,
		Value: "",
	}
}

// SetValue implements the SectionAccessory interface, sets the button's value
func (b *BlockButton) SetValue(val string) {
	b.BlockOption.Value = val
}

// HasInteraction implements the SectionAccessory interface
func (b BlockButton) HasInteraction() bool {
	return true
}

// IsImage implements the SectionAccessory interface
func (b BlockButton) IsImage() bool {
	return false
}
