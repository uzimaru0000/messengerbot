package template

import (
	"github.com/uzimaru0000/messengerbot/models"
)

type buttonTemplate struct {
	TemplateType models.TemplateType `json:"template_type"`
	Text         string              `json:"text"`
	Buttons      []models.Button     `json:"buttons"`
	Sharable     bool                `json:"sharable,omitempty"`
}

func (t *buttonTemplate) SetSharable(flag bool) {
	t.Sharable = flag
}

func (t *buttonTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

// NewButtonTemplate is create ButtonTemplate
func NewButtonTemplate(text string, buttons []models.Button) models.Template {
	t := &buttonTemplate{TemplateType: models.ButtonTemplate, Text: text, Buttons: buttons}

	return t
}
