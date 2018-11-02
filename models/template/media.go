package template

import (
	"encoding/json"

	"github.com/uzimaru0000/messengerbot/models/button"
)

type mediaTemplate struct {
	TemplateType TemplateType   `json:"type"`
	Elements     []MediaElement `json:"elements"`
	Sharable     bool           `json:"sharable,omitempty"`
}

type MediaElement struct {
	MediaType    MediaType
	AttachmentID string
	URL          string
	buttons      []button.Button
}

type MediaTemplateOption func(*mediaTemplate)

type MediaType int

const (
	Image MediaType = iota + 1
	Video
)

func (m MediaType) String() string {
	switch m {
	case Image:
		return "image"
	case Video:
		return "video"
	default:
		return ""
	}
}

func (m MediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (t *mediaTemplate) WithSharable(flag bool) MediaTemplateOption {
	return func(t *mediaTemplate) {
		t.Sharable = flag
	}
}

func (t *mediaTemplate) GetType() TemplateType {
	return t.TemplateType
}

func NewMediaTemplate(elements []MediaElement, opts ...MediaTemplateOption) Template {
	t := &mediaTemplate{TemplateType: Media, Elements: elements}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
