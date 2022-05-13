package courier

import (
	"AggreBot/internal/pkg/set_utils"
	"AggreBot/internal/pkg/tg_client"
	"github.com/mmcdole/gofeed"
	"log"
	"regexp"
	"sort"
)

type feedItemToSend struct {
	item *gofeed.Item
	hash string
}

func (c *courier) RunSender() {
	for {
		go c.senderRoutine(<-c.jobQueue)
	}
}

func (c *courier) senderRoutine(job job) {
	defer c.sourcesInWork.Remove(job.source.Id)

	if job.wasReadError {
		c.readErrorHandler(job.source)
		return
	}
	previousEntries := makeHashMapEntries(c.getSourceEntries(job.source.Id))
	currentEntries := makeHashMapFeedItems(job.entries)

	var deprecatedEntries []int64
	for _, hash := range set_utils.Difference(previousEntries, currentEntries) {
		deprecatedEntries = append(deprecatedEntries, previousEntries[hash])
	}
	c.deleteEntries(deprecatedEntries)

	var newFeedItems []feedItemToSend
	for hash, item := range set_utils.DifferenceMap(currentEntries, previousEntries) {
		newFeedItems = append(newFeedItems, feedItemToSend{item, hash})
	}
	sortFeedItems(newFeedItems)
	newEntries := c.sendFeedItems(newFeedItems, job)
	c.addEntries(job.source.Id, newEntries)
}

func sortFeedItems(feedItems []feedItemToSend) {
	sort.Slice(
		feedItems,
		func(i, k int) bool {
			return feedItems[i].item.PublishedParsed.Before(
				*feedItems[k].item.PublishedParsed)
		},
	)
}

func (c *courier) sendFeedItems(feedItems []feedItemToSend, job job) []string {
	newEntries := make([]string, 0, len(feedItems))
	for _, feedItem := range feedItems {
		matched, _ := regexp.MatchString(job.source.Filter, feedItem.item.Title)
		if matched {
			message := feedItemMessage(feedItem.item, job.source.Name)
			err := tg_client.SendMessage(c.tg, job.source.UserId, message)
			if err != nil {
				log.Printf("c.senderRoutine: %v", err)
				break
			}
		}
		newEntries = append(newEntries, feedItem.hash)
	}
	return newEntries
}
