package blockui

// BlockOption is a simple struct with a text / title field and value
type BlockOption struct {
	Text  *BlockTitleText `json:"text"`
	Value string          `json:"value"`
}

func NewSelect() *BlockSelect {
	return &BlockSelect{
		Type:    slackAccessoryStaticSelectType,
		Options: []*BlockOption{},
	}
}

// BlockSelect represents a Block Kit select UI element
type BlockSelect struct {
	Type        string          `json:"type"`
	Placeholder *BlockTitleText `json:"placeholder,omitempty"`
	Options     []*BlockOption  `json:"options,omitempty"`
	ActionID    string          `json:"action_id,omitempty"`
}

func (s BlockSelect) HasPlaceholder() bool {
	return true
}

func (s BlockSelect) GetPlaceholder() *BlockTitleText {
	return s.Placeholder
}

func (s BlockSelect) HasOptions() bool {
	return true
}

func (s BlockSelect) GetOptions() []*BlockOption {
	return s.Options
}
