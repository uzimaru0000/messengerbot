package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var accessToken = "EAAapjO8eyyoBANsd274NLUYfJixsO4GerD2moWkklHZCDMXLgWnvvrTaAOTE8iAHZAKXEGvZBDxZB3viYVjUbKDZCOx03M5OqTwbJ8OOkMgaJDgI2sMogHixalslOKKcdZBxuohZCxShdmJUZAsY0uI2pje9DLfTO4pyFMM1MJ6HdV6xZCwfaINPg"
var verifyToken = "3460"

const (
	EndPoint = "https://graph.facebook.com/v2.6/me/messages"
)

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
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int       `json:"timestamp"`
	Message   Message   `json:"message"`
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
	Quick_reply Quick_reply   `json:"quick_reply"`
	Attachments []Attachments `json:"attachments"`
}

type Quick_reply struct {
	Payload string `json:"payload"`
}

type Attachments struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"Long"`
}

type SendMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   struct {
		Text          string          `json:"text"`
		Quick_replies []Quick_replies `json:"quick_replies"`
	} `json:"message"`
}

type Quick_replies struct {
	Content_type string `json:"content_type"`
	Title        string `json:"title"`
	Payload      string `json:"payload"`
	Image_url    string `json:"image_url"`
}

type TalkJson struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Results []talkApiResult `json:"results"`
}

type talkApiResult struct {
	Perplexity float64 `json:"perplexity"`
	Reply      string  `json:"reply"`
}

func main() {
	http.HandleFunc("/", TopPageHandler)
	http.HandleFunc("/webhook", webhookHandler)
	port := "5000"
	address := fmt.Sprintf(":%s", port)
	http.ListenAndServe(address, nil)
}

func TopPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is go-bot application's top page.")
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		verifyTokenAction(w, r)
	}
	if r.Method == "POST" {
		webhookPostAction(w, r)
	}
}

func verifyTokenAction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("hub.verify_token") == verifyToken {
		log.Print("verify token success.")
		fmt.Fprintf(w, r.URL.Query().Get("hub.challenge"))
	} else {
		log.Print("Error: verify token failed.")
		fmt.Fprintf(w, "Error, wrong validation token")
	}
}

func webhookPostAction(w http.ResponseWriter, r *http.Request) {
	var receivedMessage ReceivedMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}

	if err = json.Unmarshal(body, &receivedMessage); err != nil {
		log.Print(err)
	}
	messagingEvents := receivedMessage.Entry[0].Messaging
	for i, event := range messagingEvents {
		senderID := event.Sender.ID
		log.Print(i)
		log.Print(event)
		if &event.Message != nil {
			if event.Message.Text == "QR" {
				q := []Quick_replies{
					{Content_type: "text", Title: "a", Payload: "a", Image_url: "https://user-images.githubusercontent.com/28649418/45468742-385b0500-b761-11e8-879e-2a5cef3b8ddc.png"},
					{Content_type: "location"},
					{Content_type: "user_phone_number"},
					{Content_type: "user_email"},
				}
				sendQuickReplies(senderID, "QuickReplies", q)
			} else if event.Message.Attachments != nil {
				if &event.Message.Attachments[0].Payload.Coordinates != nil {
					sendTextMessage(senderID, strconv.FormatFloat(event.Message.Attachments[0].Payload.Coordinates.Lat, 'f', 6, 64)+","+strconv.FormatFloat(event.Message.Attachments[0].Payload.Coordinates.Long, 'f', 6, 64))
				}
			} else if event.Message.Quick_reply.Payload != "" {
				switch event.Message.Quick_reply.Payload {
				case "a":
					sendTextMessage(senderID, "You selected a")

				case "b":
					sendTextMessage(senderID, "You selected b")

				case "c":
					sendTextMessage(senderID, "You selected c")
				default:
					sendTextMessage(senderID, "Payload: "+event.Message.Quick_reply.Payload)
				}
			} else {
				sendTextMessage(senderID, "yey")
			}
		}
	}
	fmt.Fprintf(w, "Success")
}

func sendQuickReplies(senderID string, text string, quick_replies []Quick_replies) {
	recipient := new(Recipient)
	recipient.ID = senderID
	m := new(SendMessage)
	m.Recipient = *recipient
	m.Message.Quick_replies = quick_replies
	m.Message.Text = text
	b, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
	}

	req, err := http.NewRequest("POST", EndPoint, bytes.NewBuffer(b))
	if err != nil {
		log.Print(err)
	}

	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
	}
	log.Print(result)
}

func sendTextMessage(senderID string, text string) {
	recipient := new(Recipient)
	recipient.ID = senderID
	m := new(SendMessage)
	m.Recipient = *recipient
	m.Message.Text = text
	b, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
	}

	req, err := http.NewRequest("POST", EndPoint, bytes.NewBuffer(b))
	if err != nil {
		log.Print(err)
	}

	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
	}
	log.Print(result)
}
