package commands

import (
	"AggreBot/api"
	"AggreBot/internal/tg_frontend/grpc_client"
	"AggreBot/internal/tg_frontend/tg_client"
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"log"
	"net/url"
	"strconv"
)

type source struct {
	srcType string
	srcRef  string
	srcName string
}

func add(c Command) *string {
	var reply string

	if len(c.args) == 0 {
		reply = "Hey! You forgot source.\n" +
			"It can be either Telegram group or RSS feed url"
		return &reply
	}
	src, ok := parseSource(c.args[0])
	if !ok {
		reply = "Hmm... I can't figure out that source."
		return &reply
	}
	err := validateSource(src)
	if err != nil {
		reply = err.Error()
		return &reply
	}

	_, err = grpc_client.Cl.AddSource(
		context.Background(),
		&api.AddSourceRequest{
			UserId:    c.userId,
			Name:      src.srcName,
			Type:      src.srcType,
			Reference: src.srcRef,
		},
	)
	if err != nil {
		reply = "Oops. Internal Error. Please try again later."
	} else {
		var icon rune
		if src.srcType == "tg" {
			icon = '🗨'
		} else if src.srcType == "rss" {
			icon = '📰'
		}
		reply = fmt.Sprintf("Added:\n%v %s", icon, src.srcName)
	}
	return &reply
}

func parseSource(rawReference string) (*source, bool) {
	var src source

	u, err := url.Parse(rawReference)
	if err == nil && u.IsAbs() {
		if u.Host == "t.me" || u.Host == "www.t.me" {
			if len(u.Path) < 2 || u.Path[1] == '+' {
				return nil, false
			}
			src.srcType = "tg"
			src.srcRef = "@" + u.Path[1:]
		} else {
			src.srcType = "rss"
			src.srcRef = rawReference
		}
	} else {
		src.srcType = "tg"
		if rawReference[0] != '@' {
			src.srcRef = "@" + rawReference
		} else {
			src.srcRef = rawReference
		}
	}

	if src.srcRef != "" {
		return &src, true
	} else {
		return nil, false
	}
}

func validateSource(src *source) error {
	if src.srcType == "tg" {
		return validateTgSource(src)
	} else if src.srcType == "rss" {
		return validateRssSource(src)
	} else {
		panic("validateSource")
		return nil
	}
}

func validateTgSource(src *source) error {
	chatConfig := tgbotapi.ChatInfoConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             0,
			SuperGroupUsername: src.srcRef,
		},
	}
	chat, err := tg_client.Cl.GetChat(chatConfig)
	if err != nil {
		log.Printf("validateTgSource: %v", err)
		return errors.New("I had troubles trying to find that Telegram group, sorry")
	}
	src.srcName = chat.Title
	src.srcRef = strconv.FormatInt(chat.ID, 10)
	return nil
}

func validateRssSource(src *source) error {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(src.srcRef)
	if err != nil {
		log.Printf("validateRssSource: %v", err)
		return errors.New("I had troubles parsing RSS from that url, sorry")
	}
	src.srcName = feed.Title
	return nil
}
