package modifire

import "encoding/json"

type WebviewHeightRatio int

const (
	Compact WebviewHeightRatio = iota
	Tall
	Full
)

func (w WebviewHeightRatio) String() string {
	switch w {
	case Compact:
		return "compact"
	case Tall:
		return "tall"
	case Full:
		return "full"
	default:
		return "compact"
	}
}

func (w WebviewHeightRatio) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String())
}
