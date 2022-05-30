package db_client

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourierSource struct {
	Id         int64
	UserId     int64
	Name       string
	Url        string
	IsActive   bool
	RetryCount int32
	Filter     string
}

type UpdateSourceRetryCountRequest struct {
	Id         int64
	RetryCount int32
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
	if err != nil {
		return nil, err
	}
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

func (db *Client) UpdateSourceRetryCount(req *UpdateSourceRetryCountRequest) error {
	rowsAffected, err := db.updateSourceRetryCountQuery(req)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceRetryCount: <%+v> not found", req),
		)
	}
	return err
}

func (db *Client) updateSourceRetryCountQuery(req *UpdateSourceRetryCountRequest) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"UPDATE sources SET retry_count = $1 WHERE id = $2",
		req.RetryCount, req.Id,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}
