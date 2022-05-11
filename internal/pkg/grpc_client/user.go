package grpc_client

import "AggreBot/internal/pkg/api"

func (c *Client) AddUser(userId int64) error {
	_, err := c.api.AddUser(
		c.ctx,
		&api.UserId{Id: userId},
	)
	return err
}

func (c *Client) GetUserFilter(userId int64) (*string, error) {
	responseUser, err := c.api.GetUser(
		c.ctx,
		&api.UserId{
			Id: userId,
		},
	)
	if err != nil {
		return nil, err
	}
	return &responseUser.Filter, nil
}

func (c *Client) UpdateUserFilter(userId int64, newFilter string) error {
	_, err := c.api.UpdateUserFilter(
		c.ctx,
		&api.User{
			Id:     userId,
			Filter: newFilter,
		},
	)
	return err
}
