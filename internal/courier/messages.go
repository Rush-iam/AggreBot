package courier

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"strings"
)

func feedItemMessage(item *gofeed.Item, sourceName string) string {
	message := []string{
		fmt.Sprintf("%s - %s", item.PublishedParsed.Format("2.01.06"), sourceName),
		item.Title,
		item.Link,
	}
	return strings.Join(message, "\n\n")
}

func readErrorMessage(sourceName string, sourceRetryCount int32) string {
	return fmt.Sprintf(
		"😵‍💫 I can't read Source \"%s\" - tried %d times",
		sourceName, sourceRetryCount,
	)
}
