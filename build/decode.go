package main

import "time"

type echoEvt struct {
	Session struct {
		SessionID   string `json:"sessionId"`
		Application struct {
			ApplicationID string `json:"applicationId"`
		} `json:"application"`
		Attributes struct {
		} `json:"attributes"`
		User struct {
			UserID      string `json:"userId"`
			AccessToken string `json:"accessToken"`
		} `json:"user"`
		New bool `json:"new"`
	} `json:"session"`
	Request struct {
		Type      string    `json:"type"`
		RequestID string    `json:"requestId"`
		Locale    string    `json:"locale"`
		Timestamp time.Time `json:"timestamp"`
		Intent    struct {
			Name  string `json:"name"`
			Slots struct {
				Type struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"type"`
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
	Version string `json:"version"`
}

type alexaResponse struct {
	Version  string `json:"version,omitempty"`
	Response struct {
		Outputspeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
		Reprompt         interface{} `json:"reprompt,omitempty"`
		Shouldendsession bool        `json:"shouldEndSession"`
	} `json:"response"`
	Sessionattributes struct {
	} `json:"sessionAttributes,omitempty"`
}

type messages struct {
	ID    string `json:"id"`
	Inbox struct {
		Data []struct {
			ID string `json:"id"`
			To struct {
				Data []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"data"`
			} `json:"to"`
			Unread      int    `json:"unread"`
			Unseen      int    `json:"unseen"`
			UpdatedTime string `json:"updated_time"`
		} `json:"data"`
		Paging struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"paging"`
		Summary struct {
			UnreadCount int    `json:"unread_count"`
			UnseenCount int    `json:"unseen_count"`
			UpdatedTime string `json:"updated_time"`
		} `json:"summary"`
	} `json:"inbox"`
}
