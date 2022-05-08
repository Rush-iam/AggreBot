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
	err := db.conn.QueryRow(db.ctx,
		"INSERT INTO sources (user_id, name, type, ref_str, ref_int)"+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		req.UserId, req.Name, req.Type, req.RefStr, req.RefInt,
	).Scan(&id.Id)
	return &id, err
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
	).Scan(&source.Id, &source.UserId, &source.Name, &source.Type,
		&source.RefStr, &source.RefInt, &source.LastChecked, &source.RetryCount)
	if err != nil {
		return nil, err
	}
	return &source, err
}

func GetUserSources(userId *api.UserId) (*api.Sources, error) {
	sources, err := getUserSourcesQuery(userId)
	return sources, err
}

func getUserSourcesQuery(userId *api.UserId) (*api.Sources, error) {
	var sources api.Sources
	rows, err := db.conn.Query(db.ctx,
		"SELECT * from sources WHERE user_id = $1", userId.Id,
	)
	defer rows.Close()
	for rows.Next() {
		var source api.Source
		err = rows.Scan(&source.Id, &source.UserId, &source.Name, &source.Type,
			&source.RefStr, &source.RefInt, &source.LastChecked,
			&source.RetryCount)
		if err != nil {
			return nil, err
		}
		sources.Sources = append(sources.Sources, &source)
	}
	return &sources, nil
}

func UpdateSource(Source *api.UpdateSourceRequest) error {
	rowsAffected, err := updateSourceQuery(Source)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateSource: <%+v> not found", Source),
		)
	}
	return err
}

func updateSourceQuery(source *api.UpdateSourceRequest) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"UPDATE sources SET name = $1 WHERE id = $2",
		source.Name, source.Id,
	)
	return cmdTag.RowsAffected(), err
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
