package db_client

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func stringCut(str string, maxLength int) string {
	strR := []rune(str)
	if len(strR) > maxLength {
		return string(strR[:maxLength])
	}
	return str
}

func (db *Client) AddSource(userId int64, name, url string) (*api.Source, error) {
	return db.addSourceQuery(userId, name, url)
}

func (db *Client) addSourceQuery(userId int64, name, url string) (*api.Source, error) {
	name = stringCut(name, 256)
	var id int64
	var isActive bool
	var retryCount int32
	err := db.conn.QueryRow(db.ctx,
		"INSERT INTO sources (user_id, name, url)"+
			"VALUES ($1, $2, $3) RETURNING id, is_active, retry_count",
		userId, name, url,
	).Scan(&id, &isActive, &retryCount)
	if err != nil {
		return nil, err
	}
	return &api.Source{
		Id:         id,
		UserId:     userId,
		Name:       name,
		Url:        url,
		IsActive:   isActive,
		RetryCount: retryCount,
	}, nil
}

func (db *Client) GetSource(id int64) (*api.Source, error) {
	source, err := db.getSourceQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.GetSource: <%+v> not found", id),
		)
	}
	return source, err
}

func (db *Client) getSourceQuery(id int64) (*api.Source, error) {
	var source api.Source
	err := db.conn.QueryRow(db.ctx,
		"SELECT * FROM sources WHERE id = $1", id,
	).Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
		&source.IsActive, &source.RetryCount)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (db *Client) GetUserSources(userId int64) ([]*api.Source, error) {
	return db.getUserSourcesQuery(userId)
}

func (db *Client) getUserSourcesQuery(userId int64) ([]*api.Source, error) {
	var sources []*api.Source
	rows, err := db.conn.Query(db.ctx,
		"SELECT * FROM sources WHERE user_id = $1 ORDER BY id", userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var source api.Source
		err = rows.Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
			&source.IsActive, &source.RetryCount)
		if err != nil {
			return nil, err
		}
		sources = append(sources, &source)
	}
	return sources, nil
}

func (db *Client) UpdateSourceName(id int64, name string) error {
	rowsAffected, err := db.updateSourceNameQuery(id, name)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceName: <%+v> not found", id),
		)
	}
	return err
}

func (db *Client) updateSourceNameQuery(id int64, name string) (int64, error) {
	name = stringCut(name, 256)
	cmdTag, err := db.conn.Exec(db.ctx,
		"UPDATE sources SET name = $1 WHERE id = $2",
		name, id,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}

func (db *Client) UpdateSourceIsActive(id int64, isActive bool) (*api.Source, error) {
	source, err := db.updateSourceIsActiveQuery(id, isActive)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceIsActive: <%+v> not found", id),
		)
	}
	return source, err
}

func (db *Client) updateSourceIsActiveQuery(id int64, isActive bool) (*api.Source, error) {
	var source api.Source
	err := db.conn.QueryRow(db.ctx,
		"UPDATE sources SET is_active = $1 WHERE id = $2 RETURNING *",
		isActive, id,
	).Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
		&source.IsActive, &source.RetryCount)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (db *Client) DeleteSource(id int64) error {
	rowsAffected, err := db.deleteSourceQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteSource: <%+v> not found", id),
		)
	}
	return err
}

func (db *Client) deleteSourceQuery(id int64) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM sources WHERE id = $1", id,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}
