package blockui

import (
	"fmt"
)

type SectionAccessory interface {
	HasInteraction() bool
	IsImage() bool
	SetText(*BlockTitleText)
	SetValue(string)
}

// SlackBlockSection represents a section Block Kit UI element
type SlackBlockSection struct {
	Type      string            `json:"type"`
	Text      TitleText         `json:"text,omitempty"`
	Accessory SectionAccessory  `json:"accessory,omitempty"`
	Fields    []*BlockTitleText `json:"fields,omitempty"`
}

// GetType implements SlackBlock interface
func (s SlackBlockSection) GetType() string {
	return s.Type
}

// SetText sets the Text field as appropriate depending on if the section has an accessory or not
func (s *SlackBlockSection) SetText(textType string, textVal string) {
	var tt TitleText
	if s.HasAccessory() {
		tt = NewBlockTitleTextEmojiless(textType)
	} else {
		tt = NewBlockTitleText(textType)
	}
	tt.SetText(textVal)
	s.Text = tt
}

func (s SlackBlockSection) HasAccessory() bool {
	return s.Fields != nil || s.Accessory != nil
}

// NewSlackBlockSection creates a new empty section UI element
func NewSlackBlockSection() *SlackBlockSection {
	return &SlackBlockSection{
		Type: slackBlockSectionType,
	}
}

// NewSlackBlockSectionWithAccessory creates a new Block Kit section UI element with provided accesssory type
func NewSlackBlockSectionWithAccessory(accessoryType string) (*SlackBlockSection, error) {
	err := fmt.Errorf("NewSlackBlockSectionWithAccessory requires a valid accessory type")
	if accessoryType == "" {
		return nil, err
	}
	s := NewSlackBlockSection()

	switch accessoryType {
	case slackAccessoryButtonType:
		s.Accessory = NewBlockButton()
	}

	return s, nil
}
