package courier

import (
	"crypto/md5"
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/db_client"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/set_utils"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/tg_client"
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
	defer c.sourcesInWork.Delete(job.source.Id)

	if job.wasReadError {
		c.readErrorHandler(job.source)
		return
	} else if job.source.RetryCount > 0 {
		c.readErrorReset(job.source)
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
			err := tg_client.SendMessage(c.tg, job.source.UserId, message, nil)
			if err != nil {
				log.Printf("c.senderRoutine: %v", err)
				break
			}
		}
		newEntries = append(newEntries, feedItem.hash)
	}
	return newEntries
}

func makeHashMapEntries(entries []*db_client.Entry) map[string]int64 {
	hashMap := make(map[string]int64, len(entries))
	for _, entry := range entries {
		hashMap[entry.Hash] = entry.Id
	}
	return hashMap
}

func makeHashMapFeedItems(feedItems []*gofeed.Item) map[string]*gofeed.Item {
	hashMap := make(map[string]*gofeed.Item, len(feedItems))
	for _, item := range feedItems {
		hash := makeEntryHash(item.Title, item.Link, item.GUID)
		hashMap[hash] = item
	}
	return hashMap
}

func makeEntryHash(title, link, GUID string) string {
	inputString := fmt.Sprintf("%s%s%s", title, link, GUID)
	return fmt.Sprintf("%x", md5.Sum([]byte(inputString)))
}
