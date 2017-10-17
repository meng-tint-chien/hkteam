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


client.replyImage(
  REPLY_TOKEN,
  'http://ddragon.leagueoflegends.com/cdn/img/champion/splash/Yasuo_1.jpg',
  'http://ddragon.leagueoflegends.com/cdn/img/champion/splash/Yasuo_1.jpg'
);