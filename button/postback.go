package button

import "github.com/uzimaru0000/messengerbot/models"

type postBackButton struct {
	ButtonType models.ButtonType `json:"type"`
	Title      string            `json:"title"`
	Payload    string            `json:"payload"`
}

// NewPostBackButton is instancing call button
func NewPostBackButton(title string, payload string) models.Button {
	return &postBackButton{ButtonType: models.PostBackButton, Title: title, Payload: payload}
}

func (b *postBackButton) GetType() models.ButtonType {
	return b.ButtonType
}
