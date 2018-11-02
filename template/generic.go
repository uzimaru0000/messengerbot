package template

import (
	"github.com/uzimaru0000/messengerbot/models"
	"github.com/uzimaru0000/messengerbot/models/modifire"
)

type genericTemplate struct {
	TemplateType     models.TemplateType       `json:"template_type"`
	Elements         []models.Element          `json:"elements"`
	Sharable         bool                      `json:"sharable,omitempty"`
	ImageAspectRatio modifire.ImageAspectRatio `json:"image_aspect_ratio,omitempty"`
}

type GenericTemplateOption func(*genericTemplate)

func (t *genericTemplate) WithSharable(flag bool) GenericTemplateOption {
	return func(t *genericTemplate) {
		t.Sharable = flag
	}
}

func (t *genericTemplate) WithImageAspectRatio(ratio modifire.ImageAspectRatio) GenericTemplateOption {
	return func(t *genericTemplate) {
		t.ImageAspectRatio = ratio
	}
}

func (t *genericTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

func NewGenericTemplate(elements []models.Element, opts ...GenericTemplateOption) models.Template {
	template := &genericTemplate{TemplateType: models.GenericTemplate, Elements: elements}

	for _, opt := range opts {
		opt(template)
	}

	return template
}
