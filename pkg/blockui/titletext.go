package blockui

type TitleText interface {
	HasEmoji() bool
	GetText() string
	SetText(string)
}

// BlockTitleText represents a title or text UI element
type BlockTitleText struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

// HasEmoji specifies if block title element uses emoji or not
func (tt BlockTitleText) HasEmoji() bool {
	return true
}

// GetText returns the string of title text
func (tt BlockTitleText) GetText() string {
	return tt.Text
}

// SetText sets the string value of the title
func (tt *BlockTitleText) SetText(text string) {
	tt.Text = text
}

// NewBlockTitleText creates a new BlockTitleText with provided type
func NewBlockTitleText(textType string) *BlockTitleText {
	return &BlockTitleText{
		Type:  textType,
		Emoji: false,
	}
}

// BlockTitleTextEmojiless represents a title or text UI element with NO EMOJI field
type BlockTitleTextEmojiless struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// HasEmoji specifies if block title element uses emoji or not
func (tt BlockTitleTextEmojiless) HasEmoji() bool {
	return false
}

// GetText returns the string of title text
func (tt BlockTitleTextEmojiless) GetText() string {
	return tt.Text
}

// SetText sets the string value of the title
func (tt *BlockTitleTextEmojiless) SetText(text string) {
	tt.Text = text
}

// NewBlockTitleTextEmojiless creates a new BloCkTitleText without emoji field with provided type
func NewBlockTitleTextEmojiless(textType string) *BlockTitleTextEmojiless {
	return &BlockTitleTextEmojiless{
		Type: textType,
	}
}
