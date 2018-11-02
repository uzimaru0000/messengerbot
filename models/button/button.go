package button

import "encoding/json"

type ButtonType int

const (
	URL ButtonType = iota + 1
	Share
	PostBack
	Buy
	Call
	GamePlay
	LogIn
	LogOut
)

func (b ButtonType) String() string {
	switch b {
	case URL:
		return "web_url"
	case Share:
		return "element_share"
	case PostBack:
		return "postback"
	case Buy:
		return "payment"
	case Call:
		return "phone_number"
	case GamePlay:
		return "game_play"
	case LogIn:
		return "account_link"
	case LogOut:
		return "account_unlink"
	default:
		return ""
	}
}

func (b ButtonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

type Button interface {
	GetType() ButtonType
}
