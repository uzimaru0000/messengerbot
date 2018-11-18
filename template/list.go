package template

import "github.com/uzimaru0000/messengerbot/models"

type listTemplate struct {
	TemplateType    models.TemplateType `json:"template_type"`
	TopElementStyle string              `json:"top_element_style,omitempty"`
	Buttons         []models.Button     `json:"buttons,omitempty"`
	Elements        []models.Element    `json:"elements"`
	Sharable        bool                `json:"sharable,omitempty"`
}

// ListTemplateOption is type of function which set option to ListTemplate
type ListTemplateOption func(*listTemplate)

func (t *listTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

// WithTopElementStyle is setting TopElementStyle to ListTemplate
func WithTopElementStyle(style string) ListTemplateOption {
	return func(t *listTemplate) {
		t.TopElementStyle = style
	}
}

// WithButtons is setting Buttons to ListTemplate
func WithButtons(btns []models.Button) ListTemplateOption {
	return func(t *listTemplate) {
		t.Buttons = btns
	}
}

func (t *listTemplate) SetSharable(flag bool) {
	t.Sharable = flag
}

// NewListTemplate is create ButtonTemplate
func NewListTemplate(elements []models.Element, opts ...ListTemplateOption) models.Template {
	t := &listTemplate{Elements: elements}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
