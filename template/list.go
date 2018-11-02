package template

import "github.com/uzimaru0000/messengerbot/models"

type listTemplate struct {
	TemplateType    models.TemplateType `json:"type"`
	TopElementStyle string              `json:"top_element_style,omitempty"`
	Buttons         []models.Button     `json:"buttons,omitempty"`
	Elements        []models.Element    `json:"elements"`
	Sharable        bool                `json:"sharable,omitempty"`
}

type ListTemplateOption func(*listTemplate)

func (t *listTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

func (t *listTemplate) WithTopElementStyle(style string) ListTemplateOption {
	return func(t *listTemplate) {
		t.TopElementStyle = style
	}
}

func (t *listTemplate) WithButtons(btns []models.Button) ListTemplateOption {
	return func(t *listTemplate) {
		t.Buttons = btns
	}
}

func (t *listTemplate) WithSharable(flag bool) ListTemplateOption {
	return func(t *listTemplate) {
		t.Sharable = flag
	}
}

func NewListTemplate(elements []models.Element, opts ...ListTemplateOption) models.Template {
	t := &listTemplate{Elements: elements}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
