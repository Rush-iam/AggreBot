package courier

import (
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/tg_client"
	"log"
)

const retryCountSendMessageThreshold = 60

func (c *courier) readErrorHandler(source *db_client.CourierSource) {
	newRetryCount := source.RetryCount + 1
	if newRetryCount%retryCountSendMessageThreshold == 0 {
		c.sendReadError(source)
	} else {
		c.updateSourceRetryCount(source.Id, newRetryCount)
	}
}

func (c *courier) sendReadError(source *db_client.CourierSource) {
	message := readErrorMessage(source.Name, source.RetryCount)
	err := tg_client.SendMessage(c.tg, source.UserId, message)
	if err != nil {
		log.Printf("c.readErrorHandler: %v", err)
	} else {
		log.Printf("[bot] %s", message)
	}
}
