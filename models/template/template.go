package template

import (
	"encoding/json"

	"github.com/uzimaru0000/messengerbot/models/button"
	"github.com/uzimaru0000/messengerbot/models/modifire"
)

type TemplateType int

type Template interface {
	GetType() TemplateType
}

type TemplateOption func(*Template)

type Element struct {
	Title         string          `json:"title"`
	SubTitle      string          `json:"sub_title,omitempty"`
	ImageURL      string          `json:"image_url,omitempty"`
	DefaultAction DefaultAction   `json:"default_action,omitempty"`
	Buttons       []button.Button `json:"buttons,omitempty"`
}

type DefaultAction struct {
	URL                 string                      `json:"url"`
	WebviewHeightRatio  modifire.WebviewHeightRatio `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool                        `json:"messenger_extensions,omitempty"`
	FallbackURL         string                      `json:"fallback_url,omitempty"`
	WebviewShareButton  string                      `json:"webview_share_button,omitempty"`
}

const (
	Generic TemplateType = iota + 1
	Button
	List
	Media
	Receipt
)

func (t TemplateType) String() string {
	switch t {
	case Generic:
		return "generic"
	case Button:
		return "button"
	case List:
		return "list"
	case Media:
		return "media"
	case Receipt:
		return "receipt"
	default:
		return "generic"
	}
}

func (t TemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}
