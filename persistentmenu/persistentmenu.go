package persistentmenu

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

func SetPersistentMenu(accessToken string, menu []models.PersistentMenu) error {
	m := struct {
		PersistentMenu []models.PersistentMenu `json:"persistent_menu"`
	}{PersistentMenu: menu}

	b, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print(string(b))

	req, err := http.NewRequest("POST", EndPoint, bytes.NewBuffer(b))
	if err != nil {
		log.Print(err)
		return err
	}

	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}

	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Print(err)
		return err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
		return err
	}
	log.Print(result)
	return nil
}

func DeletePersistendMenu(accessToken string) error {
	m := map[string][]string{
		"fields": []string{
			"get_started",
			"persistent_menu",
			"whitelisted_domains",
		},
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print(string(b))

	req, err := http.NewRequest("DELETE", EndPoint, bytes.NewBuffer(b))
	if err != nil {
		log.Print(err)
		return err
	}

	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}

	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Print(err)
		return err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
		return err
	}
	log.Print(result)
	return nil
}
