package blockui

// NewBlockImage creates a new BlockImage UI element with provided URL and alt text
func NewBlockImage(imageURL string, altText string) *BlockImage {
	return &BlockImage{
		Type:     slackBlockImageType,
		ImageURL: imageURL,
		AltText:  altText,
	}
}

// BlockImage represents a BlockUI image element
type BlockImage struct {
	Type     string          `json:"type"`
	Title    *BlockTitleText `json:"title,omitempty"`
	ImageURL string          `json:"image_url,omitempty"`
	AltText  string          `json:"alt_text,omitempty"`
}

// GetType implements the SlackBlock interface and returns the type as string
func (i BlockImage) GetType() string {
	return i.Type
}

func (i BlockImage) HasInteraction() bool {
	return false
}

func (i BlockImage) IsImage() bool {
	return true
}

func (i *BlockImage) SetText(t *BlockTitleText) {
	i.Title = t
}

// NewBlockAccessoryImage creates a new BlockAccessoryImage UI element
func NewBlockAccessoryImage(imageURL string, altText string) *BlockAccessoryImage {
	return &BlockAccessoryImage{
		Type:     slackBlockImageType,
		ImageURL: imageURL,
		AltText:  altText,
	}
}

// BlockAccessoryImage represents a BlockUI image accessory element
type BlockAccessoryImage struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
}

func (i BlockAccessoryImage) HasInteraction() bool {
	return false
}

func (i BlockAccessoryImage) IsImage() bool {
	return true
}

func (i *BlockAccessoryImage) SetText(text *BlockTitleText) {
	i.AltText = text.GetText()
}

func (i *BlockAccessoryImage) SetValue(url string) {
	i.ImageURL = url
}
