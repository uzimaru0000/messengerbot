package button

import "github.com/uzimaru0000/messengerbot/models"

type gamePlayButton struct {
	ButtonType   models.ButtonType `json:"type"`
	Title        string            `json:"title"`
	Payload      string            `json: "payload"`
	GameMetaData *GameMetaData     `json:"game_metadata"`
}

type GamePlayOption func(*gamePlayButton)

type GameMetaData struct {
	PlayerID  string `json:"player_id"`
	ContextID string `json:"context_id"`
}

func WithPayload(payload string) GamePlayOption {
	return func(b *gamePlayButton) {
		b.Payload = payload
	}
}

func WithGameMetaData(metaData *GameMetaData) GamePlayOption {
	return func(b *gamePlayButton) {
		b.GameMetaData = metaData
	}
}

func NewGamePlayButton(title string, options ...GamePlayOption) models.Button {
	button := &gamePlayButton{ButtonType: models.GamePlayButton, Title: title}
	for _, opt := range options {
		opt(button)
	}
	return button
}

func (b *gamePlayButton) GetType() models.ButtonType {
	return b.ButtonType
}
