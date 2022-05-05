package db

import (
	"AggreBot/api"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddUser(id *api.UserId) error {
	if isUserExists(id.Id) {
		return status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("db.AddUser: <%+v> already exists", id),
		)
	}
	return addUserQuery(id)
}

func addUserQuery(id *api.UserId) error {
	db.users[id.Id] = api.User{
		Id:     id.Id,
		Active: true,
	}
	return nil
}

func UpdateUser(user *api.User) error {
	if isUserExists(user.Id) == false {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateUser: <%+v> not found", user),
		)
	}
	return updateUserQuery(user)
}

func updateUserQuery(user *api.User) error {
	db.users[user.Id] = api.User{
		Id:     user.Id,
		Active: user.Active,
	}
	return nil
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
