package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	fb "github.com/huandu/facebook"
)

var (
	// AppID is Facebook App ID
	AppID string
	// AppSecret is Facebook App Sexret
	AppSecret string
	// AlexaAppID is Echo App ID from Amazon Dashboard
	AlexaAppID string
)

func Handle(evt *echoEvt, ctx *runtime.Context) (interface{}, error) {
	res := myFacebookCall(evt)
	return res, nil
}

func main() {}

func init() {
	AppID = os.Getenv("APP_ID")
	AppSecret = os.Getenv("APP_SECRET")
	AlexaAppID = os.Getenv("ALEXA_APP_ID")
}

func myFacebookCall(echoReq *echoEvt) *alexaResponse {
	s := new(echoReq.Session.User.AccessToken)
	count, unreadMsgFrom := unreadMsg(s)
	var speechText string
	if count > 0 {
		speechText = "You have " + strconv.Itoa(count) + " of unread messages. From " + strings.Join(unreadMsgFrom, "... ")
	} else {
		speechText = "You have 0 unread message."
	}
	return generateAlexaResponse(speechText)
}

func generateAlexaResponse(text string) *alexaResponse {
	res := &alexaResponse{}
	res.Version = "1.0"
	res.Response.Outputspeech.Type = "PlainText"
	res.Response.Shouldendsession = true
	res.Response.Outputspeech.Text = text
	return res
}

func new(token string) *fb.Session {
	app := fb.New(AppID, AppSecret)
	s := app.Session(token)
	s.Version = "v2.3"

	return s
}

func unreadMsg(s *fb.Session) (int, []string) {
	res, err := s.Get("/me", fb.Params{
		"fields": "inbox{from,message,subject,updated_time,to,unread,unseen,id,comments}",
	})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}

	var i inbox
	res.Decode(&i)
	var unReadMsgCount []string
	for _, msg := range i.Inbox.Data {
		if msg.Unread > 0 {
			var comments []string
			for i := len(msg.Comments.Data); i > (len(msg.Comments.Data) - msg.Unread); i-- {
				comments = append(comments, msg.Comments.Data[i-1].Message)
			}
			from := msg.To.Data[0].Name + "..." + "message" + "..." + strings.Join(comments, "...")
			unReadMsgCount = append(unReadMsgCount, from)
		}
	}
	return i.Inbox.Summary.UnreadCount, unReadMsgCount
}
