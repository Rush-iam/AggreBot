package grpc_client

import "github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"

func (c *Client) AddUser(userId int64) error {
	_, err := c.api.AddUser(
		c.ctx,
		&api.AddUserRequest{Id: userId},
	)
	return err
}

func (c *Client) GetUser(userId int64) (*api.User, error) {
	resp, err := c.api.GetUser(
		c.ctx,
		&api.GetUserRequest{
			Id: userId,
		},
	)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
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
		&api.UpdateUserFilterRequest{
			User: &api.User{
				Id:     userId,
				Filter: newFilter,
			},
		},
	)
	return err
}
