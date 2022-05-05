package handlers

import (
	"AggreBot/api"
	"AggreBot/internal/app/db"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (s Server) AddGroup(_ context.Context, req *api.AddGroupRequest) (*api.GroupId, error) {
	id, err := db.AddGroup(req)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.AddGroup: <%+v>", req)
	}
	return id, err
}

func (s Server) ListUserGroups(_ context.Context, userId *api.UserId) (*api.ListUserGroupsResponse, error) {
	groups, err := db.ListUserGroups(userId)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.ListUserGroups: <%+v>", userId)
	}
	return &api.ListUserGroupsResponse{Groups: groups}, err
}

func (s Server) UpdateGroup(_ context.Context, group *api.Group) (*empty.Empty, error) {
	err := db.UpdateGroup(group)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateGroup: <%+v>", group)
	}
	return &empty.Empty{}, err
}

func (s Server) DeleteGroup(_ context.Context, id *api.GroupId) (*empty.Empty, error) {
	err := db.DeleteGroup(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.DeleteGroup: <%+v>", id)
	}
	return &empty.Empty{}, err
}
