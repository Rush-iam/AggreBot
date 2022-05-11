package db_client

import (
	"AggreBot/internal/pkg/api"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (db *Client) AddUser(id *api.UserId) error {
	rowsAffected, err := db.addUserQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("db.AddUser: <%+v> already exists", id),
		)
	}
	return err
}

func (db *Client) addUserQuery(id *api.UserId) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"INSERT INTO users (id) VALUES ($1) ON CONFLICT DO NOTHING", id.Id,
	)
	return cmdTag.RowsAffected(), err
}

func (db *Client) GetUser(id *api.UserId) (*api.User, error) {
	user, err := db.getUserQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.GetUser: <%+v> not found", id),
		)
	}
	return user, err
}

func (db *Client) getUserQuery(id *api.UserId) (*api.User, error) {
	var user api.User
	err := db.conn.QueryRow(db.ctx,
		"SELECT * FROM users WHERE id = $1", id.Id,
	).Scan(&user.Id, &user.Filter)
	if err != nil {
		return nil, err
	}
	return &user, err
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
	return cmdTag.RowsAffected(), err
}

func (db *Client) DeleteUser(id *api.UserId) error {
	rowsAffected, err := db.deleteUserQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteUser: <%+v> not found", id),
		)
	}
	return err
}

func (db *Client) deleteUserQuery(id *api.UserId) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM users WHERE id = $1", id.Id,
	)
	return cmdTag.RowsAffected(), err
}
