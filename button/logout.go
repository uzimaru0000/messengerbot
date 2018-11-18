package button

import "github.com/uzimaru0000/messengerbot/models"

type logoutButton struct {
	ButtonType models.ButtonType `json:"type"`
}

func (b *logoutButton) GetType() models.ButtonType {
	return b.ButtonType
}

// NewLogoutButton is instancing call button
func NewLogoutButton() models.Button {
	return &logoutButton{ButtonType: models.LogOutButton}
}
