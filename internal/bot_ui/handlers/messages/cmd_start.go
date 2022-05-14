package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func cmdStartReply() string {
	return "🤖"
}

func (m *Manager) cmdStart(c *command.Command) string {
	err := m.backend.AddUser(c.UserId)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		return errors.ErrInternalError
	}

	return cmdStartReply()
}
