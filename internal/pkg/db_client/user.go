package db_client

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (db *Client) AddUser(user *api.User) error {
	rowsAffected, err := db.addUserQuery(user)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("db.AddUser: <%+v> already exists", user),
		)
	}
	return err
}

func (db *Client) addUserQuery(user *api.User) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"INSERT INTO users (id, filter) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		user.Id, user.Filter,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}

func (db *Client) GetUser(id int64) (*api.User, error) {
	user, err := db.getUserQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.GetUser: <%+v> not found", id),
		)
	}
	return user, err
}

func (db *Client) getUserQuery(id int64) (*api.User, error) {
	var user api.User
	err := db.conn.QueryRow(db.ctx,
		"SELECT * FROM users WHERE id = $1", id,
	).Scan(&user.Id, &user.Filter)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *Client) UpdateUserFilter(user *api.User) error {
	rowsAffected, err := db.updateUserFilterQuery(user)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateUser: <%+v> not found", user),
		)
	}
	return err
}

func (db *Client) updateUserFilterQuery(user *api.User) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"UPDATE users SET filter = $1 WHERE id = $2",
		user.Filter, user.Id,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}

func (db *Client) DeleteUser(id int64) error {
	rowsAffected, err := db.deleteUserQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteUser: <%+v> not found", id),
		)
	}
	return err
}

func (db *Client) deleteUserQuery(id int64) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM users WHERE id = $1", id,
	)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), err
}
