package db

import (
	"AggreBot/api"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddUser(id *api.UserId) error {
	rowsAffected, err := addUserQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("db.AddUser: <%+v> already exists", id),
		)
	}
	return err
}

func addUserQuery(id *api.UserId) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"INSERT INTO users (id) VALUES ($1) ON CONFLICT DO NOTHING", id.Id,
	)
	return cmdTag.RowsAffected(), err
}

func GetUser(id *api.UserId) (*api.User, error) {
	user, err := getUserQuery(id)
	if err == pgx.ErrNoRows {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.GetUser: <%+v> not found", id),
		)
	}
	return user, err
}

func getUserQuery(id *api.UserId) (*api.User, error) {
	var user api.User
	err := db.conn.QueryRow(db.ctx,
		"SELECT * from users WHERE id = $1", id.Id,
	).Scan(&user.Id, &user.Filter)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func UpdateUser(user *api.User) error {
	rowsAffected, err := updateUserQuery(user)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateUser: <%+v> not found", user),
		)
	}
	return err
}

func updateUserQuery(user *api.User) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"UPDATE users SET filter = $1 WHERE id = $2",
		user.Filter, user.Id,
	)
	return cmdTag.RowsAffected(), err
}

func DeleteUser(id *api.UserId) error {
	rowsAffected, err := deleteUserQuery(id)
	if rowsAffected == 0 && err == nil {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteUser: <%+v> not found", id),
		)
	}
	return err
}

func deleteUserQuery(id *api.UserId) (int64, error) {
	cmdTag, err := db.conn.Exec(db.ctx,
		"DELETE FROM users WHERE id = $1", id.Id,
	)
	return cmdTag.RowsAffected(), err
}
