package template

import (
	"github.com/uzimaru0000/messengerbot/models/modifire"
)

type genericTemplate struct {
	TemplateType     TemplateType              `json:"type"`
	Elements         []Element                 `json:"elements"`
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

func (t *genericTemplate) GetType() TemplateType {
	return t.TemplateType
}

func NewGenericTemplate(elements []Element, opts ...GenericTemplateOption) Template {
	template := &genericTemplate{TemplateType: Generic, Elements: elements}

	for _, opt := range opts {
		opt(template)
	}

	return template
}
