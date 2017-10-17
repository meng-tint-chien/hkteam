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
	var randint
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
				if message.Text=="fuck"{
					randint=(rand.Intn(55))+44
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://fapality.com/contents/albums/main/680x9999/1000/1193/509"+randint+".jpg","https://fapality.com/contents/albums/main/680x9999/1000/1193/509"+randint+".jpg")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
