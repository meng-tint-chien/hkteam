package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"math/rand"
 	"strconv"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/PuerkitoBio/goquery"
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
func GetJokes(){
  doc, err := goquery.NewDocument("https://imgur.com/a/6FdES")
  if err != nil{
    log.Fatal(err)
  }
  doc.Find(".content").Each(func(i int, s *goquery.Selection){
    fmt.Println(s.Text())
  })
}


func callbackHandler(w http.ResponseWriter, r *http.Request) {
	var guu = 0
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
				     GetJokes()
					 jj:= rand.New(rand.NewSource(time.Now().UnixNano()))
					guu=(jj.Intn(55))+44
					str2:=strconv.Itoa(guu)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://fapality.com/contents/albums/main/680x9999/1000/1193/509"+str2+".jpg","https://fapality.com/contents/albums/main/680x9999/1000/1193/509"+str2+".jpg")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
