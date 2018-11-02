package template

import (
	"github.com/uzimaru0000/messengerbot/models"
)

func NewTemplate(senderID string, template *models.Template) *models.SendMessage {
	recipient := &models.Recipient{ID: senderID}
	sm := &models.SendMessage{}
	sm.Recipient = recipient
	a := &models.TemplateAttachment{
		Type:     "template",
		Template: template,
	}
	m := &models.SendingMessage{Attachment: a}
	sm.Message = m

	return sm
}
