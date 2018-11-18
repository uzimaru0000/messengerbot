package button

import "github.com/uzimaru0000/messengerbot/models"

type loginButton struct {
	ButtonType models.ButtonType `json:"type"`
	URL        string            `json:"url"`
}

func (b *loginButton) GetType() models.ButtonType {
	return b.ButtonType
}

// NewLoginButton is instancing call button
func NewLoginButton(url string) models.Button {
	return &loginButton{ButtonType: models.LogInButton, URL: url}
}
