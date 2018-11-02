package template

import (
	"github.com/uzimaru0000/messengerbot/models"
	"github.com/uzimaru0000/messengerbot/models/template"
)

func NewTemplate(senderID string, template *template.Template) *models.SendMessage {
	recipient := &models.Recipient{ID: senderID}
	sm := &models.SendMessage{}
	sm.Recipient = recipient
	a := &models.Attachment{
		Type:    "template",
		Payload: template,
	}
	m := &models.SendingMessage{Attachment: a}
	sm.Message = m

	return sm
}
