package template

import (
	"github.com/uzimaru0000/messengerbot/models"
)

func NewTemplate(senderID string, payload *models.Payload) *models.SendMessage {
	recipient := &models.Recipient{ID: senderID}
	sm := &models.SendMessage{}
	sm.Recipient = recipient
	a := &models.Attachment{
		Type:    "template",
		Payload: payload,
	}
	m := &models.SendingMessage{Attachment: a}
	sm.Message = m

	return sm
}
