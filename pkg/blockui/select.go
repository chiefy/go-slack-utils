package blockui

// BlockOption is a simple struct with a text / title field and value
type BlockOption struct {
	Text  *BlockTitleText `json:"text"`
	Value string          `json:"value"`
}

func NewBlockSelect() *BlockSelect {
	return &BlockSelect{
		Type:    slackAccessoryStaticSelectType,
		Options: []*BlockOption{},
	}
}

// BlockSelect represents a Block Kit select UI element
type BlockSelect struct {
	Type           string          `json:"type"`
	Placeholder    *BlockTitleText `json:"placeholder,omitempty"`
	Options        []*BlockOption  `json:"options,omitempty"`
	ActionID       string          `json:"action_id,omitempty"`
	BlockID        string          `json:"block_id,omitempty"`
	SelectedOption *BlockOption    `json:"selected_option,omitempty"`
}

func (s *BlockSelect) AddOption(o *BlockOption) {
	s.Options = append(s.Options, o)
}

func (s *BlockSelect) SetText(t *BlockTitleText) {
	s.Placeholder = t
}

func (s BlockSelect) UsesPlaceholder() bool {
	return true
}

func (s BlockSelect) GetPlaceholder() *BlockTitleText {
	return s.Placeholder
}

func (s BlockSelect) HasOptions() bool {
	return true
}

func (s BlockSelect) IsImage() bool {
	return false
}

func (s BlockSelect) HasInteraction() bool {
	return true
}

func (s BlockSelect) GetOptions() []*BlockOption {
	return s.Options
}

func (s BlockSelect) GetActionID() string {
	return s.ActionID
}

func (s BlockSelect) GetValue() string {
	if s.SelectedOption == nil {
		return ""
	}
	return s.SelectedOption.Value
}

func (s *BlockSelect) SetValue(v string) {

}
