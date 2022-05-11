package commands

import "fmt"

func cmdToggleReply(sourceIsActive bool, sourceName string) string {
	return fmt.Sprintf("%c %s", boolToEmoji(sourceIsActive), sourceName)
}

func (m *Manager) cmdToggle(c *command) string {
	sourceToToggle, errReply := m.getSourceFromUserArg(c.userId, c.args)
	if errReply != "" {
		return errReply
	}

	source, err := m.backend.UpdateSourceToggleActive(sourceToToggle.Id)
	if err != nil {
		return errInternalError
	}

	return cmdToggleReply(source.IsActive, source.Name)
}
