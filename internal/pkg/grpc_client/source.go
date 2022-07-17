package grpc_client

import "github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"

func (c *Client) GetUserSources(userId int64) ([]*api.Source, error) {
	resp, err := c.api.GetUserSources(
		c.ctx,
		&api.GetUserSourcesRequest{
			Id: userId,
		},
	)
	if err != nil {
		return nil, err
	}
	return resp.Sources, nil
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
	resp, err := c.api.GetSource(
		c.ctx,
		&api.GetSourceRequest{
			Id: sourceId,
		},
	)
	if err != nil {
		return nil, err
	}
	return resp.Source, nil
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

func (c *Client) UpdateSourceIsActive(sourceId int64, isActive bool) (*api.Source, error) {
	resp, err := c.api.UpdateSourceIsActive(
		c.ctx,
		&api.UpdateSourceIsActiveRequest{
			Id:       sourceId,
			IsActive: isActive,
		},
	)
	if err != nil {
		return nil, err
	}
	return resp.Source, nil
}

func (c *Client) DeleteSource(sourceId int64) error {
	_, err := c.api.DeleteSource(
		c.ctx,
		&api.DeleteSourceRequest{
			Id: sourceId,
		},
	)
	return err
}
