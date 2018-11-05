package persistentmenu

import (
	"github.com/uzimaru0000/messengerbot/models"
	"github.com/uzimaru0000/messengerbot/models/modifire"
)

type webURL struct {
	ItemType            models.MenuItemType         `json:"type"`
	Title               string                      `json:"title"`
	URL                 string                      `json:"url"`
	WebviewHeightRatio  modifire.WebviewHeightRatio `json:"webview_height_ratio"`
	MessengerExtensions bool                        `json:"messenger_extensions"`
}

type WebURLOption func(*webURL)

func WithWebviewHeightRatio(ratio modifire.WebviewHeightRatio) WebURLOption {
	return func(i *webURL) {
		i.WebviewHeightRatio = ratio
	}
}

func (i *webURL) GetType() models.MenuItemType {
	return i.ItemType
}

func NewWebURLItem(title string, url string, opts ...WebURLOption) models.MenuItem {
	i := &webURL{ItemType: models.WebURL, Title: title, URL: url, MessengerExtensions: true}

	for _, opt := range opts {
		opt(i)
	}

	return i
}
