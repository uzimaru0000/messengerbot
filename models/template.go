package models

import (
	"encoding/json"
)

type TemplateType int

type Template interface {
	GetType() TemplateType
	SetSharable(bool)
}

type TemplateOption func(*Template)

const (
	GenericTemplate TemplateType = iota + 1
	ButtonTemplate
	ListTemplate
	MediaTemplate
	ReceiptTemplate
)

func (t TemplateType) String() string {
	switch t {
	case GenericTemplate:
		return "generic"
	case ButtonTemplate:
		return "button"
	case ListTemplate:
		return "list"
	case MediaTemplate:
		return "media"
	case ReceiptTemplate:
		return "receipt"
	default:
		return "generic"
	}
}

func (t TemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}
