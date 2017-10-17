package main

import (
	
	"log"
	"net/http"
	"os"
	
	

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	
	if err != nil {
		log.Fatal(err)
	}

	// serve /static/** files
	staticFileServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/static/", http.StripPrefix("/static/", staticFileServer).ServeHTTP)
	// serve /downloaded/** files
	downloadedFileServer := http.FileServer(http.Dir(app.downloadDir))
	http.HandleFunc("/downloaded/", http.StripPrefix("/downloaded/", downloadedFileServer).ServeHTTP)

	http.HandleFunc("/callback", app.Callback)
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
type KitchenSink struct {
	bot         *linebot.Client
	appBaseURL  string
	downloadDir string
}
func (app *KitchenSink) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	
	case "buttons":
		imageURL := "http://ddragon.leagueoflegends.com/cdn/img/champion/splash/Yasuo_1.jpg"
		template := linebot.NewButtonsTemplate(
			imageURL, "My button sample", "Hello, my button",
			linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
			linebot.NewPostbackTemplateAction("Say hello1", "hello ", ""),
			linebot.NewPostbackTemplateAction("hello2", "hello ", "hello "),
			linebot.NewMessageTemplateAction("Say message", "Rice="),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Buttons alt text", template),
		).Do(); err != nil {
			return err
		}

	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(message.Text),
		).Do(); err != nil {
			return err
		}
	}
return nil
}

