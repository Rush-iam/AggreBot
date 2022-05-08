package handlers

import (
	"AggreBot/api"
	"AggreBot/internal/app/db"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (s Server) AddUser(_ context.Context, id *api.UserId) (*api.UserId, error) {
	err := db.AddUser(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.AddUser: <%+v>", id)
	}
	return id, err
}

func (s Server) GetUser(_ context.Context, id *api.UserId) (*api.User, error) {
	user, err := db.GetUser(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetUser: <%+v>", id)
	}
	return user, err
}

func (s Server) UpdateUserFilter(_ context.Context, user *api.User) (*empty.Empty, error) {
	err := db.UpdateUserFilter(user)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateUser: <%+v>", user)
	}
	return &empty.Empty{}, err
}

func (s Server) DeleteUser(_ context.Context, id *api.UserId) (*empty.Empty, error) {
	err := db.DeleteUser(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.DeleteUser: <%+v>", id)
	}
	return &empty.Empty{}, err
}
