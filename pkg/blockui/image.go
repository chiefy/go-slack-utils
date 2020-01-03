package blockui

func NewSlackBlockImage() *SlackBlockImage {
	return &SlackBlockImage{
		Type: slackBlockImageType,
	}
}

// SlackBlockImage represents a BlockUI image element
type SlackBlockImage struct {
	Type     string          `json:"type"`
	Title    *BlockTitleText `json:"title,omitempty"`
	ImageURL string          `json:"image_url,omitempty"`
	AltText  string          `json:"alt_text,omitempty"`
}

// GetType implements the SlackBlock interface and returns the type as string
func (i SlackBlockImage) GetType() string {
	return i.Type
}

func NewSlackBlockAccessoryImage() *SlackBlockAccessoryImage {
	return &SlackBlockAccessoryImage{
		Type: slackBlockImageType,
	}
}

// SlackBlockAccessoryImage represents a BlockUI image accessory element
type SlackBlockAccessoryImage struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
}
