package commands

import (
	"fmt"
	"strings"
)

func cmdRenameReply(sourceToRenameIsActive bool, newName string) string {
	return fmt.Sprintf("%c %s", boolToEmoji(sourceToRenameIsActive), newName)
}

func (m *Manager) cmdRename(c *command) string {
	sourceToRename, errReply := m.getSourceFromUserArg(c.userId, c.args)
	if errReply != "" {
		return errReply
	}

	if len(c.args) == 1 {
		return errRenameNoName
	}

	newName := strings.Join(c.args[1:], " ")
	err := m.backend.UpdateSourceName(sourceToRename.Id, newName)
	if err != nil {
		return errInternalError
	}

	return cmdRenameReply(sourceToRename.IsActive, newName)
}
