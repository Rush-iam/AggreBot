package rss_feed

import (
	"context"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
)

func Fetch(url string) (feed *gofeed.Feed, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	feedParser := gofeed.NewParser()
	return feedParser.ParseURLWithContext(url, ctx)
}

func GetTitle(url string) (string, bool) {
	feed, err := Fetch(url)
	if err != nil {
		log.Printf("rss_feed.getTitle: %v", err)
		return "", false
	}
	return feed.Title, true
}
