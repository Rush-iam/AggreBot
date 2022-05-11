package courier

const retryCountSendMessageThreshold = 60

func (c *courier) RunSender() {
	for {
		go c.processRoutine(<-c.jobQueue)
	}
}

func (c *courier) processRoutine(job job) {
	if !job.ok {
		retryCount := job.source.RetryCount + 1
		if retryCount%retryCountSendMessageThreshold == 0 {
			//send message to user
			c.sourcesInWork.Remove(job.source.Id)
		} else {
			// update retry_count in DB
		}
		return
	}

	// get all hashes from DB for that source
	// remove from DB hashes that are absent now
	// process new hashes
	// send to user

}
