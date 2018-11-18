package template

import (
	"encoding/json"

	"github.com/uzimaru0000/messengerbot/models"
)

type mediaTemplate struct {
	TemplateType models.TemplateType `json:"template_type"`
	Elements     []MediaElement      `json:"elements"`
	Sharable     bool                `json:"sharable,omitempty"`
}

// MediaElement is element which MediaTemplate
type MediaElement struct {
	MediaType    MediaType
	AttachmentID string
	URL          string
	buttons      []models.Button
}

// MediaTemplateOption is type of function which set option to MediaTemplate
type MediaTemplateOption func(*mediaTemplate)

// MediaType is type of media to image or video.
type MediaType int

// Image -> graphics, Video -> movies
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

func (t *mediaTemplate) SetSharable(flag bool) {
	t.Sharable = flag
}

func (t *mediaTemplate) GetType() models.TemplateType {
	return t.TemplateType
}

// NewMediaTemplate is create MediaTemplate
func NewMediaTemplate(elements []MediaElement, opts ...MediaTemplateOption) models.Template {
	t := &mediaTemplate{TemplateType: models.MediaTemplate, Elements: elements}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
