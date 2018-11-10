package models

type Property interface {
	GetPropertyName() string
}

type Greeting struct {
	Locale string `json:"locale"`
	Text   string `json:"text"`
}

func (p *Greeting) GetPropertyName() string {
	return "greeting"
}

type AccountLink struct {
	URL string `json:"account_linking_url"`
}

func (p *AccountLink) delete() string {
	return "account_linking_url"
}

type PersistentMenu struct {
	Locale                string   `json:"locale"`
	ComposerInputDisabled bool     `json:"composer_input_disabled"`
	CallToActions         []Button `json:"call_to_actions"`
}

func (p *PersistentMenu) GetPropertyName() string {
	return "persistent_menu"
}

type Start struct {
	Payload string `json:"payload"`
}

func (p *Start) GetProrepryName() string {
	return "get_started"
}

type WhitelistDomain struct {
	domeins []string `json:"whitelisted_domains"`
}

func (p *WhitelistDomain) GetPropertyName() string {
	return "whitelisted_domains"
}
