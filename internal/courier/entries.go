package courier

import (
	"AggreBot/internal/pkg/db_client"
	"crypto/md5"
	"fmt"
	"github.com/mmcdole/gofeed"
)

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
