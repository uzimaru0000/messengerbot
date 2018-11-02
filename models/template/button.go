package template

import "github.com/uzimaru0000/messengerbot/models/button"

type buttonTemplate struct {
	TemplateType TemplateType    `json:"type"`
	Text         string          `json:"text"`
	Buttons      []button.Button `json:"buttons"`
	Sharable     bool            `json:"sharable,omitempty"`
}

type ButtonTemplateOption func(*buttonTemplate)

func (t *buttonTemplate) WithSharable(flag bool) ButtonTemplateOption {
	return func(t *buttonTemplate) {
		t.Sharable = flag
	}
}

func (t *buttonTemplate) GetType() TemplateType {
	return t.TemplateType
}

func NewButtonTemplate(text string, buttons []button.Button, opts ...ButtonTemplateOption) Template {
	t := &buttonTemplate{TemplateType: Button, Text: text, Buttons: buttons}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
