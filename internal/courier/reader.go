package courier

import (
	"AggreBot/internal/pkg/db_client"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
)

const readPeriod = time.Minute * 1

func (c *courier) RunReader() {
	for {
		repeatTime := time.Now().Add(readPeriod)
		sources, err := c.db.GetActiveSources()
		if err != nil {
			log.Printf("Reader: %v", err)
		} else {
			for _, source := range sources {
				if c.sourcesInWork.Has(source.Id) == false {
					c.sourcesInWork.Add(source.Id)
					go c.readerRoutine(source)
				}
			}
		}
		time.Sleep(time.Until(repeatTime))
	}
}

func (c *courier) readerRoutine(source *db_client.CourierSource) {
	newJob := job{source: source}

	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(source.Url)
	if err != nil {
		newJob.ok = false
		log.Printf("readerRoutine: %+v, %v", source, err)
	} else {
		newJob.ok = true
		newJob.entries = feed.Items
	}
	c.jobQueue <- newJob
}
