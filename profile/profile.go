package profile

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/uzimaru0000/messengerbot/models"
)

const EndPoint = "https://graph.facebook.com/v2.6/me/messenger_profile"

type deleteProperty struct {
	Fields []string `json:"fields"`
}

type Properties struct {
	PersistentMenu []models.PersistentMenu `json:"persistent_menu"`
	*models.AccountLink
	*models.Start `json:"get_started"`
	Greetings     []models.Greeting `json:"greeting"`
	*models.WhitelistDomain
}

type SettingProperies struct {
	Data []map[string]interface{} `json:"data"`
}

func (p *deleteProperty) GetPropertyName() string {
	return ""
}

func send(method string, accessToken string, body []byte) ([]byte, error) {

	req, err := http.NewRequest(method, EndPoint, bytes.NewBuffer(body))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func SetProperties(accessToken string, properties *Properties) error {
	body, err := json.Marshal(properties)
	if err != nil {
		return err
	}

	_, err = send("POST", accessToken, body)

	return err
}

func DeleteProperties(accessToken string, propertis []models.Property) error {
	del := &deleteProperty{}

	for _, p := range propertis {
		del.Fields = append(del.Fields, p.GetPropertyName())

	}
	log.Print(del)

	body, err := json.Marshal(del)
	if err != nil {
		return err
	}

	log.Print(string(body))

	body, err = send("DELETE", accessToken, body)

	log.Print(string(body))

	return err
}
