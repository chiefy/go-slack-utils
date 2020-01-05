package blockui

type SectionAccessory interface {
	HasInteraction() bool
	IsImage() bool
	SetText(*BlockTitleText)
	SetValue(string)
}

// BlockSection represents a section Block Kit UI element
type BlockSection struct {
	Type      string                   `json:"type"`
	Text      *BlockTitleTextEmojiless `json:"text,omitempty"`
	Accessory SectionAccessory         `json:"accessory,omitempty"`
	Fields    []*BlockTitleText        `json:"fields,omitempty"`
}

// GetType implements SlackBlock interface
func (s BlockSection) GetType() string {
	return s.Type
}

// SetText sets the Text field as appropriate depending on if the section has an accessory or not
func (s *BlockSection) SetText(textType string, textVal string) {
	s.Text = NewBlockTitleTextEmojiless(textType)
	s.Text.SetText(textVal)
}

func (s BlockSection) HasAccessory() bool {
	return s.Fields != nil || s.Accessory != nil
}

// SetAccessory sets the accessory
func (s *BlockSection) SetAccessory(a SectionAccessory) {
	s.Accessory = a
}

// NewBlockSection creates a new empty section UI element
func NewBlockSection() *BlockSection {
	return &BlockSection{
		Type: slackBlockSectionType,
	}
}

// NewBlockSectionWithAccessory creates a new Block Kit section UI element with provided accesssory type
func NewBlockSectionWithAccessory(a SectionAccessory) *BlockSection {
	s := NewBlockSection()
	s.SetAccessory(a)
	return s
}

// NewBlockSectionWithImage creates a new Block Kit section UI element with image accessory
func NewBlockSectionWithImage(i *BlockAccessoryImage) *BlockSection {
	return NewBlockSectionWithAccessory(i)
}

// NewBlockSectionWithButton creates a new Block Kit section UI element with button accessory
func NewBlockSectionWithButton(b *BlockButton) *BlockSection {
	return NewBlockSectionWithAccessory(b)
}

// NewBlockSectionWithSelect creates a new Block Kit section UI element with select accessory
func NewBlockSectionWithSelect(s *BlockSelect) *BlockSection {
	return NewBlockSectionWithAccessory(s)
}
