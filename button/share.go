package button

import "github.com/uzimaru0000/messengerbot/models"

type shareButton struct {
	ButtonType    models.ButtonType
	ShareContents models.Template
}

func (b *shareButton) GetType() models.ButtonType {
	return b.ButtonType
}

// NewShareButton is instancing call button
func NewShareButton(contents models.Template) models.Button {
	return &shareButton{ButtonType: models.ShareButton, ShareContents: contents}
}
