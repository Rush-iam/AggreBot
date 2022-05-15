package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"fmt"
)

func cmdDeleteReply(sourceToDeleteName string) string {
	return fmt.Sprintf("🗑 %s", sourceToDeleteName)
}

func (m *Manager) cmdDelete(c *command.Command) string {
	sourceToDelete, errReply := m.getSourceFromUserArg(c.UserId, c.Text)
	if errReply != "" {
		return errReply
	}

	err := m.backend.DeleteSource(sourceToDelete.Id)
	if err != nil {
		return errors.ErrInternalError
	}

	return cmdDeleteReply(sourceToDelete.Name)
}
