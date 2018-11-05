package persistentmenu

import "github.com/uzimaru0000/messengerbot/models"

type nestedItem struct {
	ItemType      models.MenuItemType `json:"type"`
	Title         string              `json:"title"`
	CallToActions []models.MenuItem   `json:"call_to_actions"`
}

func (i *nestedItem) GetType() models.MenuItemType {
	return i.ItemType
}

func NewNestedItem(title string, actions []models.MenuItem) models.MenuItem {
	return &nestedItem{ItemType: models.Nested, Title: title, CallToActions: actions}
}
