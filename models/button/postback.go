package button

type postBackButton struct {
	ButtonType ButtonType `json:"type"`
	Title      string     `json:"title"`
	Payload    string     `json:"payload"`
}

func NewPostBackButton(title string, payload string) Button {
	return &postBackButton{PostBack, title, payload}
}

func (b *postBackButton) GetType() ButtonType {
	return b.ButtonType
}
