package commands

import "fmt"

func cmdDeleteReply(sourceToDeleteName string) string {
	return fmt.Sprintf("🗑 %s", sourceToDeleteName)
}

func (m *Manager) cmdDelete(c *command) string {
	sourceToDelete, errReply := m.getSourceFromUserArg(c.userId, c.args)
	if errReply != "" {
		return errReply
	}

	err := m.backend.DeleteSource(sourceToDelete.Id)
	if err != nil {
		return errInternalError
	}

	return cmdDeleteReply(sourceToDelete.Name)
}
