package commands

import (
	"fmt"
	"strings"
)

func (m *Manager) cmdRename(c *command) *string {
	var reply string
	sourceToRename, errReply := m.getSourceFromUserArg(c.userId, c.args)
	if errReply != nil {
		return errReply
	}
	if len(c.args) == 1 {
		reply = "👉 Hey! You forgot to provide new name"
		return &reply
	}
	newName := strings.Join(c.args[1:], " ")
	err := m.backend.UpdateSourceName(sourceToRename.Id, newName)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	reply = fmt.Sprintf("%c %s", boolToEmoji(sourceToRename.IsActive), newName)
	return &reply
}
