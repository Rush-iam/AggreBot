package courier

import (
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/tg_client"
	"log"
)

const readErrorNotifyThreshold = 60

func (c *courier) readErrorReset(source *db_client.CourierSource) {
	c.updateSourceRetryCount(source.Id, 0)
}

func (c *courier) readErrorHandler(source *db_client.CourierSource) {
	if source.RetryCount > 0 && source.RetryCount%readErrorNotifyThreshold == 0 {
		c.readErrorNotifyUser(source)
	}
	c.updateSourceRetryCount(source.Id, source.RetryCount+1)
}

func (c *courier) readErrorNotifyUser(source *db_client.CourierSource) {
	message := readErrorMessage(source.Name, source.RetryCount)
	err := tg_client.SendMessage(c.tg, source.UserId, message)
	if err != nil {
		log.Printf("c.readErrorHandler: %v", err)
	} else {
		log.Printf("[bot] %s", message)
	}
}
