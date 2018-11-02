package models

import (
	"github.com/uzimaru0000/messengerbot/models/modifire"
)

//ReceivedMessage
type ReceivedMessage struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID        string      `json:"id"`
	Time      int         `json:"time"`
	Messaging []Messaging `json:"messaging"`
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
	MID         string       `json:"mid"`
	Seq         int          `json:"seq"`
	Text        string       `json:"text"`
	QuickReply  *QuickReply  `json:"quick_reply"`
	Attachments []Attachment `json:"attachments"`
}

type QuickReply struct {
	Payload string `json:"payload"`
}

type Attachment struct {
	Type    string    `json:"type"`
	Payload *Template `json:"payload"`
}

type Payload struct {
	Coordinates      *Coordinates `json:"coordinates,omitempty"`
	TemplateType     string       `json:"template_type,omitempty"`
	TopElementStyle  string       `json:"top_element_style,omitempty"`
	Sharable         bool         `json:"sharable,omitempty"`
	ImageAspectRatio string       `json:"image_aspect_ratio,omitempty"`
	Elements         []Element    `json:"elements,omitempty"`
}

type Element struct {
	Title         string         `json:"title"`
	Subtitle      string         `json:"subtitle"`
	ImageURL      string         `json:"image_url"`
	DefaultAction *DefaultAction `json:"default_action"`
	Buttons       []Button       `json:"buttons"`
}

type DefaultAction struct {
	Type                string                      `json:"type"`
	Title               string                      `json:"title,omitempty"`
	URL                 string                      `json:"url"`
	MessengerExtensions bool                        `json:"messenger_extensions"`
	WebViewHeightRatio  modifire.WebviewHeightRatio `json:"webview_height_ratio"`
	FallBackURL         string                      `json:"fallback_url,omitempty"`
}

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"Long"`
}

//SendMessage
type SendMessage struct {
	MessagingType    string          `json:"messaging_type,omitempty"`
	Recipient        *Recipient      `json:"recipient,omitempty"`
	Message          *SendingMessage `json:"message,omitempty"`
	SenderAction     string          `json:"sender_action,omitempty"`
	NotificationType string          `json:"notification_type,omitempty"`
	Tag              string          `json:"tag,omitempty"`
}

type SendingMessage struct {
	Text         string         `json:"text,omitempty"`
	QuickReplies []QuickReplies `json:"quick_replies,omitempty"`
	Attachment   *Attachment    `json:"attachment,omitempty"`
}

type QuickReplies struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Payload     string `json:"payload"`
	ImageURL    string `json:"image_url"`
}
