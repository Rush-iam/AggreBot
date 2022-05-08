package handlers

import (
	"AggreBot/api"
	"AggreBot/internal/app/db"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (s Server) AddSource(_ context.Context, req *api.AddSourceRequest) (*api.SourceId, error) {
	id, err := db.AddSource(req)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.AddSource: <%+v>", req)
	}
	return id, err
}

func (s Server) GetSource(_ context.Context, id *api.SourceId) (*api.Source, error) {
	source, err := db.GetSource(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetSource: <%+v>", id)
	}
	return source, err
}

func (s Server) GetUserSources(_ context.Context, userId *api.UserId) (*api.Sources, error) {
	sources, err := db.GetUserSources(userId)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetUserSources: <%+v>", userId)
	}
	return sources, err
}

func (s Server) UpdateSourceName(_ context.Context, source *api.UpdateSourceNameRequest) (*empty.Empty, error) {
	err := db.UpdateSourceName(source)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateSource: <%+v>", source)
	}
	return &empty.Empty{}, err
}

func (s Server) DeleteSource(_ context.Context, id *api.SourceId) (*empty.Empty, error) {
	err := db.DeleteSource(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.DeleteSource: <%+v>", id)
	}
	return &empty.Empty{}, err
}
