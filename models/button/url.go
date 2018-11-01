package button

import "github.com/uzimaru0000/messengerbot/models/modifire"

type urlButton struct {
	ButtonType          ButtonType                  `json:"type"`
	Title               string                      `json:"title"`
	URL                 string                      `json:"url"`
	WebviewHeightRatio  modifire.WebviewHeightRatio `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool                        `json:"messenger_extensions,omitempty"`
	FallbackURL         string                      `json:"fallback_url,omitempty"`
	WebviewShareButton  string                      `json:"webview_share_button,omitempty"`
}

type UrlButtonOption func(*urlButton)

func WithWebviewHeightRatio(ratio modifire.WebviewHeightRatio) UrlButtonOption {
	return func(b *urlButton) {
		b.WebviewHeightRatio = ratio
	}
}

func WithMessengerExtensions(flag bool) UrlButtonOption {
	return func(b *urlButton) {
		b.MessengerExtensions = flag
	}
}

func WithFallbackURL(url string) UrlButtonOption {
	return func(b *urlButton) {
		b.FallbackURL = url
	}
}

func WithWebviewShareButton(flag bool) UrlButtonOption {
	return func(b *urlButton) {
		if !flag {
			b.WebviewShareButton = "hide"
		}
	}
}

func NewURLButton(title string, url string, opts ...UrlButtonOption) Button {
	button := &urlButton{ButtonType: URL, Title: title, URL: url}

	for _, opt := range opts {
		opt(button)
	}

	return button
}

func (b *urlButton) GetType() ButtonType {
	return b.ButtonType
}
