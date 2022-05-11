package commands

import "fmt"

func (m *Manager) cmdToggle(c *command) *string {
	var reply string

	sourceToToggle, errReply := m.getSourceFromUserArg(c.userId, c.args)
	if errReply != nil {
		return errReply
	}

	source, err := m.backend.UpdateSourceToggleActive(sourceToToggle.Id)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	reply = fmt.Sprintf("%c %s", boolToEmoji(source.IsActive), source.Name)
	return &reply
}
