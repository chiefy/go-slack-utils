package payload

import (
	"encoding/json"
	"github.com/chiefy/go-slack-utils/pkg/blockui"
)

// BlockActionsContainer is the payload which comes from a Block Kit UI interaction
type BlockActionsContainer struct {
	Type      string `json:"type"`
	MessageTS string `json:"message_ts,,omitempty"`
	ChannelID string `json:"channel_id,,omitempty"`
	Ephemeral bool   `json:"is_ephemeral,omitempty"`
}

// BlockActionsPayload is the payload which comes from a Block Kit UI interaction
type BlockActionsPayload struct {
	Type        string                  `json:"type"`
	Team        map[string]string       `json:"team"`
	User        map[string]string       `json:"user"`
	Channel     map[string]string       `json:"channel"`
	APIAppID    string                  `json:"api_app_id"`
	Token       string                  `json:"token"`
	Container   *BlockActionsContainer  `json:"container"`
	TriggerID   string                  `json:"trigger_id"`
	ResponseURL string                  `json:"response_url"`
	Actions     []blockui.ActionElement `json:"actions"`
}

// UnmarshalJSON custom unmarshaller
func (b *BlockActionsPayload) UnmarshalJSON(data []byte) error {
	var objMap map[string]*json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["type"], &b.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["api_app_id"], &b.APIAppID); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["token"], &b.Token); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["trigger_id"], &b.TriggerID); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["response_url"], &b.ResponseURL); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["team"], &b.Team); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["container"], &b.Container); err != nil {
		return err
	}
	if err := json.Unmarshal(*objMap["channel"], &b.Channel); err != nil {
		return err
	}
	actions := make([]map[string]*json.RawMessage, 0)
	if err := json.Unmarshal(*objMap["actions"], &actions); err != nil {
		return err
	}
	if len(actions) > 0 {
		var rawActions []*json.RawMessage
		if err := json.Unmarshal(*objMap["actions"], &rawActions); err != nil {
			return err
		}
		for i, a := range actions {
			var actionType string
			if err := json.Unmarshal(*a["type"], &actionType); err != nil {
				return err
			}
			switch actionType {
			case "static_select":
				action := blockui.NewBlockSelect()
				if err := json.Unmarshal(*rawActions[i], &action); err != nil {
					return err
				}

				b.Actions = append(b.Actions, action)
			}
		}
	} else {
		b.Actions = nil
	}
	return nil
}
