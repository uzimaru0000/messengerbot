# messengerbot
## Usage   
### メッセージ #1
https://developers.facebook.com/docs/messenger-platform/send-messages

Messengerプラットフォームでは、テキスト、音声、画像、動画、ファイルなど、構造化されていないさまざまな種類のコンテンツを送信できます。

また、定義済みのさまざまなメッセージテンプレートを使用することによって、より工夫をこらした構造化されたメッセージを送ることができます。詳しくは、「テンプレート」をご覧ください。

すべてのAPI呼び出しとリクエストプロパティのリストについては、「送信APIリファレンス」をご覧ください。
### クイック返信 #2 (text, location, user_phone_number, user_email 対応済み)
https://developers.facebook.com/docs/messenger-platform/send-messages/quick-replies

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

#### text
Quick_repliesに下記のような構造体を追加します。
```
{
  "content_type":"text",
  "title":"<BUTTON_TEXT>",
  "image_url":"http://example.com/img/red.png",
  "payload":"<DEVELOPER_DEFINED_PAYLOAD>"
}
```

#### location   
位置情報のクイック返信では、ボットから利用者の位置情報をリクエストできます。位置情報のクイック返信をタップすると、利用者が現在地の設定に使用できる検索可能な地図が表示されます。
<img src="https://user-images.githubusercontent.com/28649418/45482681-ad423500-b789-11e8-87b8-511808ede6d5.jpg" width="540" height="200" />
<img src="https://user-images.githubusercontent.com/28649418/45483750-d617f980-b78c-11e8-9b1c-f620d6d15e4e.jpg" width="540" height="1000" />
<img src="https://user-images.githubusercontent.com/28649418/45482736-d2cf3e80-b789-11e8-98f4-966ebe63040b.jpg" width="540" height="400" />

Quick_repliesに下記のような構造体を追加します。

`{Content_type: "location"}`   

緯度と経度がwebhookイベントの**payload.coordinates**プロパティを介して利用者に送信されます。

#### user_phone_number
**この機能には、iOS用Messenger v144、またはAndroid用Messenger v142が必要です。**

Quick_repliesに下記のような構造体を追加します。
```
{
  "content_type":"user_phone_number"
}
```
利用者がクイック返信をタップすると、携帯電話番号がmessages Webhookイベントのpayload属性に渡されます。
#### user_email
**この機能には、iOS用Messenger v144、またはAndroid用Messenger v142が必要です。**

Quick_repliesに下記のような構造体を追加します。
```
{
  "content_type":"user_email"
}
```
利用者がクイック返信をタップすると、メールアドレスがmessages Webhookイベントのpayload属性に渡されます。

### メッセージテンプレート #3
https://developers.facebook.com/docs/messenger-platform/send-messages/templates

メッセージテンプレートを使うと、1つのメッセージを送信する際にボタン、画像、リストなどを統合することによって、通常のテキストメッセージよりも優れたスレッド内エクスペリエンスを提供できます。テンプレートは、製品情報を表示したり、事前に用意されたオプションから選択するよう促したり、検索結果を表示したりとさまざまな用途に使用できます。  

#### ButtonTemplate
ボタンテンプレートでは、最大3つのボタンがあるテキストメッセージを送信します。このテンプレートは、事前定義された質問への回答やアクションといった選択肢を受信者に提供する場合に便利です。  

`template.NewButtonTemplate(title string, buttons []models.Button)`  
新しいButtonTemplateを作成します

#### ListTemplate
リストテンプレートは、2～4つの構造化アイテムと、任意のグローバルボタンが下部に配置されたリストです。各アイテムにはサムネイル画像、タイトル、サブタイトル、1つのボタンを含めることができます。  
`template.NewListTemplate()`
新しいListTemplateを作成します

#### GenericTemplate
一般テンプレートは、タイトル、サブタイトル、画像、最大3つのボタンが含まれる簡単な構造化メッセージです。  
`template.NewGenericTemplate()`  
新しいGenericTemplateを作成します

#### MediaTemplate
メディアテンプレートでは、画像、GIF、動画を構造化メッセージとして、任意のボタンとともに送信できます。メディアテンプレートで送信された動画とアニメーションGIFは、スレッド内で再生できます。  
`template.NewMediaTemplate()`  
新しいMediatemplateを作成します
  
`template.NewTemplate(senderID string, template *models.Template) *models.SendMessage`  
Templateインターフェースを実装した構造体からテンプレートのSendMessageを生成します

### ボタン #4
https://developers.facebook.com/docs/messenger-platform/send-messages/buttons

ほとんどのメッセージテンプレートや固定メニューでは、さまざまなアクションを実行できるボタンがサポートされています。これらのボタンを使用すると、Messengerウェブビューで開く、支払いフローを開始する、Webhookにポストバックメッセージを送るなど、テンプレートへの応答方法をメッセージの受信者に提示できます。

メッセージテンプレートの場合、含まれるボタンはbuttons配列のオブジェクトで定義されます。固定メニューの場合、ボタンはcall_to_actions配列のオブジェクトで定義されます。各ボタンタイプの特定の目的とフォーマットについて詳しくは、以下をご覧ください。

ボタンの種類は、`URL`, `PostBack`, `Call`, `GamePlay`, `LogIn`, `LogOut`, `Share`の7種類あります。  
それぞれ`New[BUTTON_TYPE]Button(params)`で生成することができます。

### 固定メニュー #5　（未実装）
https://developers.facebook.com/docs/messenger-platform/send-messages/persistent-menu

固定メニューを使用すると、Messengerのスレッド内にユーザーインターフェイスを常に表示しておくことができます。固定メニューを表示しておけば、Messengerボットの主要な機能を利用者がいつでも簡単に利用できます。

# TMP
<img src="" width="540" height="200" />


