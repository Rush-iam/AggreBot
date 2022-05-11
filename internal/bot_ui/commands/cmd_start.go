package commands

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func cmdStartReply() string {
	return "🤖"
}

func (m *Manager) cmdStart(c *command) string {
	err := m.backend.AddUser(c.userId)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		return errInternalError
	}

	return cmdStartReply()
}
