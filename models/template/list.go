package template

import "github.com/uzimaru0000/messengerbot/models/button"

type listTemplate struct {
	TemplateType    TemplateType    `json:"type"`
	TopElementStyle string          `json:"top_element_style,omitempty"`
	Buttons         []button.Button `json:"buttons,omitempty"`
	Elements        []Element       `json:"elements"`
	Sharable        bool            `json:"sharable,omitempty"`
}

type ListTemplateOption func(*listTemplate)

func (t *listTemplate) GetType() TemplateType {
	return t.TemplateType
}

func (t *listTemplate) WithTopElementStyle(style string) ListTemplateOption {
	return func(t *listTemplate) {
		t.TopElementStyle = style
	}
}

func (t *listTemplate) WithButtons(btns []button.Button) ListTemplateOption {
	return func(t *listTemplate) {
		t.Buttons = btns
	}
}

func (t *listTemplate) WithSharable(flag bool) ListTemplateOption {
	return func(t *listTemplate) {
		t.Sharable = flag
	}
}

func NewListTemplate(elements []Element, opts ...ListTemplateOption) Template {
	t := &listTemplate{Elements: elements}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
