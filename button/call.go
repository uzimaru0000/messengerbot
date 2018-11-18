package button

import "github.com/uzimaru0000/messengerbot/models"

type callButton struct {
	ButtonType models.ButtonType `json:"type"`
	Title      string            `json:"title"`
	Payload    string            `json:"payload"`
}

// NewCallButton is instancing call button
func NewCallButton(title string, payload string) models.Button {
	return &callButton{ButtonType: models.CallButton, Title: title, Payload: payload}
}

func (b *callButton) GetType() models.ButtonType {
	return b.ButtonType
}
