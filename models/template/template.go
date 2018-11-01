package template

import "encoding/json"

type TemplateType int

const (
	Generic TemplateType = iota
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

type Template interface {
	GetType() TemplateType
}
