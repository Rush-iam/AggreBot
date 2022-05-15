package grpc_client

import "AggreBot/internal/pkg/api"

func (c *Client) GetUserSources(userId int64) ([]*api.Source, error) {
	sourcesResponse, err := c.api.GetUserSources(
		c.ctx,
		&api.UserId{
			Id: userId,
		},
	)
	return sourcesResponse.Sources, err
}

func (c *Client) AddSource(userId int64, name, url string) error {
	_, err := c.api.AddSource(
		c.ctx,
		&api.AddSourceRequest{
			UserId: userId,
			Name:   name,
			Url:    url,
		},
	)
	return err
}

func (c *Client) GetSource(sourceId int64) (*api.Source, error) {
	source, err := c.api.GetSource(
		c.ctx,
		&api.SourceId{
			Id: sourceId,
		},
	)
	return source, err
}

func (c *Client) UpdateSourceName(sourceId int64, name string) error {
	_, err := c.api.UpdateSourceName(
		c.ctx,
		&api.UpdateSourceNameRequest{
			Id:   sourceId,
			Name: name,
		},
	)
	return err
}

func (c *Client) UpdateSourceIsActive(sourceId int64, isActive bool) (*api.UpdateSourceIsActiveResponse, error) {
	toggleResponse, err := c.api.UpdateSourceIsActive(
		c.ctx,
		&api.UpdateSourceIsActiveRequest{
			Id:       sourceId,
			IsActive: isActive,
		},
	)
	return toggleResponse, err
}

func (c *Client) DeleteSource(sourceId int64) error {
	_, err := c.api.DeleteSource(
		c.ctx,
		&api.SourceId{
			Id: sourceId,
		},
	)
	return err
}
