package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"fmt"
	"strings"
)

func cmdRenameReply(sourceToRenameIsActive bool, newName string) string {
	return fmt.Sprintf("%c %s", markup.BoolToEmoji(sourceToRenameIsActive), newName)
}

func (m *Manager) cmdRename(c *command.Command) string {
	sourceToRename, errReply := m.getSourceFromUserArg(c.UserId, c.Args)
	if errReply != "" {
		return errReply
	}

	if len(c.Args) == 1 {
		return errors.ErrRenameNoName
	}

	newName := strings.Join(c.Args[1:], " ")
	err := m.backend.UpdateSourceName(sourceToRename.Id, newName)
	if err != nil {
		return errors.ErrInternalError
	}

	return cmdRenameReply(sourceToRename.IsActive, newName)
}
