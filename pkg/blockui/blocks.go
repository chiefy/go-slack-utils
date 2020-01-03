package blockui

// SlackBlock represents a BlockUI element
type SlackBlock interface {
	GetType() string
}

// SlackBlocks is a collection of SlackBlock UI elements
type SlackBlocks struct {
	Blocks []SlackBlock `json:"blocks"`
}

// NewSlackBlocks creates a new BlockUI root structure
func NewSlackBlocks() *SlackBlocks {
	return &SlackBlocks{
		Blocks: []SlackBlock{},
	}
}

// Push adds a new BlockUI element to the root blocks collection
func (b *SlackBlocks) Push(block SlackBlock) {
	b.Blocks = append(b.Blocks, block)
}

// NumBlocks returns the current length of child blocks
func (b SlackBlocks) NumBlocks() int {
	return len(b.Blocks)
}
