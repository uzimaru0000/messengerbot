package button

import (
	"github.com/uzimaru0000/messengerbot/models"
	"github.com/uzimaru0000/messengerbot/models/modifire"
)

type urlButton struct {
	ButtonType          models.ButtonType           `json:"type"`
	Title               string                      `json:"title"`
	URL                 string                      `json:"url"`
	WebviewHeightRatio  modifire.WebviewHeightRatio `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool                        `json:"messenger_extensions,omitempty"`
	FallbackURL         string                      `json:"fallback_url,omitempty"`
	WebviewShareButton  string                      `json:"webview_share_button,omitempty"`
}

// URLButtonOption is type of function which set option to UrlButton
type URLButtonOption func(*urlButton)

// WithWebviewHeightRatio is setting WebviewHeightRatio to UrlButton
func WithWebviewHeightRatio(ratio modifire.WebviewHeightRatio) URLButtonOption {
	return func(b *urlButton) {
		b.WebviewHeightRatio = ratio
	}
}

// WithMessengerExtensions is setting MessengerExtensions to UrlButton
func WithMessengerExtensions(flag bool) URLButtonOption {
	return func(b *urlButton) {
		b.MessengerExtensions = flag
	}
}

// WithFallbackURL is setting FallbackURL to UrlButton
func WithFallbackURL(url string) URLButtonOption {
	return func(b *urlButton) {
		b.FallbackURL = url
	}
}

// WithWebviewShareButton is setting WebviewShareButton to UrlButton
func WithWebviewShareButton(flag bool) URLButtonOption {
	return func(b *urlButton) {
		if !flag {
			b.WebviewShareButton = "hide"
		}
	}
}

// NewURLButton is instancing call button
func NewURLButton(title string, url string, opts ...URLButtonOption) models.Button {
	button := &urlButton{ButtonType: models.URLButton, Title: title, URL: url}

	for _, opt := range opts {
		opt(button)
	}

	return button
}

func (b *urlButton) GetType() models.ButtonType {
	return b.ButtonType
}
