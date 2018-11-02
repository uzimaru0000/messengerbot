package modifire

import "encoding/json"

type WebviewHeightRatio int

const (
	Compact WebviewHeightRatio = iota + 1
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

type ImageAspectRatio int

const (
	Horizontal ImageAspectRatio = iota + 1
	Square
)

func (i ImageAspectRatio) String() string {
	switch i {
	case Horizontal:
		return "horizontal"
	case Square:
		return "square"
	default:
		return "horizontal"
	}
}

func (i ImageAspectRatio) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}
