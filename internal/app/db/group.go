package db

import (
	"AggreBot/api"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddGroup(req *api.AddGroupRequest) (*api.GroupId, error) {
	if isUserExists(req.UserId) == false {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.AddGroup: <%+v> user not found", req),
		)
	}
	return addGroupQuery(req)
}

func addGroupQuery(req *api.AddGroupRequest) (*api.GroupId, error) {
	id := db.groupsCounter // id mutex
	db.groupsCounter++
	db.groups[id] = api.Group{
		Id:     id,
		UserId: req.UserId,
		Active: true,
		Name:   req.Name,
		Filter: req.Filter,
	}
	return &api.GroupId{Id: id}, nil
}

func ListUserGroups(userId *api.UserId) ([]*api.Group, error) {
	if isUserExists(userId.Id) == false {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.ListUserGroups: <%+v> not found", userId),
		)
	}
	return listUserGroupsQuery(userId)
}

func listUserGroupsQuery(userId *api.UserId) ([]*api.Group, error) {
	var groups []*api.Group
	for _, group := range db.groups {
		if group.UserId == userId.Id {
			groups = append(groups, &group)
		}
	}
	return groups, nil
}

func UpdateGroup(group *api.Group) error {
	if isGroupExists(group.Id) == false {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.UpdateGroup: <%+v> group not found", group),
		)
	}
	return updateGroupQuery(group)
}

func updateGroupQuery(group *api.Group) error {
	db.groups[group.Id] = api.Group{
		Id:     group.Id,
		UserId: group.UserId,
		Active: group.Active,
		Name:   group.Name,
		Filter: group.Filter,
	}
	return nil
}

func DeleteGroup(id *api.GroupId) error {
	if isGroupExists(id.Id) == false {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("db.DeleteGroup: <%+v> group not found", id),
		)
	}
	return deleteGroupQuery(id)
}

func deleteGroupQuery(id *api.GroupId) error {
	delete(db.groups, id.Id)
	return nil
}

func isGroupExists(id int64) bool {
	_, found := db.groups[id]
	return found
}
