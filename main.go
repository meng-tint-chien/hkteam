package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
		
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				message.Text="1"
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://yt3.ggpht.com/-9NvD-6i9toM/AAAAAAAAAAI/AAAAAAAAAAA/GE9FIajPWO8/s48-c-k-no-mo-rj-c0xffffff/photo.jpg","https://yt3.ggpht.com/-9NvD-6i9toM/AAAAAAAAAAI/AAAAAAAAAAA/GE9FIajPWO8/s48-c-k-no-mo-rj-c0xffffff/photo.jpg")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}