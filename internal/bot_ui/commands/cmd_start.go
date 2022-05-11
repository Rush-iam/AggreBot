package commands

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *Manager) cmdStart(c *command) *string {
	var reply string
	err := m.backend.AddUser(c.userId)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		reply = "⚠ Oops. Internal Error. Please try again later."
	} else {
		reply = "🤖"
	}
	return &reply
}
