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

// GenericTemplateOption is type of function which set option to ListTemplate
type GenericTemplateOption func(*genericTemplate)

func (t *genericTemplate) SetSharable(flag bool) {
	t.Sharable = flag
}

func (t *genericTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

// WithImageAspectRatio is setting ImageAspectRatio to GenericTemplate
func WithImageAspectRatio(ratio modifire.ImageAspectRatio) GenericTemplateOption {
	return func(t *genericTemplate) {
		t.ImageAspectRatio = ratio
	}
}

// NewGenericTemplate is create ButtonTemplate
func NewGenericTemplate(elements []models.Element, opts ...GenericTemplateOption) models.Template {
	template := &genericTemplate{TemplateType: models.GenericTemplate, Elements: elements}

	for _, opt := range opts {
		opt(template)
	}

	return template
}
