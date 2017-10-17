package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"math/rand"
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
				if message.Text=="fuck"{
					 jj:= rand.New(rand.NewSource(time.Now().UnixNano()))
					
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://fapality.com/contents/albums/main/680x9999/1000/1193/509"+((jj.Intn(55))+44)+".jpg","https://fapality.com/contents/albums/main/680x9999/1000/1193/509"+((jj.Intn(55))+44)+".jpg")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
