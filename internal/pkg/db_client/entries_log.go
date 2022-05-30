package db_client

import (
	"github.com/jackc/pgx/v4"
)

type Entry struct {
	Id       int64
	SourceId int64
	Hash     string
}

type AddEntryRequest struct {
	SourceId int64
	Hash     string
}

func (db *Client) GetSourceEntries(id int64) ([]*Entry, error) {
	return db.getSourceEntriesQuery(id)
}

func (db *Client) getSourceEntriesQuery(id int64) ([]*Entry, error) {
	rows, err := db.conn.Query(db.ctx,
		"SELECT * FROM entries_log WHERE source_id = $1", id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entries []*Entry
	for rows.Next() {
		var entry Entry
		err = rows.Scan(
			&entry.Id,
			&entry.SourceId,
			&entry.Hash,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &entry)
	}
	return entries, nil
}

func (db *Client) AddEntries(entries []AddEntryRequest) (int64, error) {
	return db.addEntriesQuery(entries)
}

func (db *Client) addEntriesQuery(entries []AddEntryRequest) (int64, error) {
	table := "entries_log"
	columns := []string{"source_id", "hash"}
	rowsAffected, err := db.conn.CopyFrom(
		db.ctx, pgx.Identifier{table}, columns,
		pgx.CopyFromSlice(
			len(entries),
			func(i int) ([]interface{}, error) {
				return []interface{}{
					entries[i].SourceId, entries[i].Hash,
				}, nil
			},
		),
	)
	return rowsAffected, err
}

func (db *Client) DeleteEntries(ids []int64) (int64, error) {
	return db.deleteEntriesQuery(ids)
}

func (db *Client) deleteEntriesQuery(ids []int64) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM entries_log WHERE id = ANY($1)", ids,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}
