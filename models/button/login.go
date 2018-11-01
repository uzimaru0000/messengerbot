package button

type loginButton struct {
	ButtonType ButtonType `json:"type"`
	URL        string     `json:"url"`
}

func (b *loginButton) GetType() ButtonType {
	return b.ButtonType
}

func NewLoginButton(url string) Button {
	return &loginButton{LogIn, url}
}
