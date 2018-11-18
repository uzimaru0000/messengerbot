package button

import "github.com/uzimaru0000/messengerbot/models"

type nestedItem struct {
	ItemType      models.ButtonType `json:"type"`
	Title         string            `json:"title"`
	CallToActions []models.Button   `json:"call_to_actions"`
}

func (i *nestedItem) GetType() models.ButtonType {
	return i.ItemType
}

// NewNestedButton is instancing call button
func NewNestedButton(title string, actions []models.Button) models.Button {
	return &nestedItem{ItemType: models.Nested, Title: title, CallToActions: actions}
}
