package blockui

type ActionElement interface {
	GetActionID() string
	UsesPlaceholder() bool
	GetValue() string
	HasOptions() bool
	GetOptions() []*BlockOption
}

// ActionsBlock represents a Block Kit actions block UI element
type ActionsBlock struct {
	Type     string          `json:"type"`
	BlockID  string          `json:"block_id,omitempty"`
	Elements []ActionElement `json:"elements"`
}

// NewActionsBlock creates a new ActionsBlock struct with appropriate type field
func NewActionsBlock(blockID string) *ActionsBlock {
	return &ActionsBlock{
		Type:     slackBlockActionsType,
		BlockID:  blockID,
		Elements: []ActionElement{},
	}
}

func (a ActionsBlock) GetType() string {
	return a.Type
}

func (a *ActionsBlock) Push(e ActionElement) {
	a.Elements = append(a.Elements, e)
}

func (a ActionsBlock) NumElements() int {
	return len(a.Elements)
}
