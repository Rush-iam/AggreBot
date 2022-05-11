package commands

import "fmt"

func (m *Manager) cmdDelete(c *command) *string {
	var reply string
	sourceToDelete, errReply := m.getSourceFromUserArg(c.userId, c.args)
	if errReply != nil {
		return errReply
	}
	err := m.backend.DeleteSource(sourceToDelete.Id)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	reply = fmt.Sprintf("🗑 %s", sourceToDelete.Name)
	return &reply
}
