package models

import "encoding/json"

type Button interface {
	GetType() ButtonType
}

type ButtonType int

const (
	URLButton ButtonType = iota + 1
	ShareButton
	PostBackButton
	BuyButton
	CallButton
	GamePlayButton
	LogInButton
	LogOutButton
	Nested
)

func (b ButtonType) String() string {
	switch b {
	case URLButton:
		return "web_url"
	case ShareButton:
		return "element_share"
	case PostBackButton:
		return "postback"
	case BuyButton:
		return "payment"
	case CallButton:
		return "phone_number"
	case GamePlayButton:
		return "game_play"
	case LogInButton:
		return "account_link"
	case LogOutButton:
		return "account_unlink"
	case Nested:
		return "nested"
	default:
		return ""
	}
}

func (b ButtonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}
