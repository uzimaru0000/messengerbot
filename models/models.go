package models

//ReceivedMessage
type ReceivedMessage struct {
	Object string   `json:"object"`
	Entry  *[]Entry `json:"entry"`
}

type Entry struct {
	ID        string       `json:"id"`
	Time      int          `json:"time"`
	Messaging *[]Messaging `json:"messaging"`
}

type Messaging struct {
	Sender    *Sender    `json:"sender"`
	Recipient *Recipient `json:"recipient"`
	Timestamp int        `json:"timestamp"`
	Message   *Message   `json:"message"`
}

type Sender struct {
	ID string `json:"id"`
}

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	MID         string        `json:"mid"`
	Seq         int           `json:"seq"`
	Text        string        `json:"text"`
	QuickReply  *Quick_reply  `json:"quick_reply"`
	Attachments *[]Attachment `json:"attachments"`
}

type Quick_reply struct {
	Payload string `json:"payload"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Coordinates      *Coordinates `json:"coordinates,omitempty"`
	TemplateType     string       `json:"template_type"`
	Sharable         bool         `json:"sharable"`
	ImageAspectRatio string       `json:"image_aspect_ratio,omitempty"`
	Elements         *[]Element   `json:"elements"`
}
type Element struct {
	Title         string          `json:"title"`
	Subtitle      string          `json:"subtitle"`
	ImageURL      string          `json:"image_url"`
	DefaultAction *Default_action `json:"default_action"`
	Buttons       *[]Button       `json:"buttons"`
}

type Default_action struct {
	Type                string `json:"type"`
	Title               string `json:"title,omitempty"`
	URL                 string `json:"url"`
	MessengerExtensions bool   `json:"messenger_extensions"`
	WebViewHeightRatio  string `json:"webview_height_ratio"`
	FallBackURL         string `json:"fallback_url,omitempty"`
}

type Button struct {
	Type                string `json:"type"`
	Title               string `json:"title"`
	URL                 string `json:"url"`
	WebViewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool   `json:"messenger_extensions,omitempty"`
	FallBackURL         string `json:"fallback_url,omitempty"`
	WebviewShareButton  string `json:"webview_share_button,omitempty"`
}

type Coordinates struct {
	Lat  float64 `json:"lat,omitempty"`
	Long float64 `json:"Long,omitempty"`
}

//SendMessage
type SendMessage struct {
	MessagingType    string     `json:"messaging_type"`
	Recipient        *Recipient `json:"recipient"`
	Message          *Message   `json:"message"`
	SenderAction     string     `json:"sender_action,omitempty"`
	NotificationType string     `json:"notification_type,omitempty"`
	Tag              string     `json:"tag,omitempty"`
}

type QuickReplies struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Payload     string `json:"payload"`
	ImageURL    string `json:"image_url"`
}
