package db

import (
	"AggreBot/api"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddSource(req *api.AddSourceRequest) (*api.SourceId, error) {
	return addSourceQuery(req)
}

func addSourceQuery(req *api.AddSourceRequest) (*api.SourceId, error) {
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

func GetSource(id *api.SourceId) (*api.Source, error) {
	source, err := getSourceQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.GetSource: <%+v> not found", id),
		)
	}
	return source, err
}

func getSourceQuery(id *api.SourceId) (*api.Source, error) {
	var source api.Source
	err := db.conn.QueryRow(db.ctx,
		"SELECT * from sources WHERE id = $1", id.Id,
	).Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
		&source.IsActive, &source.LastChecked, &source.RetryCount)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func GetUserSources(userId *api.UserId) (*api.Sources, error) {
	sources, err := getUserSourcesQuery(userId)
	return sources, err
}

func getUserSourcesQuery(userId *api.UserId) (*api.Sources, error) {
	var sources api.Sources
	rows, err := db.conn.Query(db.ctx,
		"SELECT * from sources WHERE user_id = $1 ORDER BY id", userId.Id,
	)
	defer rows.Close()
	for rows.Next() {
		var source api.Source
		err = rows.Scan(&source.Id, &source.UserId, &source.Name, &source.Url,
			&source.IsActive, &source.LastChecked, &source.RetryCount)
		if err != nil {
			return nil, err
		}
		sources.Sources = append(sources.Sources, &source)
	}
	return &sources, nil
}

func UpdateSourceName(req *api.UpdateSourceNameRequest) error {
	rowsAffected, err := updateSourceNameQuery(req)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceName: <%+v> not found", req),
		)
	}
	return err
}

func updateSourceNameQuery(req *api.UpdateSourceNameRequest) (int64, error) {
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

func UpdateSourceToggleActive(id *api.SourceId) (*api.SourceToggleActiveResponse, error) {
	source, err := updateSourceToggleActiveQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSourceToggleActive: <%+v> not found", id),
		)
	}
	return source, nil
}

func updateSourceToggleActiveQuery(id *api.SourceId) (*api.SourceToggleActiveResponse, error) {
	var sourceInfo api.SourceToggleActiveResponse
	err := db.conn.QueryRow(db.ctx,
		"UPDATE sources SET is_active = NOT is_active WHERE id = $1 "+
			"RETURNING name, url, is_active",
		id.Id,
	).Scan(&sourceInfo.Name, &sourceInfo.Url, &sourceInfo.IsActive)
	if err != nil {
		return nil, err
	}
	return &sourceInfo, nil
}

func DeleteSource(id *api.SourceId) error {
	rowsAffected, err := deleteSourceQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteSource: <%+v> not found", id),
		)
	}
	return err
}

func deleteSourceQuery(id *api.SourceId) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM sources WHERE id = $1", id.Id,
	)
	return cmdTag.RowsAffected(), err
}
