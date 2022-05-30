package courier

import (
	"AggreBot/internal/pkg/db_client"
	"log"
)

func (c *courier) updateSourceRetryCount(sourceId int64, retryCount int32) {
	req := db_client.UpdateSourceRetryCountRequest{
		Id:         sourceId,
		RetryCount: retryCount,
	}
	err := c.db.UpdateSourceRetryCount(&req)
	if err != nil {
		log.Printf("c.updateSourceRetryCount: %v", err)
	}
}

func (c *courier) getSourceEntries(sourceId int64) []*db_client.Entry {
	entries, err := c.db.GetSourceEntries(sourceId)
	if err != nil {
		log.Printf("c.getSourceEntries: %v", err)
		return []*db_client.Entry{}
	}
	return entries
}

func (c *courier) addEntries(sourceId int64, hashes []string) {
	entries := make([]db_client.AddEntryRequest, len(hashes))
	for i, hash := range hashes {
		entries[i] = db_client.AddEntryRequest{
			SourceId: sourceId,
			Hash:     hash,
		}
	}
	_, err := c.db.AddEntries(entries)
	if err != nil {
		log.Printf("c.addEntries: %v", err)
	}
}

func (c *courier) deleteEntries(ids []int64) {
	_, err := c.db.DeleteEntries(ids)
	if err != nil {
		log.Printf("c.deleteEntries: %v", err)
	}
}
