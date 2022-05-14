package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"fmt"
)

func cmdToggleReply(sourceIsActive bool, sourceName string) string {
	return fmt.Sprintf("%c %s", markup.BoolToEmoji(sourceIsActive), sourceName)
}

func (m *Manager) cmdToggle(c *command.Command) string {
	sourceToToggle, errReply := m.getSourceFromUserArg(c.UserId, c.Args)
	if errReply != "" {
		return errReply
	}

	source, err := m.backend.UpdateSourceToggleActive(sourceToToggle.Id)
	if err != nil {
		return errors.ErrInternalError
	}

	return cmdToggleReply(source.IsActive, source.Name)
}
