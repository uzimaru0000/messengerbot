package template

import (
	"github.com/uzimaru0000/messengerbot/models"
)

type buttonTemplate struct {
	TemplateType models.TemplateType `json:"type"`
	Text         string              `json:"text"`
	Buttons      []models.Button     `json:"buttons"`
	Sharable     bool                `json:"sharable,omitempty"`
}

type ButtonTemplateOption func(*buttonTemplate)

func (t *buttonTemplate) WithSharable(flag bool) ButtonTemplateOption {
	return func(t *buttonTemplate) {
		t.Sharable = flag
	}
}

func (t *buttonTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

func NewButtonTemplate(text string, buttons []models.Button, opts ...ButtonTemplateOption) models.Template {
	t := &buttonTemplate{TemplateType: models.ButtonTemplate, Text: text, Buttons: buttons}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
