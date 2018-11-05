package persistentmenu

import "github.com/uzimaru0000/messengerbot/models"

type postBack struct {
	ItemType models.MenuItemType `json:"type"`
	Title    string              `json:"title"`
	Payload  string              `json:"payload"`
}

func (i *postBack) GetType() models.MenuItemType {
	return i.ItemType
}

func NewPostBackItem(title string, payload string) models.MenuItem {
	i := &postBack{ItemType: models.PostBackItem, Title: title, Payload: payload}

	return i
}
