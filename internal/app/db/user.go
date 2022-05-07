package db

import (
	"AggreBot/api"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddUser(id *api.UserId) error {
	return addUserQuery(id)
}

func addUserQuery(id *api.UserId) error {
	commandTag, err := db.conn.Exec(db.ctx,
		"INSERT INTO users VALUES ($1, $2) ON CONFLICT DO NOTHING", id.Id, true,
	)
	if err == nil && commandTag.RowsAffected() == 0 {
		err = status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("db.AddUser: <%+v> already exists", id),
		)
	}
	return err
}

func UpdateUser(user *api.User) error {
	return updateUserQuery(user)
}

func updateUserQuery(user *api.User) error {
	commandTag, err := db.conn.Exec(db.ctx,
		"UPDATE users SET active = $1 WHERE id = $2", user.Active, user.Id,
	)
	if err == nil && commandTag.RowsAffected() == 0 {
		err = status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateUser: <%+v> not found", user),
		)
	}
	return err
}

func DeleteUser(id *api.UserId) error {
	if isUserExists(id.Id) == false {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteUser: <%+v> not found", id),
		)
	}
	return deleteUserQuery(id)
}

func deleteUserQuery(id *api.UserId) error {
	delete(db.users, id.Id)
	return nil
}

func isUserExists(id int64) bool {
	_, found := db.users[id]
	return found
}
