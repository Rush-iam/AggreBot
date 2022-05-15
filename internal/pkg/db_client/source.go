package db_client

import (
	"AggreBot/internal/pkg/api"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (db *Client) AddSource(req *api.AddSourceRequest) (*api.SourceId, error) {
	return db.addSourceQuery(req)
}

func (db *Client) addSourceQuery(req *api.AddSourceRequest) (*api.SourceId, error) {
	var id api.SourceId
	name := []rune(req.Name)
	if len(name) > 256 {
		name = name[:256]
	}
	err := db.conn.QueryRow(db.ctx,
		"INSERT INTO sources (user_id, name, url)"+
			"VALUES ($1, $2, $3) RETURNING id",
		req.UserId, string(name), req.Url,
	).Scan(&id.Id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (db *Client) GetSource(id *api.SourceId) (*api.Source, error) {
	source, err := db.getSourceQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.GetSource: <%+v> not found", id),
		)
	}
	return source, err
}

func (db *Client) getSourceQuery(id *api.SourceId) (*api.Source, error) {
	var source api.Source
	err := db.conn.QueryRow(db.ctx,
		"SELECT * FROM sources WHERE id = $1", id.Id,
	).Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
		&source.IsActive, &source.RetryCount)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (db *Client) GetUserSources(userId *api.UserId) (*api.Sources, error) {
	return db.getUserSourcesQuery(userId)
}

func (db *Client) getUserSourcesQuery(userId *api.UserId) (*api.Sources, error) {
	var sources api.Sources
	rows, err := db.conn.Query(db.ctx,
		"SELECT * FROM sources WHERE user_id = $1 ORDER BY id", userId.Id,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var source api.Source
		err = rows.Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
			&source.IsActive, &source.RetryCount)
		if err != nil {
			return nil, err
		}
		sources.Sources = append(sources.Sources, &source)
	}
	return &sources, nil
}

func (db *Client) UpdateSourceName(req *api.UpdateSourceNameRequest) error {
	rowsAffected, err := db.updateSourceNameQuery(req)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceName: <%+v> not found", req),
		)
	}
	return err
}

func (db *Client) updateSourceNameQuery(req *api.UpdateSourceNameRequest) (int64, error) {
	name := []rune(req.Name)
	if len(name) > 256 {
		name = name[:256]
	}
	cmdTag, err := db.conn.Exec(db.ctx,
		"UPDATE sources SET name = $1 WHERE id = $2",
		string(name), req.Id,
	)
	return cmdTag.RowsAffected(), err
}

func (db *Client) UpdateSourceIsActive(req *api.UpdateSourceIsActiveRequest) (*api.UpdateSourceIsActiveResponse, error) {
	resp, err := db.updateSourceIsActiveQuery(req)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceIsActive: <%+v> not found", req),
		)
	}
	return resp, err
}

func (db *Client) updateSourceIsActiveQuery(req *api.UpdateSourceIsActiveRequest) (*api.UpdateSourceIsActiveResponse, error) {
	var sourceInfo api.UpdateSourceIsActiveResponse
	err := db.conn.QueryRow(db.ctx,
		"UPDATE sources SET is_active = $1 WHERE id = $2 "+
			"RETURNING name, url, is_active",
		req.IsActive, req.Id,
	).Scan(&sourceInfo.Name, &sourceInfo.Url, &sourceInfo.IsActive)
	if err != nil {
		return nil, err
	}
	return &sourceInfo, nil
}

func (db *Client) DeleteSource(id *api.SourceId) error {
	rowsAffected, err := db.deleteSourceQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteSource: <%+v> not found", id),
		)
	}
	return err
}

func (db *Client) deleteSourceQuery(id *api.SourceId) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM sources WHERE id = $1", id.Id,
	)
	return cmdTag.RowsAffected(), err
}
