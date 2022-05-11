package commands

import (
	"fmt"
	"regexp"
)

func cmdFilterReply(userFilter string) string {
	if userFilter == "" {
		return "🔍 No Filter set"
	} else {
		return fmt.Sprintf("🔍 RegExp Filter: '%s'", userFilter)
	}
}

func (m *Manager) cmdFilter(c *command) string {
	var userFilter string
	if len(c.args) > 0 {
		userFilter = c.args[0]
	}

	if len([]rune(userFilter)) > 256 {
		return errFilterTooLong
	}

	_, err := regexp.Compile(userFilter)
	if err != nil {
		return errFilterRegExp
	}

	err = m.backend.UpdateUserFilter(c.userId, userFilter)
	if err != nil {
		return errInternalError
	}

	return cmdFilterReply(userFilter)
}
