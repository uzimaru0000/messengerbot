# messengerbot
## Usage   
### クイック返信 #2
クイック返信では、タイトルと画像(任意)の付いた最大11個のボタン一式をスレッド内で使用し、作成画面にわかりやすく表示できます。また、クイック返信を使用すると、利用者の現在地、メールアドレス、電話番号をリクエストできます。

`sendQuickReplies(senderID, "text", quick_replies)`


引数のquick_repliesに下記の構造体リストを参照することで最大１１個のボタンを指定できます。
```
type Quick_replies struct {
	Content_type string `json:"content_type"`
	Title        string `json:"title"`
	Payload      string `json:"payload"`
	Image_url    string `json:"image_url"`
}
```

<img src="https://user-images.githubusercontent.com/28649418/45469371-275fc300-b764-11e8-871e-46fbacad47df.jpg" width="540" height="200" />
https://developers.facebook.com/docs/messenger-platform/send-messages/quick-replies
