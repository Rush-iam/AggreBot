package db_client

type CourierSource struct {
	Id         int64
	UserId     int64
	Name       string
	Url        string
	IsActive   bool
	RetryCount int64
	Filter     string
}

func (db *Client) GetActiveSources() ([]*CourierSource, error) {
	return db.getActiveSourcesQuery()
}

func (db *Client) getActiveSourcesQuery() ([]*CourierSource, error) {
	rows, err := db.conn.Query(db.ctx,
		"SELECT sources.*, users.filter FROM sources "+
			"JOIN users ON sources.user_id = users.id "+
			"WHERE is_active = true",
	)
	defer rows.Close()
	var sources []*CourierSource
	for rows.Next() {
		var source CourierSource
		err = rows.Scan(
			&source.Id,
			&source.UserId,
			&source.Name,
			&source.Url,
			&source.IsActive,
			&source.RetryCount,
			&source.Filter,
		)
		if err != nil {
			return nil, err
		}
		sources = append(sources, &source)
	}
	return sources, nil
}
