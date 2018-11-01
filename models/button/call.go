package button

type callButton struct {
	ButtonType ButtonType `json:"type"`
	Title      string     `json:"title"`
	Payload    string     `json:"payload"`
}

func NewCallButton(title string, payload string) Button {
	return &callButton{Call, title, payload}
}

func (b *callButton) GetType() ButtonType {
	return b.ButtonType
}
