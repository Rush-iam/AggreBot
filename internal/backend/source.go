package backend

import (
	"AggreBot/internal/pkg/api"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (s Server) AddSource(_ context.Context, req *api.AddSourceRequest) (*api.SourceId, error) {
	id, err := s.db.AddSource(req)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.AddSource: <%+v>", req)
	}
	return id, err
}

func (s Server) GetSource(_ context.Context, id *api.SourceId) (*api.Source, error) {
	source, err := s.db.GetSource(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetSource: <%+v>", id)
	}
	return source, err
}

func (s Server) GetUserSources(_ context.Context, userId *api.UserId) (*api.Sources, error) {
	sources, err := s.db.GetUserSources(userId)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetUserSources: <%+v>", userId)
	}
	return sources, err
}

func (s Server) UpdateSourceName(_ context.Context, req *api.UpdateSourceNameRequest) (*empty.Empty, error) {
	err := s.db.UpdateSourceName(req)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateSource: <%+v>", req)
	}
	return &empty.Empty{}, err
}

func (s Server) UpdateSourceToggleActive(_ context.Context, id *api.SourceId) (*api.SourceToggleActiveResponse, error) {
	source, err := s.db.UpdateSourceToggleActive(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateSourceToggleActive: <%+v>", id)
	}
	return source, err
}

func (s Server) DeleteSource(_ context.Context, id *api.SourceId) (*empty.Empty, error) {
	err := s.db.DeleteSource(id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.DeleteSource: <%+v>", id)
	}
	return &empty.Empty{}, err
}
