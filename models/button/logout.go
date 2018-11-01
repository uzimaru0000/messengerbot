package button

type logoutButton struct {
	ButtonType ButtonType `json:"type"`
}

func (b *logoutButton) GetType() ButtonType {
	return b.ButtonType
}

func NewLogoutButton() Button {
	return &logoutButton{LogOut}
}
