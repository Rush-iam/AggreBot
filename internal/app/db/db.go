package db

import (
	"AggreBot/api"
)

type dummyDB struct {
	users         map[int64]api.User
	groups        map[int64]api.Group
	groupsCounter int64
	sources       map[int64]api.Source
}

var db dummyDB

func init() {
	// connect to DB here
	db.users = make(map[int64]api.User)
	db.groups = make(map[int64]api.Group)
	db.sources = make(map[int64]api.Source)
}
