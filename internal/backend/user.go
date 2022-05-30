package backend

import (
	"AggreBot/internal/pkg/api"
	"context"
	"log"
)

func (s Server) AddUser(_ context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	err := s.db.AddUser(&api.User{Id: req.Id})
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.AddUser: <%+v>", req.Id)
	}
	return &api.AddUserResponse{}, err
}

func (s Server) GetUser(_ context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	user, err := s.db.GetUser(req.Id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetUser: <%+v>", req.Id)
	}
	return &api.GetUserResponse{User: user}, err
}

func (s Server) UpdateUserFilter(_ context.Context, req *api.UpdateUserFilterRequest) (*api.UpdateUserFilterResponse, error) {
	err := s.db.UpdateUserFilter(req.User)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateUser: <%+v>", req.User)
	}
	return &api.UpdateUserFilterResponse{}, err
}

func (s Server) DeleteUser(_ context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	err := s.db.DeleteUser(req.Id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.DeleteUser: <%+v>", req.Id)
	}
	return &api.DeleteUserResponse{}, err
}
