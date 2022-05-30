package backend

import (
	"AggreBot/internal/pkg/api"
	"context"
	"log"
)

func (s Server) AddSource(_ context.Context, req *api.AddSourceRequest) (*api.AddSourceResponse, error) {
	source, err := s.db.AddSource(req.UserId, req.Name, req.Url)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.AddSource: <%+v>", req)
	}
	return &api.AddSourceResponse{Source: source}, err
}

func (s Server) GetSource(_ context.Context, req *api.GetSourceRequest) (*api.GetSourceResponse, error) {
	source, err := s.db.GetSource(req.Id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetSource: <%+v>", req.Id)
	}
	return &api.GetSourceResponse{Source: source}, err
}

func (s Server) GetUserSources(_ context.Context, req *api.GetUserSourcesRequest) (*api.GetUserSourcesResponse, error) {
	sources, err := s.db.GetUserSources(req.Id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.GetUserSources: <%+v>", req.Id)
	}
	return &api.GetUserSourcesResponse{Sources: sources}, err
}

func (s Server) UpdateSourceName(_ context.Context, req *api.UpdateSourceNameRequest) (*api.UpdateSourceNameResponse, error) {
	err := s.db.UpdateSourceName(req.Id, req.Name)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateSource: <%+v>", req)
	}
	return &api.UpdateSourceNameResponse{}, err
}

func (s Server) UpdateSourceIsActive(_ context.Context, req *api.UpdateSourceIsActiveRequest) (*api.UpdateSourceIsActiveResponse, error) {
	source, err := s.db.UpdateSourceIsActive(req.Id, req.IsActive)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.UpdateSourceIsActive: <%+v>", req)
	}
	return &api.UpdateSourceIsActiveResponse{Source: source}, err
}

func (s Server) DeleteSource(_ context.Context, req *api.DeleteSourceRequest) (*api.DeleteSourceResponse, error) {
	err := s.db.DeleteSource(req.Id)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("s.DeleteSource: <%+v>", req.Id)
	}
	return &api.DeleteSourceResponse{}, err
}
