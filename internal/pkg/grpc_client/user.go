package grpc_client

import "AggreBot/internal/pkg/api"

func (c *Client) AddUser(userId int64) error {
	_, err := c.api.AddUser(
		c.ctx,
		&api.UserId{Id: userId},
	)
	return err
}

func (c *Client) GetUser(userId int64) (*api.User, error) {
	user, err := c.api.GetUser(
		c.ctx,
		&api.UserId{
			Id: userId,
		},
	)
	return user, err
}

func (c *Client) GetUserFilter(userId int64) (*string, error) {
	user, err := c.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return &user.Filter, nil
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
