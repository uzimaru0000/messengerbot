package models

import "encoding/json"

type PersistentMenu struct {
	Locale                string     `json:"locale"`
	ComposerInputDisabled bool       `json:"composer_input_disabled"`
	CallToActions         []MenuItem `json:"call_to_actions"`
}

type MenuItemType int

const (
	WebURL MenuItemType = iota
	PostBack
	Nested
)

func (t MenuItemType) String() string {
	switch t {
	case WebURL:
		return "web_url"
	case PostBack:
		return "postback"
	case Nested:
		return "nested"
	default:
		return "web_url"
	}
}

func (t MenuItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

type MenuItem interface {
	GetType() MenuItemType
}
